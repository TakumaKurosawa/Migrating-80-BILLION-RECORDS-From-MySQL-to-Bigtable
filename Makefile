.DEFAULT_GOAL := help

##### Setup #####

.PHONY: setup
setup: ## setup for test ## make setup
	docker compose up -d
	make data/prepare

##### Data prepare #####

.PHONY: data/prepare
data/prepare: ## prepare data for test ## make data/prepare
	cd preparer && go run main.go model.go

##### Go benchmark test #####

.PHONY: bench/basic-migration
bench/basic-migration: ## benchmark test for basic-migration ## make bench/basic-migration
	cd basic-migration && go test -bench . -benchmem -benchtime 6x -count 6 > ../basic-migration.out

.PHONY: bench/shard-concurrency
bench/shard-concurrency: ## benchmark test for shard-concurrency ## make bench/shard-concurrency
	cd shard-concurrency && go test -bench . -benchmem -benchtime 6x -count 6 > ../shard-concurrency.out

.PHONY: bench/worker-pattern
bench/worker-pattern: ## benchmark test for worker-pattern ## make bench/worker-pattern
	cd worker-pattern && go test -bench . -benchmem -benchtime 6x -count 6 > ../worker-pattern.out

.PHONY: bench/all
bench/all: ## benchmark test for all ## make bench/all
	make bench/basic-migration
	@echo ""
	make bench/shard-concurrency
	@echo ""
	make bench/worker-pattern

##### HELP #####

.PHONY: help
help: ## Display this help screen ## make or make help
	@echo ""
	@echo "Usage: make SUB_COMMAND argument_name=argument_value"
	@echo ""
	@echo "Command list:"
	@echo ""
	@printf "\033[36m%-30s\033[0m %-50s %s\n" "[Sub command]" "[Description]" "[Example]"
	@grep -E '^[/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | perl -pe 's%^([/a-zA-Z_-]+):.*?(##)%$$1 $$2%' | awk -F " *?## *?" '{printf "\033[36m%-30s\033[0m %-50s %s\n", $$1, $$2, $$3}'
