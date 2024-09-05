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
## ADR 1: Use of Google Cloud App Engine for Deployment
##### Context
This application requires a deployment platform that supports scalable web services, integrates seamlessly with other cloud services, and minimizes infrastructure management overhead. 
##### Decision
Google Cloud App Engine is a fully managed serverless platform that enables developers to build and deploy applications without managing the underlying infrastructure.
##### Justification
1. **Serverless Architecture**: No need to manage servers or infrastructure, allowing focus on the code.
2. **Automatic Scaling**: Automatically adjusts to handle varying traffic loads without manual intervention.
3. **Seamless Integration**: Works well with other Google Cloud services like Firestore and Cloud Functions.
4. **High Availability**: Built-in redundancy and load balancing ensure reliable performance.
5. **Simplified Deployment**: Easy to deploy updates and manage application releases.
6. **Cost Efficiency**: Pay-as-you-go model aligns costs with actual usage.
##### Consequences
- Pros:
    - Reduced infrastructure management with serverless architecture..
    - Automatic scaling and load balancing.
    - Seamless integration with Google Cloud services.
    - Simplified deployment and cost efficiency.
- Cons
    - Dependency on Google Cloud infrastructure.
    - Costs can grow with high traffic volumes if not managed properly.
##### Conclusion
Using Google Cloud App Engine is beneficial for this application because it offers a serverless architecture, automatic scaling, easy integration with other Google Cloud services, and strong security features. This choice ensures efficient deployment, effective operational management, and reliable performance under varying traffic conditions.
## ADR 2: Use of Firestore for Data Storage
##### Context
This application needs a mechanism to store and retrieve information about satellites and the messages they receive. The system must handle partial updates to the information and consolidate data from different satellites to provide the emitter's location and the complete message. 
##### Decision
Firestore is the storage solution for the application. Firestore is a NoSQL database provided by Google Cloud that offers flexible and scalable cloud storage.
##### Justification
1. **Automatic Scalability**: Firestore automatically scales to handle varying traffic and data volumes, ensuring the application performs well under different loads.
2. **Flexible Data Storage**: Firestore's flexible document-based structure supports both structured and unstructured data, which is ideal for managing satellite data and message fragments.
3. **Real-Time Updates**: It provides real-time synchronization, so any updates to satellite data are instantly reflected in the application.
4. **Seamless Integration**: It integrates smoothly with other Google Cloud services, simplifying deployment and management.
5. **Robust Security**: Firestore includes advanced security rules for managing access and protecting data integrity.
6. **Data Durability**: It ensures data persistence through replication and redundancy across multiple locations.
##### Consequences
- Pros:
    - Automatic scalability and handling of variable traffic without manual management.
    - Flexibility in storing unstructured data and real-time capabilities.
    - Integration with the Google Cloud ecosystem simplifies deployment and management.
    - Advanced security features and access control mechanisms.
- Cons
    - Firestore might incur higher costs compared to self-managed solutions for certain use cases, especially with high read/write frequencies.
    - The NoSQL data structure may require adaptation in the data model design if the application in the future has very complex query needs.
##### Conclusion
Using Firestore for this application is a suitable choice due to its scalability, flexibility in data handling, and seamless integration with Google Cloud. These features are crucial for ensuring that the application can grow and adapt to future requirements (like adding much more satellites) while providing reliable performance and an efficient user experience.
## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

For any issues or feature requests, please open an issue on the [GitHub repository](https://github.com/AlexanderIsaac/rebel-communication-decoder/issues).
