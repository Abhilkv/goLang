# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build main.go

# Expose the port your application listens on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

# build docker image 
# docker build -t go_gorm_postgress .
# docker run --name my_app_container --link some-postgres-name:postgres_image_name  -p 8080:8080 go_gorm_postgress
