# DevOps Lecture Project
This is a project for the lecture "Einführung in DevOps, Continuous Delivery Tools und Mindset" at the DHBW.

Luxify is a simple clothes webshop developed with the Go programming language.

## Building and running the application with Docker

### Build the Docker Image
Navigate to the root directory of the application and run the following command:
```bash
docker build -t shop-backend .
```

### Run the Application with Docker
After building the image, run the application using Docker with the following command:
```bash
docker run -p  8080:8080 shop-backend
```

### Accessing the Application
Once the container is running, you can access the REST API by navigating to ``http://localhost:8080``.