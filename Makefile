
test:
	go test -short -v `go list ./... | grep -v ./docker` 
test-integration:
	go test github.com/peknur/nginx-unit-sdk/unit 
docker-run:
	docker run -p 8080:8080 -d --rm unit:test
docker-build:
	docker build -t unit:test -f docker/Dockerfile docker/	
docker-init: docker-build docker-run
