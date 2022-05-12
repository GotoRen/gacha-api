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
CONTAINER ID   IMAGE           COMMAND                  CREATED          STATUS          PORTS                                NAMES
ea5ca2f50bf2   gacha-api_db    "docker-entrypoint.s…"   58 seconds ago   Up 53 seconds   33060/tcp, 0.0.0.0:13306->3306/tcp   gacha-api_db
fa3719719d37   gacha-api/app   "./main"                 58 seconds ago   Up 54 seconds   0.0.0.0:8080->8080/tcp               gacha-api_app

=== * 作成されるDockerイメージ * ===
$ docker images
REPOSITORY                           TAG        IMAGE ID       CREATED         SIZE
gacha-api/app                        latest     e58639cff39e   2 minutes ago   13.5MB
gacha-api_db                         latest     8c6bc13c1a11   6 days ago      546MB

=== * 作成されるDockerネットワーク * ===
$ docker network ls
NETWORK ID     NAME             DRIVER    SCOPE
de4ec0aef8b7   gacha-api_link   bridge    local
```
