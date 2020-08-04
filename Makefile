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
