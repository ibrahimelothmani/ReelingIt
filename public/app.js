import "./components/AnimatedLoading.js";
import "./components/YouTubeEmbed.js";
import { API } from "./services/API.js";
import Router from "./services/Router.js";
import Store from "./services/Store.js";

window.app = {
  API,
  Router,
  Store,
  showError: (
    message = "There was an error loading the page",
    goToHome = false
  ) => {
    document.querySelector("#alert-modal").showModal();
    document.querySelector("#alert-modal p").textContent = message;
    if (goToHome) app.Router.go("/");
    return;
  },
  closeError: () => {
    document.getElementById("alert-modal").close();
  },
  search: (event) => {
    event.preventDefault();
    const keywords = document.querySelector("input[type=search]").value;
    if (keywords.length > 1) {
      app.Router.go(`/movies?q=${keywords}`);
    }
  },
  searchOrderChange: (order) => {
    const urlParams = new URLSearchParams(window.location.search);
    const q = urlParams.get("q");
    const genre = urlParams.get("genre") ?? "";
    app.Router.go(`/movies?q=${q}&order=${order}&genre=${genre}`);
  },
  searchFilterChange: (genre) => {
    const urlParams = new URLSearchParams(window.location.search);
    const q = urlParams.get("q");
    const order = urlParams.get("order") ?? "";
    app.Router.go(`/movies?q=${q}&order=${order}&genre=${genre}`);
  },
  register: async (event) => {
    event.preventDefault();
    const name = document.getElementById("register-name").value;
    const email = document.getElementById("register-email").value;
    const password = document.getElementById("register-password").value;
    const passwordConfirmation = document.getElementById(
      "register-password-confirm"
    ).value;

    const errors = [];
    if (name.length < 4) errors.push("Name must be at least 4 characters long");
    if (password.length < 7)
      errors.push("Enter a password with at least 7 characters");
    if (email.length < 4) errors.push("Enter your complete email.");
    if (password != passwordConfirmation) errors.push("Passwords do not match");

    if (errors.length == 0) {
      const response = await API.register(name, email, password);
      if (response.success) {
        app.Store.jwt = response.jwt;
        app.Router.go("/account/");
      } else {
        app.showError(response.message);
      }
    } else {
      app.showError(errors.join(". "));
    }
  },
  login: async (event) => {
    event.preventDefault();
    let errors = [];
    const email = document.getElementById("login-email").value;
    const password = document.getElementById("login-password").value;

    if (email.length < 8) errors.push("Enter your complete email");
    if (password.length < 6) errors.push("Enter a password with 6 characters");
    if (errors.length == 0) {
      const response = await API.login(email, password);
      if (response.success) {
        app.Store.jwt = response.jwt;
        app.Router.go("/account/");
      } else {
        app.showError(response.message, false);
      }
    } else {
      app.showError(errors.join(". "), false);
    }
  },
  logout: () => {
    Store.jwt = null;
    app.Router.go("/");
  },
  saveToCollection: async (movie_id, collection) => {
    if (app.Store.loggedIn) {
      try {
        const response = await API.saveToCollection(movie_id, collection);
        if (response.success) {
          switch (collection) {
            case "favorite":
              app.Router.go("/account/favorites");
              break;
            case "watchlist":
              app.Router.go("/account/watchlist");
          }
        } else {
          app.showError("We couldn't save the movie.");
        }
      } catch (e) {
        console.log(e);
      }
    } else {
      app.Router.go("/account/");
    }
  },
};
window.addEventListener("DOMContentLoaded", () => {
  app.Router.init();
});
