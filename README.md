# web-service-gin

I am developing a RESTful API with Go and Gin.

## API Endpoints

This API provides access to a store selling vintage recordings on vinyl. A client can get and add albums from vinyl.com

### `/albums`

*   **GET**: Get a list of all albums, returned as JSON.
*   **POST**: Add a new album from request data sent as JSON.

### `/albums/:id`

*   **GET**: Get an album by its ID, returning the album data as JSON.