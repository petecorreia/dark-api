# build
FROM golang:1.15-alpine as build
ARG SERVICE
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . /go/src/app
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -v -o /go/bin/app .

# run
FROM gcr.io/distroless/base
COPY --from=build /go/bin/app /
CMD ["/app"]
