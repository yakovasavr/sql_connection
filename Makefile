build:
	go build -o /sql_connection ./

.DEFAULT_GOAL := build
.PHONY: build