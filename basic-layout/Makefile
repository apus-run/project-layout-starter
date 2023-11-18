# 初始化环境
.PHONY: setup
setup:
	@sh ./script/setup.sh

.PHONY: lint
lint:
	golangci-lint run

.PHONY: fmt
fmt:
	@sh ./script/fmt.sh

.PHONY: tidy
tidy:
	@go mod tidy -v

.PHONY: check
check:
	@$(MAKE) --no-print-directory fmt
	@$(MAKE) --no-print-directory tidy


# 生成 mock 文件
.PHONY: mock
mock:
	@go generate -tags=wireinject ./...
	@$(MAKE) tidy
	@echo "go mock finished"

# 单元测试
.PHONY: test
	@go test -race ./..
	@echo "go tests finished"

# e2e 测试
.PHONY: e2e
e2e:
	sh ./scripts/integrate_test.sh
	@echo "go e2e finished"