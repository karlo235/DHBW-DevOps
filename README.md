# DevOps Lecture Project
This is a project for the lecture "Einführung in DevOps, Continuous Delivery Tools und Mindset" at the DHBW.

Luxify is a simple clothes webshop developed with the Go programming language.

## Building the application with Make

### Build the Application
Navigate to the root directory of the application and run the following command:
```bash
make build
```

### Run the Application
After building the application, it can be run using the following command:
```bash
go run main.go
```

## Building and running the application with Vagrant

### Setting Up a Vagrant Environment
Navigate to the `infra/vagrant` directory of the application and run the following command:
```bash
vagrant up
```
This will create a virtual machine, build the application and run it inside of the machine.

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

## Deploying the application

### Kubernetes
To deploy the application to a Kubernetes cluster, you can use Kubernetes manifests located in the `manifests` folder.

### Terraform
Configuration files for Terraform can be found inside the `infra/terraform` folder. You can use these to provision cloud resources in Azure.

### ArgoCD
Inside the `manifests/argocd-apps` folder, there are configuration files, that you can use if you wish to deploy the application using ArgoCD. This also gives you an opportunity to deploy the kube-prometheus-stack to observe your system.