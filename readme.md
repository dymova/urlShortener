1. User Registration and Authentication:
    - [x]  Users can create an account by providing a username, password.
    - [x]  User authentication is required to access certain functionalities.
2. URL Shortening:
    - [x]  Users can submit a long URL to the service.
    - [x]  The service generates a unique short URL for the submitted long URL.
    - [x]  The short URL redirects to the original long URL when accessed.
    - [x]  Shortened URLs should be short and easy to remember.
3. URL Management:
    - Users can view a list of their shortened URLs.
    - Each shortened URL is associated with the original long URL, creation timestamp, and the number of times it has been accessed.
    - Users can edit or delete their shortened URLs if needed.
4. URL Analytics:
    - Users can view analytics for their shortened URLs, for example the number of clicks 
5. Public Access:
    - Shortened URLs can be accessed by anyone, even without authentication.
6. Security:
    - [x]  Implement secure user authentication and password storage using encryption and hashing techniques.
    - Protect against common security vulnerabilities such as cross-site scripting (XSS) and cross-site request forgery (CSRF).
7. User Interface:
    - Provide a user-friendly web interface for users to interact with the service.
    - Users should be able to easily navigate through the application, view their shortened URLs, and access analytics.
8. API Endpoints:
    - Implement RESTful API endpoints to allow programmatic access to the URL shortening service.
    - API endpoints should support operations such as creating shortened URLs, retrieving URL analytics, and accessing the original long URL.
9. Error Handling and Validation:
    - Handle and display appropriate error messages for cases such as invalid URLs, duplicate URLs, or incorrect login credentials.
    - Validate user input to ensure the URLs are properly formatted and prevent potential security issues.