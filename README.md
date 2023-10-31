
# UPLOAD VERSE
This service uploads file to IPFS. The IPFS node is hosted on docker and exposes an API that interacts with the CLI using golang.

#### Requirements:
- Docker: The system must have docker installed

#### Use the following command to start the project:
```
docker-compose up -d
```

This will startup every service in the `docker-compose.yaml` file.

Use the following command below to add CORS to the web UI of our IPFS node

```bash
docker-compose exec ipfs ipfs config --json API.HTTPHeaders.Access-Control-Allow-Origin '["http://0.0.0.0:5001", "http://localhost:3000", "http://127.0.0.1:5001", "https://webui.ipfs.io"]'
```

```bash
docker-compose exec ipfs ipfs config --json API.HTTPHeaders.Access-Control-Allow-Methods '["PUT", "POST"]'
```

To view IPFS mode UI, go to [IPFS web UI ](http://localhost:5001/webui)

The URL to the postman collection is  [Postman Collection](https://documenter.getpostman.com/view/4455541/2s9YXb95fj)


### Folder structure
The structure of this code is like plug and play, in the sense that any component can be removed and replaced with minimal effort and *NO* bug,
let's take the DB driver for instance, you can take this code and remove postgres and replace it with another DB driver, all you have to do is change the connection
and implement the file service interface.
The approach is simple 3 steps:
- Root package is for domain types
- Group sub-packages by dependency
- Main package ties together dependencies

> "These rules help isolate our packages and define a clear domain language across the entire application"

The code uses dependency injection to manage the flow, and the cons of this folder structure includes but not limited to:
- Removes chances of circular dependencies
- Takes care of atrocious types names like controller.FileController. (Names are the best form of documentation)

## Tests
To test this application, I'll first have to write a mock service. 
This mock lets me inject functions into anything that uses the uploadverse.UserService interface to validate arguments, return expected data, or inject failures.

#### Unit test
For unit test, I'd be testing the individual function for primary services, e.g. ipfs upload and get file service.

#### End 2 End Test
For E2E test, I would use the `net/http/httptest` to test the API and use the mock service to functions and return data or inject failure.

> go test files is always in the format *_test.go e.g. postgres_test.go

