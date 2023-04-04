# auth-guard-v2

auth-guard-v2 is a project that includes api-gateway and rest-api, providing authentication with client credentials flow.

The auth0.com platform developed by Okta was used in the development of this project.

## Usage

```
- open a terminal and go to that folder where you want to save the project.

- git clone https://github.com/jiin-yang/auth-guard-v2.git

- cd api-gateway
- go run .

- cd ..

- cd employee-api
- go run .
```


## Requests

Creating a new employee is protected. You can access all employees publicly.

Example curl requests for these requests;

```
1-) GET ACCESS TOKEN

curl --location 'http://localhost:3010/token'
```

```
2-) CREATE EMPLOYEE

curl --location 'http://localhost:3010/employees' \
--header 'authorization: Bearer <YOUR_ACCESS_TOKEN>' \
--header 'Content-Type: application/json' \
--data '{
    "name": "EMPLOYEE_NAME",
    "department": "DEPARTMENT_NAME"
}'
```
```
3-) GET ALL EMPLOYEES

curl --location 'http://localhost:3010/employees'
```