unit_test:
	go test ./routes

pact_test:
	go test ./main_test.go

docker_build:
	docker build . -t service

docker_build_test:
	docker build . -t todo_api_test --target=test

docker_run:
	docker run --publish 4000:4000 service