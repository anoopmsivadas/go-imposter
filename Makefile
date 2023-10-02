build:
	go mod tidy
	go mod vendor
	go build -mod=vendor -v -o build/imposter cmd/imposter/main.go

clean:
	rm -rf build