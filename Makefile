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

upprodd:
	docker-compose -f docker-compose.prod.yml -p jocprod --env-file .env.prod up -d

downprod:
	docker-compose -f docker-compose.prod.yml -p jocprod --env-file .env.prod down

reupprod:
	docker-compose -f docker-compose.prod.yml -p jocprod --env-file .env.prod up --build

cleanprod:
	docker-compose -f docker-compose.prod.yml -p jocprod --env-file .env.prod down 
	docker volume prune
