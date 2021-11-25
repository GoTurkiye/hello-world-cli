FROM golang:1.17.1-alpine as build

WORKDIR /src

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN go build

FROM scratch as final

COPY --from=build /src/hello-world-cli ./

ENTRYPOINT [ "./hello-world-cli" ]



