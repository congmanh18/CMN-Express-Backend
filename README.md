
# CMN Express Backend (DeliveryApp)
## Installation
**Clone the repository**
   ```bash
   git clone https://github.com/congmanh18/CMN-Express-Backend.git .
   ```
**Set up environment variables**
   Create a `.env` file in the root directory and add the necessary environment variables. Refer to the `.env.example` for guidance.

**Install dependencies**
   ```bash
   go mod tidy
   ```

**Run database migrations**
   ```bash
   go run migrate.go
   ```
**Start the server**
   ```bash
   go run main.go
   ```

## Usage
- To start the server, run:
  ```bash
  go run main.go
  ```
- The server will start at `http://localhost:8080` by default.

## API Documentation
The API documentation is available at `http://localhost:8080/swagger/index.html` after starting the server.

### Example API Calls
#### Register
```http
POST /user/register
```
```
Content-Type: application/json
{
    "first_name":"Nguyen",
    "last_name":"Manh",
    "phone":"0977683533",
    "username":"nguyenmanh",
    "password":"12345678",
    "role":"Administrator"
}
```
| Parameter    | Type     | Description                   |
| :----------  | :------- | :-----------------------------|
| `first_name` | `string` |                               |
| `last_name`  | `string` |                               |
| `phone`      | `string` |**Required**. Your phone number|
| `username`   | `string` |                               |
| `password`   | `string` |**Required**. Your password    |
| `role`       | `string` |                               |
#### Login
```http
POST /user/register
```
```
Content-Type: application/json
{
    "phone":"0977683533",
    "password":"12345678",
}
```
| Parameter    | Type     | Description                   |
| :----------  | :------- | :-----------------------------|
| `phone`      | `string` |**Required**. Your phone number|
| `password`   | `string` |**Required**. Your password    |
#### Update info user

## Configuration
All configuration options can be set using environment variables. Below are some of the key variables:
- `DB_HOST`: Database host
- `DB_PORT`: Database port
- `DB_USER`: Database user
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact
If you have any questions or need further assistance, feel free to contact us at [nguyenmanh180102@gmail.com](mailto:nguyenmanh180102@gmail.com).
```

