# ReelingIt

## Project Description
ReelingIt is a full-stack web application designed to manage and explore movies, actors, and genres, allowing users to create and manage their personal watchlists and favorite movie collections. The application provides a seamless experience for browsing movie details, searching for specific titles, and interacting with a personalized movie library.

## Features
- **User Authentication:** Secure user registration and login.
- **Movie Browsing:** Explore a wide collection of movies with detailed information.
- **Search Functionality:** Easily find movies by title, actor, or genre.
- **Personalized Watchlist:** Add movies to your watchlist for future viewing.
- **Favorite Movies:** Mark movies as favorites for quick access.
- **Actor and Genre Management:** View details about actors and genres associated with movies.
- **Responsive Design:** A user-friendly interface that adapts to various screen sizes.

## Technologies Used

### Backend
- **Go (Golang):** For building high-performance and scalable API services.
- **Gorilla Mux:** A powerful URL router and dispatcher for Go.
- **JWT (JSON Web Tokens):** For secure user authentication and authorization.
- **SQL Database:** (e.g., PostgreSQL, MySQL) for data storage, managed via `database-dump.sql`.
- **Logging:** Custom logging implementation for application monitoring (`logger/`).

### Frontend
- **Vanilla JavaScript:** For dynamic and interactive user interfaces.
- **HTML5:** Structure of the web pages.
- **CSS3:** Styling and visual presentation.
- **Client-side Routing:** Managed by `public/services/Router.js` and `public/services/Routes.js`.
- **API Interaction:** Handled by `public/services/API.js` for fetching and sending data.
- **State Management:** Simple global state management using `public/services/Store.js`.
- **Components:** Modular UI elements (`public/components/`) for various parts of the application (e.g., `HomePage.js`, `MovieDetails.js`, `LoginPage.js`).

### Development & Deployment
- **Docker & Docker Compose:** For containerizing the application and its services, ensuring consistent development and deployment environments.
- **Air:** (Optional) For live-reloading Go applications during development.

## Project Structure
- `handlers/`: Contains Go handlers for API endpoints (account, movie).
- `models/`: Defines Go structs for data models (actor, genre, movie, user).
- `providers/`: Contains Go interfaces and implementations for data repositories.
- `public/`: Frontend assets including HTML, CSS, JavaScript, and images.
  - `public/components/`: Reusable JavaScript components for the frontend.
  - `public/services/`: JavaScript services for API interaction, routing, and state management.
- `token/`: Go code for JWT token creation and validation.
- `logger/`: Go package for logging.
- `import/`: Database dump and installation scripts.

## Setup Instructions

### Prerequisites
- Go (version 1.16 or higher recommended)
- Docker and Docker Compose
- A SQL database (e.g., PostgreSQL, MySQL)

### Backend Setup
1. **Clone the repository:**
   ```bash
   git clone https://github.com/ibrahimelothmani/ReelingIt.git
   cd ReelingIt
   ```
2. **Set up the Go environment:**
   - Make sure Go is installed and set up.
   - Create a `.env` file in the root directory and configure your database connection.
3. **Build and run the Docker containers:**
   ```bash
   docker-compose up --build
   ```
4. **Access the application:**
   - Open your browser and visit `http://localhost:8080` for the frontend.

### Additional Notes
- The application is designed for demonstration purposes and may not be production-ready.
- For a production environment, consider additional security measures and considerations.
- The database schema is defined in `database-dump.sql`.
- The application uses JWT for authentication, and the secret key is stored in the `.env` file.

## Contributing
Contributions to ReelingIt are welcome! If you have suggestions, bug reports, or feature requests, please open an issue or submit a pull request.


