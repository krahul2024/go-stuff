.PHONY: build run test clean

APP_NAME := crud-mux
BUILD_DIR := build

build: $(BUILD_DIR)/$(APP_NAME)

$(BUILD_DIR)/$(APP_NAME):
	go build -o $(BUILD_DIR)/$(APP_NAME) .

run: build
	$(BUILD_DIR)/$(APP_NAME)

test:
	go test -v ./...

clean:
	rm -rf $(BUILD_DIR)
