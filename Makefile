GOPATH=${HOME}/go

.PHONY: fmt
fmt:
	./scripts/fmt.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: test
test:
	./scripts/test.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: proto
proto:
	./scripts/proto.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: install
install:
	./scripts/install.sh
