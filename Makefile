.PHONY: deps clean live

deps:
	go get -u ./...

clean:
	rm -rf ./output

live: clean
	GOOS=linux go build -o output/maze_api ./main.go

deploy: live
	scp output/maze_api skenderi:/lib/maze/maze_api
