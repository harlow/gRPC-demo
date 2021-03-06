#!/usr/bin/env sh

# compile the protos
default:
	for f in proto/*.proto; do \
		echo compiled: $$f; \
		protoc --go_out=plugins=grpc:. $$f; \
	done
