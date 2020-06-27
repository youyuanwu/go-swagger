

GEN=cmd/swagger/swagger.go

DIR=examples/yy-todo-list

SRC=${DIR}/code

generate:
	mkdir -p ${SRC}
	go run ${GEN} generate server --target=${SRC} --spec=${DIR}/swagger.yml

clean:
	rm -rf ${SRC}

run:
	go run ${SRC}/cmd/simple-to-do-list-api-server/main.go --scheme http --host localhost --port 12345



# curl --header 'x-todolist-token:dummy' -X GET http://127.0.0.1:12345