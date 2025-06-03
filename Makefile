.PHONY: build

BINARY_NAME=wasmdash

# build builds the tailwind css sheet, and compiles the binary into a usable thing.
build:
	go mod tidy && \
	tailwindcss -m -o static/css/tailwind.css && \
	go generate && \
	go build -ldflags="-w -s" -o ${BINARY_NAME}

# dev runs the development server where it builds the tailwind css sheet,
# and compiles the project whenever a file is changed.
dev:
	templ generate --watch --cmd="go generate" &\
	templ generate --watch --cmd="go run ."

clean:
	go clean

help:
