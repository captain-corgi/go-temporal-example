# Temporal server
.PHONY: start-temporal
start-temporal:
	temporal server start-dev --ui-port 18080 --db-filename ./temporal/db/persist_db_01.db

# Demo
.PHONY: clone-demo
clone-demo:
	git clone https://github.com/temporalio/money-transfer-project-template-go
.PHONY: run-demo-workflow
run-demo-workflow:
	go run money-transfer-project-template-go/start/main.go
.PHONY: run-demo-worker
run-demo-worker:
	go run money-transfer-project-template-go/worker/main.go

# Practice
.PHONY: run-workflow
run-workflow:
	go run go-temporal-example-01/cmd/app/main.go
.PHONY: run-worker
run-worker:
	go run go-temporal-example-01/cmd/worker/main.go

