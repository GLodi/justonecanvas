# justonecanvas

## development

A complete dev environment is available under docker-compose. Just `docker-compose up` under root.
Both client and server are setup for hot-reloading.

If you don't want to use docker, then:

### server

Runs on 8080.

```
go mod download
go run ./main.go
```

### client

Runs on 3000

```
npm install
npm start
```

## production

Create a .env.prod under root with your variables and
`docker-compose -f docker-compose.prod.yml up` under root folder.
