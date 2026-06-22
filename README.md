# shortavee
URL shortener using Golang and React
<br>
This project is focused on the Golang workflow, which uses a 
- router: maps endpoints into a function/handler. Using Gin (high-performance HTTP web framework)
- handler: accepts and validates the request and giving respond to user.
- services: the logic and decision-making. Generate a short code for the URL using the UUID (Universally Unique ID) library from Google.
- repository: communicates with the database using the GORM library (Golang object-relational Mapping).
<br>
In this project, so far, we have 2 endpoints,
- /api/shorten: POST request to shorten a URL
- /api/urls: GET request to show all URLs and its informations.

