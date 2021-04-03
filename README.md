# gacha-api
## 💡 Overview
This repository is TechTrain mission Golang GachaAPI.<br>
It can be run locally with docker-compose.<br>
The production will be operated using AWS ECS, AWS RDS (Aurora), etc...

## 🚀 Running
1. You run the following command.
    ```
    << Env files. >>
    ### api/app
    $ cp .env{.sample,}
    ### api/db
    $ cp .env.db{.sample,}

    << Node.js >>
    ### ui/gacha
    $ npm install
    $ npm run build
    $ npm run start
    ```
2. Please rewrite ".env" and ".env.db" as needed.

## 🌱 API EndPoint
```
### /user/create
$ curl -X POST "Content-Type: application/json" -d '{"name": "KobaFumi"}'  localhost:8080/user/create

### /user/get
$ curl -X GET -H "x-token: abc" -H "Content-Type: application/json" -d '{"id": "1"}' localhost:8080/user/get
$ curl -X GET -H "x-token: abc" -H "Content-Type: application/json" localhost:8080/user/get

### /user/update
$ curl -X PUT -H "x-token: abc" -H "Content-Type: application/json" -d '{"name" : "KobayashiFumiaki"}' localhost:8080/user/update

### /gacha/draw
$ curl -X GET -H "x-token: abc" -H "Content-Type: application/json" -d '{"count": 10}' localhost:8080/gacha/draw

###/character/list
$ curl -X GET -H "x-token: abc" -H "Content-Type: application/json" localhost:8080/character/list
```

## 🦆 Init DB env
- `.env`
  ```
  DB_HOST=db
  DB_USER=user
  DB_PASSWORD=password
  DB_DATABASE=go_db
  ```
- `.env.db`
  ```
  MYSQL_ROOT_PASSWORD=password
  MYSQL_DATABASE=go_db
  MYSQL_USER=user
  MYSQL_PASSWORD=password
  ```

## 📝 make-command specifications
```
# make all
$ docker-compose build --no-cache && docker-compose up -d

# make build
$ docker-compose build --no-cache

# make up
$ docker-compose up -d

# make stop
$ docker-compose up -d

# make down
$ docker-compose down

# make app
$ docker-compose exec app sh

# make db 
$ docker-compose exec db bash

# make db/mysql
$ docker-compose exec db bash
> mysql -u user -ppassword

# make log
$ docker-compose logs -f

# make logs/app
$ docker-compose logs -f app

# make logs/db
$ docker-compose logs -f db
```
