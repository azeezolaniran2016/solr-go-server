.PHONY: dev/start
dev/start:
	-set -a && . ./env/dev.env && go run main.go