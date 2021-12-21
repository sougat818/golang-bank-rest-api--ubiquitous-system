FROM golang:alpine AS build
WORKDIR /go/src
COPY . .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bank/ ./...

FROM scratch AS runtime
COPY --from=build /go/src/bank/ /go/bin/bank/

EXPOSE 8080/tcp
ENTRYPOINT ["/go/bin/bank/server"]