.PHONY: clean
	clean: CRUFT=$(shell find . -name '._*')
	clean: ; rm -f $(CRUFT)

swagger:
	swag init --parseInternal --generalInfo cmd/sse/main.go