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

### Now we are using JWT Authentication, 
it's a token-based auth. So once we logged in we get token that will be used to access shorten and urls endppoints. 

### React progress;
- Now we have login page and dashboard+logout button
- Unauthenticated user can not access dashboard, so once user will be directed to main path once they access dashboard enpoint. How this is works? react read wheter the user have token or not in localStorage. if they do not have token, they willl be redirected to "/". 
- The lack of this concept, react will only see wheter it has the token or not doesn't care wheter the token is right or not.


