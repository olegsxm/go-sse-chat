clean:
	clean: CRUFT=$(shell find . -name '._*')
	clean: ; rm -f $(CRUFT)

prepare:
	 go get github.com/mailru/easyjson && go install github.com/mailru/easyjson/...@latest
	 go install github.com/swaggo/swag/cmd/swag@latest

swagger:
	swag init --parseInternal --generalInfo cmd/sse/main.go

json:
	easyjson ./internal/models/*.go