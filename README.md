# Boolipi - A Go API for booleans for CRUD operations

A Go API for booleans, supporting CRUD operations for booleans with JWT authorization

## API

An entry consists of an `id` (uuid), a boolean `value` (boolean not string) and a `key` (string)

### Add a new Boolean
``` 
POST /
request:

{
  "value":true,
   "key": "name" // this is optional
}

response:

{
  "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "value": true,
  "key": "name"
} 
```

Usage :
* Without Authorization : 
  ```console
  curl -X POST http://localhost:8080 --header "Content-Type: application/json" --data '{"value": true, "label": "Hello world!"}'
  ```
* With Authorization :
  ```console
  curl -X POST http://localhost:8080 --header "Content-Type: application/json" --data '{"value": true, "label": "Hello world!"}' --header "Authorization: Token [token]
  ```

### Retrieve a new Boolean

```
GET /:id
response:

{
  "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "value": true,
  "key": "name"
}
```

* Without Authorization : 
  ```console
  curl -X GET http://localhost:8080/[id]
  ```
* With Authorization :
  ```console
  curl -X GET http://localhost:8080/[id] --header "Authorization: Token [token]
  ```

### Update a new Boolean
```
PATCH /:id
request:

{
  "value":false,
  "key": "new name" // this is optional
}

response:

{
  "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "value": false,
  "key": "new name"
}
```

Usage :
* Without Authorization : 
  ```console
  curl -X PATCH http://localhost:8080/[id] --header "Content-Type: application/json" --data '{"value": true, "label": "Hello world!"}'
  ```
* With Authorization :
  ```console
  curl -X PATCH http://localhost:8080/[id] --header "Content-Type: application/json" --data '{"value": true, "label": "Hello world!"}' --header "Authorization: Token [token]
  ```

### Delete and existing Boolean
```
DELETE /:id
response:
HTTP 204 No Content
```

* Without Authorization : 
  ```console
  curl -X DELETE http://localhost:8080/[id]
  ```
* With Authorization :
  ```console
  curl -X DELETE http://localhost:8080/[id] --header "Authorization: Token [token]
  ```

## HTTP Status codes returned

| HTTP Status Codes   | Explanation  | 
| :------------- | :---------- | 
| 200 | GET, POST, PATCH Request successfully executed | 
| 204  | DELETE request succesfully executed | 
| 400  | Request format is not correct | 
| 401  | Authorization credentials not correct |
| 404  | Requested resource does not exist |
| 500  | Internal server error occured |  

## How to Run

You will need a running MySQL server instance to run the API service. 
### Build from source

1. Clone this repository : 

    ``` git clone http://github.com/theamrendrasingh/boolipi.git ```

2. cd into the cloned source

    ``` cd boolipi ```
3. Build the docker image from the Dockerfile

    ``` docker build -t theamrendrasingh/boolipi```

### Using dockerhub image

Pull the docker image

`docker pull theamrendrasingh/boolipi

### Running the docker container
After building or pulling the dcoker image from dockerhub, use the following command to run it :

    ```console
    % docker run -i -d -t -p 8080:8080 -e DB_USER='user' -e DB_PASS='pass' -e DB_NAME='boolipi' -e DOCKER_MODE=true -e DB_PORT='3306' -e DB_HOST='host.docker.internal' -e USE_AUTH=true theamrendrasingh/boolipi 
    ``` 
  * Pass the required DB_USER (MySQL username), DB_PASS, DB_NAME (the name of database that the API will create/use), DB_PORT
  * If you are using the host's MySQL server or anyother thing which is mapped to hosts localhost, set DB_HOST='host.dcoker.internal' 
  * If you wish to use JWT token authorization, set USE_AUTH=true

## Reference 
https://booleans.io/