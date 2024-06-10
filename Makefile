VECTOR 			= amvector
VECTOR_SERVER   = ./cmd/${VECTOR}

GO 				= go
CONFIG_FILE := $(shell pwd)/configs/config.toml

.PHONY: bin
# build bin
bin:
	$(GO) build -o $(VECTOR) $(VECTOR_SERVER)

.PHONY: wire
# generate wire
wire:
	wire ./service/
