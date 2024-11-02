WeatherProfiles_Service
Welcome to WeatherProfiles_Service! This application is a robust user authentication and weather profiling service built with Go (Golang), leveraging modern software design patterns and secure authentication mechanisms.

Table of Contents
Features
Technologies Used
Software Design Patterns
Challenges Faced
Getting Started
Prerequisites
Installation
Running the Application
Contributing
License
Features
User Registration and Authentication: Secure user signup and login using hashed passwords and JWT tokens.
Profile Management: Access and update user profiles securely.
Weather Data Integration: Fetch and display weather data (future implementation).
RESTful API: Clean and well-structured API endpoints.
Secure Password Handling: Passwords are hashed using bcrypt before storage.
JWT Authentication: Secure token-based authentication for API endpoints.
Technologies Used
Go (Golang): Main programming language.
Gin: Web framework for building HTTP web services.
GORM: ORM library for Golang, handling database operations.
PostgreSQL: Relational database for data storage.
JWT (JSON Web Tokens): For secure user authentication.
bcrypt: Password hashing function.
Docker: Containerization for consistent development environments (planned).
godotenv: Loading environment variables from .env files.
Software Design Patterns
The application employs several software design patterns to ensure code modularity, maintainability, and scalability:

MVC (Model-View-Controller): Separating the application into Models, Views (not used in API-only applications), and Controllers.
Repository Pattern: Abstracting data access logic to repositories, providing a cleaner separation between the data and domain layers.
Service Layer Pattern: Encapsulating the business logic of the application, making controllers thin and focused on HTTP handling.
Singleton Pattern: Ensuring a single instance of the database connection throughout the application.
Factory Pattern: Used in initializing controllers and services, promoting loose coupling.
Challenges Faced
During the development of WeatherProfiles_Service, several challenges were encountered:

1. JWT_SECRET Environment Variable Not Set
Problem: The application failed to start due to the JWT_SECRET environment variable not being set.

Solution:

Ensured the JWT_SECRET was set in the environment or loaded from a .env file using godotenv.
Modified the initialization function to load the secret and handle cases where it's missing.
2. Password Comparison Failing with Bcrypt
Problem: Users couldn't log in because password comparison failed, even with correct credentials.

Cause:

The Password field in the User model had the json:"-" tag, preventing it from being populated during JSON binding.
This resulted in an empty password being hashed and stored during registration.
Solution:

Changed the JSON tag to json:"password" to allow the password to be received from the client.
Ensured that the password is not returned in responses by omitting it or using separate response structs.
3. Type Mismatch with JWT Key
Problem: An error occurred during token generation due to a type mismatch in the JWT signing key.

Solution:

Updated the jwtKey variable to be of type []byte instead of string, as required by the SignedString method.
Loaded the key from the environment securely.
4. Improper Error Handling and Logging
Problem: Difficulty in debugging due to insufficient logging and error messages.

Solution:

Enhanced logging throughout the application, especially in areas dealing with authentication and token generation.
Provided clear and descriptive error messages to assist in debugging.
Getting Started
Prerequisites
Go: Version 1.16 or higher.
PostgreSQL: Ensure you have a PostgreSQL server running.
Git: For cloning the repository.
cURL or Postman: For testing API endpoints.
Installation
Clone the repository:

bash
Copy code
git clone https://github.com/yourusername/WeatherProfiles_Service.git
cd WeatherProfiles_Service
Set up environment variables:

Create a .env file in the root directory:

env
Copy code
JWT_SECRET=your_secret_key
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_db_password
DB_NAME=userdb
DB_PORT=5432
Install dependencies:

bash
Copy code
go mod download
Set up the database:

Ensure PostgreSQL is running.
Create a database named userdb or as specified in your .env.
Running the Application
bash
Copy code
go run main.go
The server will start on http://localhost:8080.

Testing Endpoints
Registration:

bash
Copy code
curl -X POST -H "Content-Type: application/json" \
-d '{"username":"testuser","email":"test@example.com","password":"password123456"}' \
http://localhost:8080/register
Login:

bash
Copy code
curl -X POST -H "Content-Type: application/json" \
-d '{"email":"test@example.com","password":"password123456"}' \
http://localhost:8080/login
Access Protected Route:

bash
Copy code
curl -H "Authorization: Bearer your_jwt_token" \
http://localhost:8080/profile
Contributing
Contributions are welcome! Please fork the repository and submit a pull request for any enhancements or bug fixes.

License
This project is licensed under the MIT License.

