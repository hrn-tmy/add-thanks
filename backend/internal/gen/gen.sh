#!/bin/bash

oapi-codegen -generate "types" -package gen ./internal/gen/gen.yaml > ./internal/gen/gen.types.go
oapi-codegen -generate "server" -package gen ./internal/gen/gen.yaml > ./internal/gen/gen.server.go