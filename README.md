# Microservices in Golang

### What is microservices?

Microservices is an architectural approach to building software applications as a collection of small, loosely coupled services that can be developed, deployed, and scaled independently.
Each microservice focuses on a specific business capability.

### Illustration of Microservices

<img src="docs/go.microservices.png" height=250 width=500>

### Why do we use microservices?

- **Scalability:** Microservices allow applications to scale more efficiently. Each microservice can be independently scaled based on its specific needs, which enables better resource allocation and utilization.
- **Flexibility and Agility:** Microservices promote flexibility and agility in software development. Since each microservice is developed and deployed independently, it enables teams to work autonomously and make changes or updates to a specific service without impacting the entire system.
- **Modularity and Reusability:** Microservices promote modularity and reusability of code. Services can be developed as independent modules with well-defined interfaces, making it easier to reuse them in different projects or integrate them into new systems.
- **Team Scalability and Autonomy:** Microservices enable teams to work on different services independently, which promotes autonomy and scalability within the development process.

## Drawbacks of microservices.

- **Increased Complexity:** Microservices introduce a higher level of complexity compared to monolithic architectures. Managing multiple independent services requires additional effort in terms of deployment, monitoring, and inter-service communication.
- **Distributed System Challenges:** Microservices are distributed systems, and communication between services typically occurs over the network. This introduces challenges such as network latency, potential points of failure, and the need for robust error handling and retries.
- **Operational Overhead:** Scaling, updating, and ensuring the availability of multiple services also adds to the operational overhead.
- **Development and Testing:** Each service needs to be developed, tested, and deployed independently, which requires additional coordination and setup. End-to-end testing across multiple services can be challenging due to the distributed nature of the system.

### How can the project be structured for the microservices architecture?

When structuring a project with a microservices architecture, it's important to consider scalability, maintainability, and separation of concerns. Here's a suggested structure:

**Root Directory:**

- README.md: Document the project overview, setup instructions, and any relevant information.
  LICENSE: Include the license under which the project is distributed.
- .gitignore: Specify files and directories to be ignored by version control.
- main.go: Entry point of the application.

**Api Gateways:** Create a separate directory for the API gateway, which serves as the entry point for client requests:
`api-gateway/`

- **Routes:** Define the routing logic for forwarding requests to the appropriate microservices.
- **Middlewares:** Implement gateway-specific middleware functions, such as authentication and rate limiting.

**Services:** Create a separate directory for each microservice in the project. For example:

```
user-service/
auth-service/
...
```

Each microservice directory should have its own structure following the principles of a typical project, including:

- **Handlers:** Handle incoming requests, validate data, and interact with the corresponding service.
- **Services:** Implement the business logic and perform the necessary data operations.
- **Repositories:** Interact with the database or any other data storage mechanism.
- **Models:** Define the data models used within the microservice.
- **Routes:** Define the API endpoints and routing logic specific to the microservice.
- **Tests:** Include unit tests and integration tests to ensure the microservice's functionality and reliability.

**Shared Code:**
Create a directory for shared code that can be used by multiple microservices:
` common/`

- **Utilities:** Include utility functions, helpers, and common code snippets.
- **Database:** Define common database configurations and connection pools.
  API Gateway:


**Documentation:** Includes any project-specific documentation, such as API specifications, architectural diagrams, and setup instructions.
Tests and Scripts:

Create a directory for project-wide tests and scripts:
tests/

- **Integration:** Includes integration tests that span multiple microservices.

  **scripts:** Includes utility scripts for common tasks, such as database migration or test data generation.

---

### Call flow of the application

The call flow for the application looks like

- Whenever client performs the request, the api-gateway handles the request and forward the request to the respective microservice.
- The forwarded request checks for the middlewares and if any of the middleware is used for the requests then, the respective middleware is being called.
    - If the request is not verified by the middlewares, it then stops the request.
    - If the request is verified, then it will forward the request.
- Now, the path to the request calls the methods from the `handler` class. The handler is responsible for validating the incoming request data and interact with the corresponding services.
- The `service` class is responsible for the business logic of the particular microservice. So, it performs the data operations, and pass the data to the `repositories.`
- The `repositories` are responsible for interacting with the database or external storage services.
- Now, the data received from the `repositories` is passed back to the `handlers`.
- The `handlers` send the response back to the client.