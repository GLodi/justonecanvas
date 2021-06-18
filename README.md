# justonecanvas

If you use Firefox, or any canvas-blocking mechanism, you need to accept canvas permissions 
in your browser BEFORE you make a move, or it may lag a lot.

Website no longer online!

[Blog post](https://giuliolodi.dev/how-i-made-justonecanvas)

![image](https://github.com/GLodi/justonecanvas/blob/master/picture.png)

This is essentially a copy of Reddit's 2017 [r/place](https://redditblog.com/2017/04/13/how-we-built-rplace/). Just tinier
(65x65) and designed to work on a 5$ Digital Ocean droplet.

All users connected to justonecanvas actually share the same canvas, and the backend
updates all players of each other's moves thanks to WebSockets.

This project follows a combination of [this](https://github.com/L04DB4L4NC3R/clean-architecture-sample) and [this](https://github.com/AkbaraliShaikh/denti) implementations
of Clean Architecture

Websockets in Golang implements this [chat](https://github.com/gorilla/websocket/tree/master/examples/chat) example, while frontend
follows [this](https://dev.to/finallynero/using-websockets-in-react-4fkp)

Just FYI: frontend kinda sucks. Being my first experience doing frontend, it's very basic and
not at all optimized for mobile.

## development

A complete dev environment is available under docker-compose. Make sure to have docker 19.03 and docker-compose 1.26.2

Check Makefile or just `make up` under root.
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

`npx create-react-app client --use-npm --typescript --scripts-version=react-scripts@3.4.0`

Runs on 3000

```
npm install
npm start
```

### load testing

You can do load testing thanks to Artillery:
	- `npm install -g artillery`
	- cd in testing directory
	- `artillery run loadtest.yml`

## production

Check makefile or `make upprod` under root folder.

The production environment makes two assumptions:
 - There must be a .env.prod under root with your variables
 - It must find let's encrypt's certs under `/etc/letsencrypt/live/justonecanvas.live`


## profiling

Thanks to pprof package, you can:

`wget http://localhost:8080/debug/pprof/trace\?seconds\=5` to profile 5 seconds of execution.

`go tool trace 'trace?seconds=5'` to analyze.

Remember to add `pprof.Register(g)` in main.go.
