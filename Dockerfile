# Use a smaller base image for the build stage
FROM golang:1.22.2-alpine as builder

WORKDIR /src
COPY . .

# Enable Go modules
ENV GO111MODULE=on

RUN go mod tidy
# Compile the action
RUN go build -o /bin/action

# Use a smaller base image for the final container
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the compiled Go program from the builder stage
COPY --from=builder /bin/action /usr/local/bin/hugo-notion

# Copy the entrypoint script
COPY entrypoint.sh  /usr/local/bin/entrypoint.sh

# Make the entrypoint script executable
RUN chmod +x  /usr/local/bin/entrypoint.sh
RUN ls -la  /usr/local/bin/
# Set entrypoint
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]