# web-service-gin

I am developing a RESTful API with Go and Gin.

## API Endpoints

This API provides access to a store selling vintage recordings on vinyl. A client can get and add albums from vinyl.com

### `/albums`

*   **GET**: Get a list of all albums, returned as JSON.
*   **POST**: Add a new album from request data sent as JSON.

### `/albums/:id`

*   **GET**: Get an album by its ID, returning the album data as JSON.

## To Test my APIs follow the following steps:
### 1- git clone https://github.com/thiernohgradiagram/web-service-gin.git
### 2- cd web-service-gin
### 3- go get
### 4- go run .
### 5- open another terminal and run the following curl commands:
<details>
<summary>Get all albums</summary>

```bash
curl http://localhost:8080/albums
```
</details>

<details>
<summary>Add a new album</summary>

```bash
curl http://localhost:8080/albums \
    --include --header \
    "Content-Type: application/json" \
    --request "POST" --data \
    '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

</details>

<details>
<summary>Get album by ID</summary>

```bash
curl http://localhost:8080/albums/2
```
</details>

<details>
<summary>Get album by non-existent ID</summary>

```bash
curl http://localhost:8080/albums/222
```
</details>

### 6- navigate to http://localhost:8080/albums in any web browser 