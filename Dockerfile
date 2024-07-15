FROM golang:1.22.5 AS build

WORKDIR /go/src/app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/fserve .

FROM gcr.io/distroless/static-debian12

COPY --from=build /go/bin/fserve /

ENTRYPOINT ["/fserve"]