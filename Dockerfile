# Use an official Go image as the base image
FROM golang:1.21.1

# Set the working directory in the container to /go/src/event-router
WORKDIR /app

# Copy the main.go file into the container
COPY . .

# Install dependencies
RUN go mod download
# Build the Go application
RUN ls -a
RUN cd src/app && go build -o /event-router-app

# Expose port 8080
EXPOSE 8080

# Start the Go application when the container starts
CMD ["/event-router-app"]
