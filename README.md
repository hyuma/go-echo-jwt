# go-echo-jwt
Test JWT authentication using golang with Echo framework

- Token expiration time
   - 30 sec


# APIs
## Login
`POST /login`

### Request Parameters
#### HTTP Header
- N/A

#### HTTP Body
|   Key    | Req? |    Description     |
| -------- | ---- | ------------------ |
| name     | Y    | User name to login |
| password | Y    | User password      |

### Response
#### Success
- HTTP Status: `200`
- Body: N/A

#### Error
- WIP


## Get my profile
`GET /profile`

### Request Parameters
#### HTTP Header
- Authorization: Bearer
  - JWT token

#### HTTP Body
N/A

### Response
#### Success
- HTTP Status: `200`
- Body: 
  HTML Web page

#### Error
- WIP
