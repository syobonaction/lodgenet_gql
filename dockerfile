# syntax=docker/dockerfile:1

FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. 
COPY . ./

RUN ls

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-lodgenet-gql

# Expose the port
EXPOSE 8080

# Run
CMD ["/docker-lodgenet-gql"]