# Boolipi - A Go API for booleans for CRUD operations

## Representation and Actions

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

### Delete and existing Boolean
```
DELETE /:id
response:
HTTP 204 No Content
```

## How to Run

You will need a running MySQL server instance to run the container. 

### Using dockerhub image

``` console
% docker run -i -d -t -p 8080:8080 -e DB_USER='user' -e DB_PASS='pass' -e DB_NAME='boolipi' -e DOCKER_MODE=true -e DB_PORT='3306' -e DB_HOST='host.docker.internal'  theamrendrasingh/boolipi
``` 

### Build from source

1. Clone this repository : 

    ``` git clone http://github.com/theamrendrasingh/boolipi.git ```

2. cd into the cloned source

    ``` cd boolipi ```
3. Build the docker image from the Dockerfile

    ``` docker build -t theamrendrasingh/boolipi```
4. Run the docker container
    ```console
    % docker run -i -d -t -p 8080:8080 -e DB_USER='user' -e DB_PASS='pass' -e DB_NAME='boolipi' -e DOCKER_MODE=true -e DB_PORT='3306' -e DB_HOST='host.docker.internal'  theamrendrasingh/boolipi 
    ``` 
## Reference 
https://booleans.io/