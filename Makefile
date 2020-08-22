up:
	docker-compose up

upd:
	docker-compose up -d

reup:
	docker-compose up --build

down:
	docker-compose down

logs:
	docker-compose logs -f

prune:
	docker volume prune

clean:
	docker-compose down
	docker volume prune

upprod:
	docker-compose -f docker-compose.prod.yml -p jocprod --env-file .env.prod up 

reupprod:
	docker-compose -f docker-compose.prod.yml -p jocprod --env-file .env.prod up --build

reupprodserver:
	docker-compose -f docker-compose.prod.yml -p jocprod --env-file .env.prod up --no-deps --build server

reupprodclient:
	docker-compose -f docker-compose.prod.yml -p jocprod --env-file .env.prod up --no-deps --build client

cleanprod:
	docker-compose -f docker-compose.prod.yml -p jocprod --env-file .env.prod down 
	docker volume prune
