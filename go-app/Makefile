test:
	@go test -cover ./... -coverprofile cover.out
	@echo "-------------------------------------------------------------------------------------"
	@go tool cover -func cover.out
	@echo "-------------------------------------------------------------------------------------"

lint:
	@golangci-lint run

vendor:
	go mod vendor

bench:
	go test -bench=BenchmarkExecute -benchmem -benchtime 5s -count=5

pprof:
	go test -bench=BenchmarkExecute -benchmem -benchtime=5s -count=5 -cpuprofile cpu.out -memprofile mem.out

pprof-mem:
	go tool pprof -http :9000 mem.out

pprof-cpu:
	go tool pprof -http :9000 cpu.out

build:
	go build -buildmode=c-shared -o go_app.so main.go