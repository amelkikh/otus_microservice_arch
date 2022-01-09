FROM golang:1.16.5-alpine as base
ENV CGO_ENABLED=0
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

FROM base as builder
ARG GIT_COMMIT_HASH
ARG GIT_TAG
COPY . .
RUN go build -ldflags="-s -w -X 'main.Version=${GIT_TAG}' -X 'main.Commit=${GIT_COMMIT_HASH}' -X 'main.BuildTime=\"$(LC_TIME=en_US.utf8;date -u)\"'" -o main ./cmd/app

FROM builder as unit-test
RUN go test -v -count=1 ./...

FROM alpine:3.14.2 as release
EXPOSE 8000
COPY --from=builder /src/main /server
ENTRYPOINT ["/server"]
CMD ["-port", "8000"]
