# The Dockerfile is a text file that contains the instructions 
# needed to create a new container image. These instructions include identification of 
# an existing image to be used as a base, commands to be run during the image creation process, 
# and a command that will run when new instances of the container image are deployed.

# Define what docker image will be used
FROM golang:alpine

# Create a directory
RUN mkdir /files

# Copy and paste a file
COPY hw.go /files

# Change working directory
WORKDIR /files

# Create a binary executable from the Go source file
RUN go build -o /files/hw hw.go

# Entrypoint specifies the path of the binary file when running Docker image
ENTRYPOINT ["/files/hw"]
