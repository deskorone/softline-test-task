run:
	docker run -p 8080:8080 --name server_test softline-test-task
build:
	docker build .  -t softline-test-task
connectDb:
	docker exec -it docker_db psql -U psql authorization_db
up:
	docker-compose up --build
gen:
	mockgen  -source=internal/service/service.go  -package=mock_service -destination=mocks/mock_service.go softline-test-taskinternal/service/service.go
test:
	go test ./internal/...
