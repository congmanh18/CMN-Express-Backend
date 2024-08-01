
# CMN Express Backend (DeliveryApp)
## Installation
**Clone the repository**
   ```bash
   git clone https://github.com/yourusername/deliveryapp-backend.git
   cd deliveryapp-backend
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
#### Create an Order
```http
POST /api/orders
Content-Type: application/json
{
    "userId": "12345",
    "items": [
        {
            "name": "Pizza",
            "quantity": 1
        },
        {
            "name": "Soda",
            "quantity": 2
        }
    ],
    "deliveryAddress": "123 Main St, City, Country"
}
```

#### Get Order Details
```http
GET /api/orders/{orderId}
```

## Configuration
All configuration options can be set using environment variables. Below are some of the key variables:
- `DB_HOST`: Database host
- `DB_PORT`: Database port
- `DB_USER`: Database user
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name

## Contributing
We welcome contributions! Please read our [Contributing Guidelines](CONTRIBUTING.md) for more details.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact
If you have any questions or need further assistance, feel free to contact us at [nguyenmanh180102@gmail.com](mailto:nguyenmanh180102@gmail.com).
```

Bạn có thể tùy chỉnh thêm nội dung và cấu trúc để phù hợp với yêu cầu và tính năng cụ thể của project của bạn.
