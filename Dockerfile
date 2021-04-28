FROM golang:latest AS build_base
WORKDIR /tmp/app
COPY . .
RUN CGO_ENABLED=0 make build-server

FROM alpine:latest
RUN apk add ca-certificates
COPY --from=build_base /tmp/app/server/bin/app /mailcheck/server
EXPOSE 8080
CMD ["sh", "/mailcheck/server"]
