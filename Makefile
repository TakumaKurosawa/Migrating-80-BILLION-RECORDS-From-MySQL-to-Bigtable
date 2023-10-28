.DEFAULT_GOAL := help

##### Data prepare #####

.PHONY: data/prepare
data/prepare: ## prepare data for test ## make data/prepare
	cd preparer && go run main.go model.go

##### Go benchmark test #####

.PHONY: bench/basic-migration
bench/basic-migration: ## benchmark test for basic-migration ## make bench/basic-migration
	cd basic-migration && go test -bench . -benchmem -benchtime 6x

.PHONY: bench/shard-concurrency
bench/shard-concurrency: ## benchmark test for shard-concurrency ## make bench/shard-concurrency
	cd shard-concurrency && go test -bench . -benchmem -benchtime 6x

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
