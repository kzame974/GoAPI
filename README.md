# GoAPI

Ce projet est un script pour envoyer des données hexadécimales sous forme de chaîne de caractères dans le corps d'une requête POST. Il utilise le framework Gin pour la gestion des routes et InfluxDB pour stocker les données.



## Prérequis
-Avoir installé Docker sur votre machine.
```bash
https://www.docker.com/
```
-Avoir un fichier .env à la racine du projet.


## Comment lancer le projet
Clonez ce dépôt sur votre ordinateur

```bash
git clone https://github.com/kzame974/GoAPI.git
```
Placez votre fichier .env à la racine du projet
Construisez l'image Docker à l'aide de la commande suivante dans le répertoire racine de votre projet:


```bash
docker build -t goapi:latest .
```
Lancez le container à l'aide de la commande suivante :

```bash
docker run -p 8083:8083 goapi:latest
```
A la racine lancez la commande pour excécuter main:
```bash
go run main.go

```

Dans un nouveau terminal, accédez au répertoire src/script et exécutez la commande
```bash
go run scriptPostRequest.go

```

## Remarques

Le script envoie des données de façon aléatoire, vous pouvez les modifier dans le fichier scriptPostRequest.go si nécessaire.

Le script envoie des données à l'adresse http://localhost:8083/sensors. Assurez-vous que l'adresse est correcte et que le service est bien en cours d'exécution avant de lancer le script.