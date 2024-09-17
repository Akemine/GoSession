# GoSession

Ce projet est une API REST en Golang pour un système d'authentification et de gestion de profils utilisateurs. Il offre des fonctionnalités de base telles que la connexion, la déconnexion, l'affichage de profil, ainsi que des exemples de routes protégées et publiques.

## Prérequis
- PostgreSQL
- Go
- Gorilla/session
- Bcrypt
- CORS

# Routes

## Login
Cette route permet à un utilisateur de se connecter.

Méthode : POST
URL : /login
Corps de la requête (JSON) :
{
    "email": "utilisateur@exemple.com",
    "password": "motdepasse123"
}

Exemple de réponse en cas de succès (Status 200 OK) :
"Connexion réussie"

Exemple de réponse en cas d'échec (Status 401 Unauthorized) :
"Identifiants invalides"

## Logout
Cette route permet à un utilisateur connecté de se déconnecter.

Méthode : POST
URL : /logout

Exemple de réponse en cas de succès (Status 200 OK) :
"Déconnexion réussie"

Exemple de réponse en cas d'échec (Status 400 Bad Request) :
"Vous n'êtes pas connecté"

## Profile
Cette route permet à un utilisateur connecté de voir son profil.

Méthode : GET
URL : /profile

Exemple de réponse en cas de succès (Status 200 OK) :
{
    "id": 1,
    "username": "JohnDoe",
    "email": "utilisateur@exemple.com",
    "type": "standard",
    "created_at": "2023-09-17T15:30:45Z"
}

Exemple de réponse en cas d'échec (Status 401 Unauthorized) :
"Vous n'êtes pas autorisé à accéder à cette ressource"

## Protected
Cette route est un exemple de route protégée qui nécessite une authentification.

Méthode : GET
URL : /protected

Exemple de réponse si l'utilisateur est connecté (Status 200 OK) :
"Vous êtes connecté"

Exemple de réponse si l'utilisateur n'est pas connecté (Status 401 Unauthorized) :
"Vous n'êtes pas connecté"

## Public
Cette route est un exemple de route publique qui ne nécessite pas d'authentification.

Méthode : GET
URL : /public

Exemple de réponse en cas de succès (Status 200 OK) :
"Vous êtes sur la page publique"

## Installation

1. Clonez le dépôt :
   ```
   git clone https://github.com/Akemine/GoSession
   ```

2. Installez les dépendances :
   ```
   go mod tidy
   ```

3. Configurez votre base de données PostgreSQL et mettez à jour les informations de connexion dans le fichier de configuration.

4. Lancez l'application :
   ```
   go run cmd/main.go
   ```

## Utilisation

L'API est accessible à l'adresse `http://localhost:8080`. Utilisez un client HTTP comme cURL ou Postman pour interagir avec les différentes routes.

## Sécurité

Cette API utilise des sessions pour l'authentification. Assurez-vous d'utiliser HTTPS en production pour sécuriser les communications.

## Contribution

Les contributions sont les bienvenues ! Veuillez consulter le fichier CONTRIBUTING.md pour plus d'informations.

## Licence

Ce projet est sous licence [insérer le type de licence ici]. Voir le fichier LICENSE pour plus de détails.


## Script SQL pour la création de la table "users"

Voici le script SQL pour créer la table "users" dans votre base de données PostgreSQL :

CREATE TABLE USERS (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    type BOOLEAN NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);





