#BUILD STAGE
FROM golang:1.24-alpine3.21 AS builder
WORKDIR /app
COPY . .
#build
RUN go build -o main main.go

#RUN STAGE
FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

#publish port 
EXPOSE 9090

#commands to run when container starts
CMD ["/app/main"]
