# alphabet-api
Alphabet API checks if all the alphabet exists in a given string.

[![Build Status](https://travis-ci.com/b-entangled/alphabet-api.svg?branch=master)](https://travis-ci.com/b-entangled/alphabet-api)

## Use Case
For a given string as input, potentially having a mixture of upper and lower case, numbers, special characters etc. The API should determine if the string contains at least one of each letter of the alphabet. Return true if all are found and false if not.

## API Endpoints

The Alphabet API contains a single endpoint and a single parameter 'inputstring' to that endpoint.

```
http://localhost:8080/alphabet/validate
```

Parameter
```
inputstring = <string>
```

## Documentation

An OpenAPI documentation for the service is generated using Swagger Inspector

[Open Documentation](https://github.com/b-entangled/alphabet-api/blob/master/alphabet_api_docs.yaml)

## Deployment

The Service can be deployed in Kubernetes Cluster.
AWS provides AWS EKS which is a managed kubernetes service and we can use Docker Hub or AWS ECR to store and use Docker Images.
For local development a Minikube or Docker-Desktop can be used to build and test the service in a kubernetes environment.
[Makefile](https://github.com/b-entangled/alphabet-api/blob/master/Makefile) contains the set of commands to build and deploy the service

The deployment involves,

** 1> Creation of Docker File **
DockerFile is a collection of commands to build docker images.
[DockerFile](https://github.com/b-entangled/alphabet-api/blob/master/Dockerfile)

** 2> Docker Image Repository **
To share, store and use Docker images requires an image repositories.
We can either use DockerHub or AWS ECR.
AWS ECR involves a creation of Repository using AWS CLI or AWS Console.
The Docker images are built locally or in a CI Server (Travis-CI) and can be pushed into ECR Registry.

Building Images

Makefiles provides a command called ```build-alphabet``` to build Docker images.
[build-alphabet](https://github.com/b-entangled/alphabet-api/blob/master/Makefile#L13-L15)

Run ```make build-alphabet``` to build service

The built docker images can be pushed into ECR using

```
docker push 123456-dkr.ecr.us-west-2.amazonaws.com/bentangled/alphabet:latest
```

** 3> Creation of EKS Cluster **

The creation of EKS Cluster involves creation of AWS VPC, Subnets, Internet Gateway,
AWS EKS Master and Worker Nodes etc.

We can either use AWS Cloudformation or Terraform to create EKS Cluster.
A generation of Kubeconfig file is necessary to connect to cluster.

** 4> Creation and Deployment of Kubernetes Manifests **

Once the cluster is ready we can create and deploy kubernetes deployment and service inside the cluster.

The kubernetes manifest contains a Deployment and a Service.
[Kubernetes Manifests](https://github.com/b-entangled/alphabet-api/blob/master/alphabet.yaml)

We can deploy the service using the command ```make deploy-alphabet```
The command creates a deployment and a service for the API.
[deploy-alphabet](https://github.com/b-entangled/alphabet-api/blob/master/alphabet.yaml)

** 5> Accessing Service **

For testing the service can be accessed by port-forwarding the service

```kubectl port-forward svc/alphabet 8080:80```

To expose the service externally involves creation of service type Loadbalancer.
