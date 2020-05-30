up:
	docker-compose up --build

down:
	docker-compose down

logs:
	docker-compose logs -f

prune:
	docker volume prune

clean:
	docker system prune -f
