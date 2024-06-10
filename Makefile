VECTOR 			= amvector
VECTOR_SERVER   = ./cmd/${VECTOR}

BOOTSTRAP       = ambootstrap
BOOTSTRAP_SERVER = ./cmd/${BOOTSTRAP}

GO 				= go
CONFIG_FILE := $(shell pwd)/configs/config.toml

.PHONY: dev
# run dev
dev:
	@$(GO) run $(SERVER)/main.go run -c $(CONFIG_FILE)

.PHONY: bin
# build bin
bin:
	$(GO) build -o $(APP) $(SERVER)

.PHONY: wire
# generate wire
wire:
	wire ./service/
