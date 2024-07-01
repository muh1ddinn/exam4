CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

gen-proto-module:
	sudo rm -rf ${CURRENT_DIR}/genproto
	./scripts/gen_proto.sh ${CURRENT_DIR}

migration-up:
	migrate -database 'postgres://muhiddin:1@localhost:5432/manager_crm?sslmode=disable' -path migration up

		
migration-down:
	migrate -database 'postgres://muhiddin:1@localhost:5432/manager_crm?sslmode=disable' -path ./migration down

	
	