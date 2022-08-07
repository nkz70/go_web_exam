ps:
	docker-compose ps
up:
	docker-compose up -d
down:
	docker-compose down
reboot: down up

exec:
	docker-compose exec ${name} sh
rebuild:
	docker-compose build --no-cache
logs:
	docker-compose logs ${name}
genmock:
	genmock -source=${source} destination=${mock} -package=${pkg}
