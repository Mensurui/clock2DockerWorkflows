# Build stage
FROM golang:1.23-bookworm as build
WORKDIR /app

COPY go.mod ./
COPY . .

# Ensure static binary compilation
RUN CGO_ENABLED=0 GOOS=linux go build -o clock2

# Final stage
FROM scratch
COPY --from=build /app/clock2 /clock2

CMD ["/clock2"]
