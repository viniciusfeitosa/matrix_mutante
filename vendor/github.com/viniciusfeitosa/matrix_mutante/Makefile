.PHONY: install
install:
	@go get

.PHONY: test
test: 
	@go test ./... -v -coverprofile=. -timeout 30s

.PHONY: run
run:
	@go run main.go app.go