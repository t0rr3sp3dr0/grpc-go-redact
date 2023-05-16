all: clean build test
.PHONY: all

clean:
	rm -fR ./bin/*
	rm -fR ./grpc-go-redact
	rm -fR ./test/input.pb.go
	rm -fR ./test/input.txt
	rm -fR ./test/output.pb.go
	rm -fR ./test/output.txt
.PHONY: clean

build: bin/grpc-go-redact
.PHONY: build

test: test/output.pb.go
	go test ./...
.PHONY: test

bin/grpc-go-redact:
	go build -o ./bin/grpc-go-redact .

test/input.txt:
	sed -E -e 's/^package internal$$/package test/g' -e '/^(\/\*|\*\/)$$/d' ./generator/internal/stringfunc.pb.go > ./test/input.txt

test/output.txt: bin/grpc-go-redact test/input.txt
	cp ./test/input.txt ./test/output.txt
	for _ in $$(seq 10); do ./bin/grpc-go-redact -input ./test/output.txt -output ./test/output.txt; done

test/output.pb.go: test/output.txt
	cp ./test/output.txt ./test/output.pb.go
