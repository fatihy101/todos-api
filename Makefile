docker_build_test:
	docker build . -t todo_api_test --target=test

test: docker_build_test
	docker-compose down
	docker-compose up -d
	docker-compose exec -T http go test ./...
	docker-compose down


run_app: docker_build docker_run

docker_build:
	docker build . -t todos_api

docker_run:
	docker run --publish 4000:4000 todos_api