FROM golang:alpine AS builderGo
RUN apk --no-cache -U upgrade && \
    apk --no-cache add --upgrade make build-base
WORKDIR /go/src/github.com/JDinABox/tool-server
COPY go.* ./
RUN go mod download
COPY ./cmd/ ./cmd/
COPY ./Makefile ./Makefile
COPY *.go ./
RUN --mount=type=cache,target=/root/.cache/go-build make build

# Docker build
FROM alpine:latest

RUN apk --no-cache -U upgrade \
    && apk --no-cache add --upgrade ca-certificates \
    && wget -O /bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.5/dumb-init_1.2.5_x86_64 \
    && chmod +x /bin/dumb-init

COPY --from=builderGo /go/src/github.com/JDinABox/tool-server/cmd/tool-server/tool-server /bin/tool-server
WORKDIR /etc/tool-server/

# Use dumb-init to prevent gofiber prefork from failing as PID 1
ENTRYPOINT ["/bin/dumb-init", "--"]
CMD ["/bin/tool-server"]