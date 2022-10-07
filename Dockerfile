# Start from golang base image
FROM golang:latest as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Thomas Ngo <tuanit168@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .       

# Expose port 8080 to the outside world
EXPOSE 5000

# Install CompileDaemon which is used for hot reload each time a file is changed
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

# The ENTRYPOINT defines the command that will be ran when the container starts up
# The "go build" command here build from the current directory
# We will also execute the binary so that the server starts up. CompileDaemon handles the rest - anytime any .go file changes in the directory
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ." -command="./forum"

#Command to run the executable
CMD ["./main"]