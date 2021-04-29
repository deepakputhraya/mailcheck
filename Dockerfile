FROM golang:latest AS build_base
WORKDIR /tmp/app
COPY . .
RUN make build-server

FROM alpine:latest
COPY --from=build_base /tmp/app/server/bin/app /mailcheck/server
EXPOSE 8080
CMD ["/mailcheck/server"]
