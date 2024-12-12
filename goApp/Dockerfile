FROM golang:1.23 AS builder

WORKDIR /app

# Copier le code source
COPY .. .

# Compiler le binaire statiquement
ENV CGO_ENABLED=0
RUN go build -o server main.go

FROM alpine:3.17
WORKDIR /app

# Copier le binaire depuis l'image builder
COPY --from=builder /app/server .

# Donner les droits d'exécution (normalement déjà exécutables, mais par sécurité)
RUN chmod +x server

EXPOSE 8080
CMD ["./server"]