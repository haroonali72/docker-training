# This command defines the base image
FROM golang:1.18
# This command create a directory, in our case app directory will be created
WORKDIR /app
# This command copies the code to the directory, created with WORKDIR
COPY .. .
# This command is used to run any shell command, and this is run while creating the image.
RUN go mod download
# This will set the environment variable for the container
# currently we have defined default value for ENV but we can override
# this value while running the container using docker run -e VAR1.
# One more thing to remember environment variable only used while running the container not
# while creating the image.
ENV VAR1=defaultValue
# CMD will be used to run the container, but we can override CMD command while running the container.
# e.g docker run <container> go run main.go
#CMD ["go", "run", "main.go"]

# ENTRYPOINT used to run container same as CMD but we can't override ENTRYPOINT arguments instead
# we can append new arguments while running container. e.g docker run <container> main.go, under the hood it will run
# this command "go run main.go"
ENTRYPOINT ["go", "run"]