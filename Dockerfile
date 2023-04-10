# Stage 1: Build binary
FROM golang:alpine AS build
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# Stage 2: Final image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /app/app .
RUN chmod 777 ./app
EXPOSE 8080
CMD ["./app"]

