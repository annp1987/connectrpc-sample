FROM golang:1.21-alpine3.19 as build
RUN mkdir -p /app
WORKDIR /go/src/sms-app
COPY . .

ARG release
ENV CGO_ENABLED=0 RELEASE=$release GOSUMDB=off

RUN go clean --modcache && \
    go mod download && \
    go build -o app ./cmd/main.go


FROM alpine:3.19 as prod
WORKDIR /

COPY --from=build /go/src/sms-app/app /usr/local/bin/app

EXPOSE 8080
ENTRYPOINT [ "/usr/local/bin/app" ]