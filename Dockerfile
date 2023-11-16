FROM golang:alpine AS builderGo
RUN apk --no-cache -U upgrade && \
    apk --no-cache add --upgrade make build-base
WORKDIR /go/src/github.com/JDinABox/tool-server
COPY go.* ./
RUN go mod download
COPY ./cmd/ ./cmd/
COPY ./Makefile ./Makefile
COPY *.go ./
RUN --mount=type=cache,target=/root/.cache/go-build make build-docker

FROM scratch

COPY --from=builderGo /go/src/github.com/JDinABox/tool-server/cmd/tool-server/tool-server /tool-server

CMD ["/tool-server"]