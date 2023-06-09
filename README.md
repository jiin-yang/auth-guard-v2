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

## Notes

My purpose; Preventing external access to employee-api methods that need to be protected and only accessing requested resources through api-gateway.

#### Flow:
```
1-) First of all, you must have a token. You can send a GET /token request to the api gateway.

This request is for you; Filling in client_id, client_secret, grand_type and auth0_audience information,
Retrieves the access token generated by the auth0.com platform.

2-) Now we will use this token to access protected methods.
You must add key:"authorization" value:"Bearer <ACCESS_TOKEN> to the Header section when making a request.

3-) This token you send in the header is verified before going to the employee-api in the api-gateway project.
Validation is done by the 'CustomValidToken' middleware function.

By creating a machine-to-machine application on the Auth0 platform,
We have the necessary tools to implement the client credentials flow.
The client_id, client_secret, domain and target audience addresses of my application are specified in the .env file.

4-) For protected methods in employee service; It is necessary to validate the access token in the incoming request.
If we don't verify, apps with employee service endpoints can access employee service without any protection.

For this, the incoming token is parsed and using the RSA256 algorithm,
Authentication is done with the public key JSON Web Key Set.
For details of this process, check the verifyToken function in employee-api.

```