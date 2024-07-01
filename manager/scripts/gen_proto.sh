#!/bin/bash

# Get the current directory from the first argument

CURRENT_DIR=$1

# Delete old generated files

rm -rf "${CURRENT_DIR}/genproto/*"

#  Check that the protoc-gen-go and protoc-gen-go-grpc plugins are installed and are in PATH
if ! command -v protoc-gen-go &> /dev/null || ! command -v protoc-gen-go-grpc &> /dev/null; then
  echo "protoc-gen-go or protoc-gen-go-grpc not found in PATH"
  exit 1
fi

# # Generate code for each subdirectory in the proto folder

for x in $(find ${CURRENT_DIR}/proto/* -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/proto -I /usr/local/include \
         --go_out=${CURRENT_DIR} \
         --go-grpc_out=require_unimplemented_servers=false:${CURRENT_DIR} \
         ${x}/*.proto
done