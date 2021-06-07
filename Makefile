build:
	env GOOS=linux GOARXH=amd64 go build -ldflags="-s -w" -o bin/server

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose