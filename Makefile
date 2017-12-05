GO ?= go
GOPATH := $(CURDIR)

game:
	@echo "----start build weiqi game server----"
	@pwd
	@rm -rf bin/server
	@export GOPATH=$(CURDIR)
	@go build src/main.go
	@mv main server
	@cp server ./bin/
	@rm -rf server
make_run:
	@make game
	@./bin/server
linux:
	@echo "----start build weiqi game server----"
	@pwd
	@rm -rf bin/linux/server
	@export GOPATH=$(CURDIR)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build src/main.go
	@mv main server
	@cp server ./bin/linux/
	@rm -rf server
run_linux:
	@make linux
	@./bin/linux/server