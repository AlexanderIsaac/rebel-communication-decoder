# Rebel Communication Decoder

Welcome to the Rebel Communication Decoder! This Go application is designed to decode distress signals from space using data from three (in the future more!) satellites.
## Features
## Core Functions

#### `GetLocation`

**Description**: Calculates the coordinates `(x, y)` of the message sender based on distances from the satellites.

**Parameters**:
- `distances`: A variadic list of distances from each satellite to the message sender.

**Returns**:
- `x, y`: Coordinates of the message sender.

#### `GetMessage`

**Description**: Reconstructs the original message from the fragments received by the satellites.

**Parameters**:
- `messages`: A variadic list of message arrays from each satellite.

**Returns**:
- `message`: The reconstructed message.

## Getting Started

1. **Clone the repository**:
   ```sh
   git clone https://github.com/AlexanderIsaac/rebel-communication-decoder.git
2. **Install dependencies**:
    ```sh
    cd rebel-communication-decoder
    go mod tidy
3. **Create a .env file**:
   1. **Copy the example .env file:**:
         ```sh
         cp .env.example .env
    2. **Open the .env file in a text editor and configure it**:
         ```sh
        GOOGLE_CLOUD_PROJECT=
        GOOGLE_CLOUD_FIRESTORE_DB=
        GOOGLE_CLOUD_CREDENTIALS=
    Ensure that the .env file contains all necessary environment variables required.
4. **Run the application:**:
     ```sh
     go run main.go
5. **Explore the API documentation:**
Access the Swagger UI for live API documentation at http://localhost:8080/swagger/index.html after running the application. The Swagger configuration can also be found in the [docs/swagger.json](docs/swagger.json) file.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

For any issues or feature requests, please open an issue on the [GitHub repository](https://github.com/AlexanderIsaac/rebel-communication-decoder/issues).
