SHELL := /bin/bash

.PHONY: dev/start
dev/start:
	@docker-compose up --remove-orphans --build