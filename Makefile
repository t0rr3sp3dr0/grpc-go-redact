


.PHONY: build
build:
	go build -o ./bin/grpc-go-redact .

.PHONY: test
test: build
	./bin/grpc-go-redact -input ./test/base/test.pb.go -output ./test/test.pb.go
	go test ./... -v