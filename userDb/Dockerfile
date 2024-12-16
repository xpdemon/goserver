FROM golang:1.23-alpine

WORKDIR /app

# Créer le répertoire pour la base de données
RUN mkdir -p /data

# Installer les dépendances nécessaires pour CGO
RUN apk add --no-cache gcc musl-dev

# Copier les fichiers sources dans le conteneur
COPY . .

# Définir un volume pour la base de données SQLite
VOLUME ["/data"]

# Construire l'application avec CGO activé
RUN CGO_ENABLED=1 go build -o user_db .

# Exposer le port sur lequel l'application écoute
EXPOSE 8085

# Commande pour exécuter l'application
CMD ["./user_db"]