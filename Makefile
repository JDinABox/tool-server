build:
	go build -ldflags="-s -w" -o ./cmd/tool-server/tool-server ./cmd/tool-server

build-docker:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static" -s -w' -o ./cmd/tool-server/tool-server ./cmd/tool-server