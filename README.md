# mygram-api-go

## Pembagian Tugas

| Bagian     | Detail             | Dikerjakan oleh          |
|------------|--------------------|--------------------------|
| Helper     | Generate Token     | Muhammad Rifqi Al Furqon |
| Helper     | Validate Token     | Muhammad Rifqi Al Furqon |
| Middleware | Authorization      | Muhammad Rifqi Al Furqon |
| Endpoint   | GET /users         | Muhammad Rifqi Al Furqon |
| Endpoint   | POST /users/:id    | Muhammad Rifqi Al Furqon |
| Endpoint   | PUT /users/:id     | Muhammad Rifqi Al Furqon |
| Endpoint   | DELETE /users/:id  | Muhammad Rifqi Al Furqon |
| Endpoint   | GET /photos        | Muhammad Rifqi Al Furqon |
| Endpoint   | POST /photos       | Muhammad Rifqi Al Furqon |
| Endpoint   | DELETE /photos/:id | Muhammad Rifqi Al Furqon |
| Endpoint   | PUT /photos/:id    | Muhammad Rifqi Al Furqon |
| Endpoint   | GET /comments        | Juan Simon Damanik |
| Endpoint   | POST /comments       | Juan Simon Damanik |
| Endpoint   | DELETE /comments/:id | Juan Simon Damanik |
| Endpoint   | PUT /comments/:id    | Juan Simon Damanik |
| Swagger   | Add Swagger Docs    | Juan Simon Damanik |
| Endpoint	 | POST /socialmedia/:id| Dion Fauzi 			   |
| Endpoint	 | GET /socialmedia/ :id| Dion Fauzi			   |
| Endpoint	 | DELETE /socialmedia 	| Dion Fauzi			   |
| Endpoint   | PUT /socialmedia	    | Dion Fauzi			   |

## How to Run
### Locally
- Clone this repo
```
git clone https://github.com/jusidama18/mygram-api-go/
```
- Run PostgreSQL Docker script
```
chmod +x ./scripts/run-postgres.sh && ./scripts/run-postgres.sh
```
- Copy .env.example to .env
```
cp .env.example .env
```
- Run go webserver
```
go run ./main.go
```
- Enjoy!

### Docker Compose
- Clone this repo
```
git clone https://github.com/jusidama18/mygram-api-go/
```
- Copy .env.example to .env
```
cp .env.example .env
```
- Edit **DB_HOST** variable in .env file
- Run `docker compose up`
- Enjoy!

## Deployment
Project deployed on Google Compute Engine with this URL `https://mygram-api-go-production.up.railway.app/`

# Documentation

## Users

### Register new user

- Method: POST
- Endpoint: /users/register
- Request Body:

```json
{
	"age": "integer",
	"email": "string",
	"password": "string",
	"username": "string"
}
```

- Response Body:

  - Status: 201,
  - Body:

  ```json
  {
  	"data": {
  		"age": "integer",
  		"email": "string",
  		"password": "string",
  		"username": "string"
  	}
  }
  ```

### Login

- Method: POST
- Endpoint: /users/login
- Request Body:

```json
{
	"age": "integer",
	"email": "string",
	"password": "string",
	"username": "string"
}
```

- Response:

  - Status: 201
  - Body:

    ```json
    {
    	"age": "integer",
    	"email": "string",
    	"password": "string",
    	"username": "string"
    }
    ```

### Edit User

- Method: PUT
- Endpoint: /users
- Headers: Authorization (Bearer Token)
- Request Body:

```json
{
	"email": "string",
	"username": "string"
}
```

- Response:

  - Status: 200
  - Body:

  ```json
  {
  	"age": "integer",
  	"email": "string",
  	"password": "string",
  	"username": "string"
  }
  ```

Notes: Endpoint ini perlu melewati proses autentikasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken.

### Delete User

- Method: DELETE
- Endpoint: /users
- Response:

  - Status: 200
  - Body:

  ```json
  {
  	"message": "Your account has been successfully deleted"
  }
  ```

Notes: Endpoint ini perlu melewati proses autentikasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken.

## Photos

### Create new photo

- Method: POST
- Endpoint: /photos
- Headers: Authorization (Bearer Token)
- Request Body:

```json
{
	"title": "string",
	"caption": "string",
	"photo_url": "string"
}
```

- Response:

  - Status: 201
  - Body:

  ```json
  {
  	"id": "integer",
  	"title": "string",
  	"caption": "string",
  	"photo_url": "string",
  	"user_id": "integer",
  	"created_at": "date"
  }
  ```

Notes: Endpoint ini perlu melewati proses autentikasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken.

### Get all photos

- Method: GET
- Endpoint: /photos
- Headers: Authorization (Bearer Token)

- Response Body:

  - Status: 200
  - Body:

    ```json
    [
    	{
    		"id": "integer",
    		"title": "string",
    		"caption": "string",
    		"photo_url": "string",
    		"user_id": "integer",
    		"created_at": "date",
    		"updated_at": "date",
    		"user": {
    			"username": "string",
    			"email": "username"
    		}
    	},
    	{
    		"id": "integer",
    		"title": "string",
    		"caption": "string",
    		"photo_url": "string",
    		"user_id": "integer",
    		"created_at": "date",
    		"updated_at": "date",
    		"user": {
    			"username": "string",
    			"email": "username"
    		}
    	}
    ]
    ```

Notes: Endpoint ini perlu melewati proses autentikasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken.

### Edit photo

- Method: PUT
- Endpoint: /photos/:photoId
- Headers: Authorization (Bearer Token)
- Request Body:

```json
{
	"title": "string",
	"caption": "string",
	"photo_url": "string"
}
```

- Response Body:

  - Status: 200
  - Body:

  ```json
  {
      "data": {
          "title": "string",
          "caption": "string",
          "photo_url": "string",
          "user_id": "string",
          "updated_at": "string"
  }
  ```

Notes: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken. Dan alur proses autorisasinya adalah user hanya boleh mengupdate data photo miliknya sendiri.

### Delete photo

- Method: DELETE
- Endpoint: /photos/:photoId
- Headers: Authorization (Bearer Token)
- Response:

  - Status: 200
  - Body

  ```json
  {
  	"message": "Your photo has been successfully deleted"
  }
  ```

Notes: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken. Dan alur proses autorisasinya adalah user hanya boleh menghapus data photo miliknya sendiri.

## Comments

### Create new comment

- Method: POST
- Endpoint: /comments
- Headers: Authorization (Bearer Token)
- Request Body:

```json
{
	"message": "string",
	"photo_id": "integer"
}
```

- Response:

  - Status: 201
  - Body:

  ```json
  {
  	"id": "integer",
  	"message": "string",
  	"photo_id": "string",
  	"user_id": "integer",
  	"created_at": "date"
  }
  ```

Notes: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken.

### Get all comments

- Method: GET
- Endpoint: /comments
- Headers: Authorization (Bearer Token)
- Response:

  - Status: 200
  - Body:

    ```json
    {
    	"id": "integer",
    	"message": "string",
    	"photo_id": "string",
    	"user_id": "integer",
    	"created_at": "date",
    	"user": {
    		"id": "integer",
    		"email": "string",
    		"username": "string"
    	},
    	"photo": {
    		"id": "string",
    		"title": "string",
    		"photo_url": "string",
    		"user_id": "integer"
    	}
    }
    ```

Notes: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken.

### Edit comment

- Method: PUT
- Endpoint: /comments/:commentId
- Headers: Authorization (Bearer Token)
- Request Body:

```json
{
	"message": "string"
}
```

- Response:

  - Status: 200
  - Body:

  ```json
  {
  	"id": "integer",
  	"message": "string",
  	"photo_id": "string",
  	"user_id": "integer",
  	"created_at": "date"
  }
  ```

Notes: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken. Dan alur proses autorisasinya adalah user hanya boleh mengupdate data comment miliknya sendiri.

### Delete comment

- Method: DELETE
- Endpoint: /comments/:commentId
- Headers: Authorization (Bearer Token)
- Response:

  - Status: 200
  - Body:

  ```json
  {
  	"message": "Your photo has been successfully deleted"
  }
  ```

Notes: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken. Dan alur proses autorisasinya adalah user hanya boleh menghapus data comment miliknya sendiri.

## Social Media

### Create new socialmedia

- Method: POST
- Endpoint: /socialmedias
- Headers: Authorization (Bearer Token)
- Request Body:

```json
{
	"name": "string",
	"social_media_url": "string"
}
```

- Response Body:
  - Status: 201
  - Body:

```json
{
	"id": "integer",
	"caption": "string",
	"social_media_url": "string",
	"user_id": "integer",
	"created_at": "date"
}
```

Notes: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken.

### Get all socialmedia

- Method: GET
- Endpoint: /socialmedias
- Headers: Authorization (Bearer Token)

- Response Body:
  - Status: 200

```json
{
	"id": "integer",
	"caption": "string",
	"social_media_url": "string",
	"user_id": "integer",
	"created_at": "date",
	"updated_at": "date",
	"user": {
		"id": "integer",
		"username": "string",
		"profile_image_url": "string"
	}
}
```

Notes: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken.

### Edit socialmedias

- Method: PUT
- Endpoint: /socialmedias/:socialMediaId
- Headers: Authorization (Bearer Token)
- Request Body:

```json
{
	"name": "string",
	"social_media_url": "string"
}
```

- Response:

  - Status Code: 200
  - Body:

    ```json
    {
    	"age": "integer",
    	"email": "string",
    	"password": "string",
    	"username": "string"
    }
    ```

Notes: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken. Dan alur proses autorisasinya adalah user hanya boleh mengupdate data social media miliknya sendiri.

### Delete socialmedias

- Method: DELETE
- Endpoint: /socialmedias/:socialMediaId
- Response:

  - Status: 200
  - Body:

    ```json
    {
    	"message": "Your social media has been successfully deleted"
    }
    ```

Notes: Endpoint ini perlu melewati proses autentikasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken.
