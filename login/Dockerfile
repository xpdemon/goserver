FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN go build -o go_login .
EXPOSE 8081
CMD ["./go_login"]