# gacha-api
## 💡 Overview
This repository is TechTrain mission Golang GachaAPI.<br>
It can be run locally with docker-compose.<br>
The production will be operated using AWS ECS, AWS RDS (Aurora), etc...

## 🚀 Running
You run the following command.
### Frontend
```
<< Nuxt.js >>
### ui/gacha
$ npm install
$ npm run build
$ npm run start
```
### Backend
```
<< Env files. >>
### api
$ cp api/.env{.sample,}

### db
$ cp db/.env.db{.sample,}
```
Please rewrite ".env" and ".env.db" as needed.

## 🌱 API EndPoint
```zsh
### /user/create
$ curl -X POST "Content-Type: application/json" -d '{"name": "KobaFumi"}'  localhost:8080/user/create

### /user/login
$ curl -X POST -H "Content-Type: application/json" -d '{"id":1, "name":"RenGoto"}' localhost:8080/user/login

### /user/get
$ curl -X GET -H "x-token: abc" -H "Content-Type: application/json" localhost:8080/user/get

### /user/update
$ curl -X PUT -H "x-token: abc" -H "Content-Type: application/json" -d '{"name" : "KobaKoba"}' localhost:8080/user/update

### /gacha/draw
$ curl -X GET -H "x-token: abc" -H "Content-Type: application/json" 'localhost:8080/gacha/draw?count=10'

### /character/list
$ curl -X GET -H "x-token: abc" -H "Content-Type: application/json" localhost:8080/character/list
```

## 🦆 Init DB env
`.env`
```
DB_HOST=db
DB_USER=user
DB_PASSWORD=password
DB_DATABASE=go_db
```
`.env.db`
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

# make app/api
$ docker-compose exec app sh

# make app/db 
$ docker-compose exec db bash

# make mysql
$ docker-compose exec db bash
> mysql -u user -ppassword

# make log
$ docker-compose logs -f

# make logs/app
$ docker-compose logs -f app

# make logs/db
$ docker-compose logs -f db
```

## 🍬 Created
```
### 確認
=== * 起動するDockerコンテナ * ===
$ docker ps
CONTAINER ID   IMAGE                    COMMAND                  CREATED         STATUS         PORTS                                NAMES
2c03853063ec   gacha-api/db:gacha-v1    "docker-entrypoint.s…"   7 seconds ago   Up 1 second    33060/tcp, 0.0.0.0:13306->3306/tcp   gacha-api_db
cfc8a77dd2de   gacha-api/app:gacha-v1   "./main"                 7 seconds ago   Up 3 seconds   0.0.0.0:8080->8080/tcp               gacha-api_app

=== * 作成されるDockerイメージ * ===
$ docker images
REPOSITORY                           TAG        IMAGE ID       CREATED          SIZE
gacha-api/app                        gacha-v1   c7d45815b55b   29 seconds ago   13.5MB
gacha-api/db                         gacha-v1   8c6bc13c1a11   3 days ago       546MB

=== * 作成されるDockerネットワーク * ===
$ docker network ls
NETWORK ID     NAME                    DRIVER    SCOPE
1cd300b2d326   gacha-api_link          bridge    local
```

## 🚧 DevOps Architecture
![architecture](https://user-images.githubusercontent.com/63791288/113522998-0c822200-95e0-11eb-851a-ee61c69076f1.png)

