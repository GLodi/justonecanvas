# justonecanvas

This project follows [this](https://github.com/L04DB4L4NC3R/clean-architecture-sample) implementation
of Clean Architecture

## development

A complete dev environment is available under docker-compose. Check Makefile or just `docker-compose up` under root.
Both client and server are setup for hot-reloading.

If you don't want to use docker, then:

### server

Runs on 8080.

```
go mod download
go run ./main.go
```

### client

Because of a problem regarding Docker and create-react-app, the project was created with:

`npx create-react-app client2 --use-npm --typescript --scripts-version=react-scripts@3.4.0`

Runs on 3000

```
npm install
npm start
```

## production

Create a .env.prod under root with your variables and
`docker-compose -f docker-compose.prod.yml up` under root folder.

## profiling

Thanks to pprof package, you can:

`wget http://localhost:8080/debug/pprof/trace\?seconds\=5` to profile 5 seconds of execution.

`go tool trace 'trace?seconds=5'` to analyze.
