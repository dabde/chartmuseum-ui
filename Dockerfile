FROM library/golang:1 as builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build cmd/chartmuseum-ui.go

FROM alpine:3.11
RUN apk add --no-cache curl cifs-utils ca-certificates \
    && adduser -D -u 1000 chartmuseum
COPY --from=builder /build/chartmuseum-ui /chartmuseum-ui
COPY /views /views
COPY /static /static
COPY /conf /conf
USER 1000
ENTRYPOINT ["/chartmuseum-ui"]
