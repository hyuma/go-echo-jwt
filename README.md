# go-echo-jwt
Test JWT authentication using golang with Echo framework

## APIs
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
