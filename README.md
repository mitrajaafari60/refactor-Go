# Refactored Codebase - Shopping Cart Manager

Welcome to the refactored version of the Shopping Cart Manager codebase! This repository was initially created to test coding skills by refactoring a simple tax calculator. The tax calculator was later replaced with a shopping cart manager, and your task was to refactor the code while maintaining the functionality of the application. Below, you'll find information on the updated codebase and how to navigate and run the application.

# Overview

The Shopping Cart Manager is a straightforward web application that allows users to interact with a shopping cart. Three main functionalities are provided:

#### Show Form: Display the form to add or remove products from the cart.

#### Add Products: Add products to the shopping cart.

#### Remove Products: Remove items from the shopping cart.

## Changes Made

Throughout the refactoring process, several improvements have been made to enhance the codebase:

 Cleaner Code: The code has been refactored for improved readability, maintainability, and adherence to best practices.

 Test Coverage: Comprehensive testing has been added to ensure the reliability of the code.

Structural Improvements: The project structure has been organized for better clarity and separation of concerns.

# Getting Started

Follow these steps to run the application:

Clone the Repository: Clone this repository to your local machine.

Build Docker Containers: If you don't have MySQL installed locally, Docker is used to set up containers. Navigate to the docker directory and run:

bash
Copy code
docker-compose up -d --build
Run the Application: Once the containers are ready, navigate to cmd/web-api and execute:

bash
Copy code
go run main.go
Access the Application: Open your browser and visit http://localhost:8088/ to interact with the application.

Testing
The refactored code includes a comprehensive set of tests. To run the tests, execute:

bash
Copy code
go test ./...