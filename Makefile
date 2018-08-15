.PHONY: fmt gen

fmt:
	go fmt
	find . -type f -name "*.proto" | xargs clang-format -verbose -style file -i

gen:
	go generate
	protoc --proto_path=. --go_out=. protobuf.proto