
GEN=cmd/swagger/swagger.go

DIR=examples/tutorials/todo-list/cli

SRC=${DIR}/code

generate:
	mkdir -p ${SRC}
	go run ${GEN} generate cli --target=${SRC} --spec=${DIR}/swagger.yml

clean:
	rm -rf ${SRC}

run:
	go run ${SRC}/cmd/cli/main.go --hostname localhost:12345 --x-todolist-token "example token" todos addOne --item.description "hi"
	
run-get:
	go run ${SRC}/cmd/cli/main.go --hostname localhost:12345 --x-todolist-token "example token" todos findTodos

run-update:
	go run ${SRC}/cmd/cli/main.go --hostname localhost:12345 --x-todolist-token "example token" Todos UpdateOne --id 1 --ItemCompleted true --ItemDescription "done"
#--hostname hello --scheme https
#AddOne
#DestroyOne
curl:
	curl http://localhost:12345/

run-server:
	go run examples/auto-configure/cmd/a-to-do-list-application-server/main.go --port=12345