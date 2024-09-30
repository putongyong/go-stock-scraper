# Step 1: Use the official Golang image to build the app
FROM golang:1.23-alpine AS builder

# Step 2: Set the current working directory inside the container
WORKDIR /app

# Step 3: Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Step 4: Download the Go module dependencies
RUN go mod download

# Step 5: Copy the entire project into the container
COPY . .

# Step 6: Build the Go app and generate a binary called 'app'
RUN go build -o app ./main.go

# Step 7: Use a small image to run the compiled Go binary
FROM alpine:latest

# Step 8: Set the current working directory inside the container
WORKDIR /root/

# Step 9: Copy the built binary from the builder stage
COPY --from=builder /app/app .

# Step 10: Copy the tickers.txt file to the container
COPY tickers.txt .

# Step 11: Expose port if needed (optional)
# EXPOSE 8080

# Step 12: Command to run the binary
CMD ["./app"]
