APP_BIN=build/app
APP_CLI=build/cli

build-app: clean $(APP_BIN)
build-cli: clean $(APP_CLI)
build: clean $(APP_BIN) $(APP_CLI)

$(APP_BIN):
	go build -o $(APP_BIN) ./cmd/app/main.go
$(APP_CLI):
	go build -o $(APP_CLI) ./cmd/cli/main.go

clean:
	rm -rf ./build || true

