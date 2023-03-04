.PHONY: build clean deploy

build:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/game_calender cmd/game/calender/main.go

clean:
	rm -rf ./bin ./vendor

deploy: clean build
	sls deploy --verbose
