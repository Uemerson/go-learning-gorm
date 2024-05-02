# About
Basic Go application for learning GORM ORM.

---

See the GORM implementations [here](./user/repository/)

# Instructions

This project uses [Docker and Docker Compose, so it's necessary to install them.](https://docs.docker.com/compose/install/)

After installation, rename the files `.env.example` to `.env`. The file containing environment variables with default values.

- `.env`: It's used to set the environment variables used by the application.

Afterwards, you can bring up the containers using the bash script below:

```
$ bash up.dev.v2.sh
```

or if you don't want to use the bash script:

```
$ docker compose -f docker-compose.dev.yml --env-file .env.dev up -d --build --remove-orphans
```

Now you can run:

```
$ go run main.go
```
