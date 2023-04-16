# hack-nu-umag
Code is inside Umag folder
##  Running the Application with Docker
This readme provides step-by-step instructions on how to run the application using Docker.
##  Prerequisites
Before proceeding, ensure that you have the following:
    Docker installed on your local machine. You can download and install Docker from their official website: https://www.docker.com/get-started
##  Clone the Repository
First, clone the repository to your local machine using the following command:

    git clone https://github.com/gyraseisfortorsion/hack-nu-umag
## Building the MySQL Image
To build the MySQL image, run the following command from the root directory of the project:

    docker build -t my-mysql-image .

This will create a Docker image named my-mysql-image based on the Dockerfile in the root directory of the project.
## Running the MySQL Container
To run the MySQL container, use the following command:

    docker run -d --name my-mysql-cont -p 3306:3306 my-mysql-image

## Building the Application Image
To build the application image, run the following command from the root directory of the project:

    docker build -t back ./Umag

## Running the Application Container

    docker run -d --name backend1 -p 8080:8080 back
    
