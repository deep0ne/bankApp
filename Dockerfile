# build stage
FROM golang AS builder
WORKDIR /app
COPY . . 
RUN go build -o main main.go

# Run Stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD [ "/app/main" ]