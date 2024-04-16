dev:
	docker-compose up -d

logs:
	docker-compose logs --no-log-prefix -f app

clean:
	docker-compose down

sh:
	docker-compose exec -it app sh

prod:
	docker build -t johanbx:prod -f docker/Dockerfile .
	docker run -v $$(pwd)/dev.db:/app/db/dev.db:rw -e GIN_MODE=release -e SQLITE_URI=./db/dev.db -p 8080:8080 johanbx:prod

test:
	echo $$(pwd)