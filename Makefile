


.PHONY: build
build:
	go build -o ./bin/grpc-go-redact .

.PHONY: test
test: build
	./bin/grpc-go-redact -input ./test/test.pb.go -output ./bin/output.go