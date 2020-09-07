

GEN=cmd/swagger/swagger.go

DIR=fixtures/enhancements/936

SRC=${DIR}/code

generate:
	mkdir -p ${SRC}
	#go run ${GEN} generate model --name=object_1 --name=object_2 --name=object_6 --target=${SRC} --spec=${DIR}/fixture-936.yml
	go run ${GEN} generate server --target=${SRC} --spec=${DIR}/fixture-936.yml
clean:
	rm -rf ${SRC}

run:
	go run ${SRC}/cmd/simple-to-do-list-api-server/main.go --scheme http --host localhost --port 12345


test:
	go test -v --run=^TestMoreModelValidations/codegen-fixture-936 ./generator/...
# curl --header 'x-todolist-token:dummy' -X GET http://127.0.0.1:12345 