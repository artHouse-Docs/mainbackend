FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# Download all the dependencies
RUN go mod download

# Build the Go application
RUN go build ./cmd/arthouse-dock/

# Expose the port that the application will run on
EXPOSE 8081

# Command to run the application
CMD ["./arthouse-dock"]