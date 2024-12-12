FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN go build -o ext_authz .
EXPOSE 9000
CMD ["./ext_authz"]