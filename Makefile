
.PHONY: build
build:
	go build -o ./bin/grpc-go-redact .

.PHONY: run
run: build
	./bin/grpc-go-redact -input ./test/base/test.pb.go -output ./test/output.pb.go


.PHONY: test
test: run
	go test ./...