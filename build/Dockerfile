FROM golang:1.15-alpine AS builder

RUN /sbin/apk update && \
	/sbin/apk --no-cache add ca-certificates git tzdata && \
	/usr/sbin/update-ca-certificates

RUN adduser -D -g '' depot
WORKDIR /home/depot

COPY go.mod /home/depot
COPY go.sum /home/depot
COPY cmd    /home/depot/cmd
COPY pkg    /home/depot/pkg
COPY *.go   /home/depot

ARG VERSION

RUN CGO_ENABLED=0 go build -a -tags netgo,osusergo \
    -ldflags "-extldflags '-static' -s -w" \
    -ldflags "-X main.version=$VERSION" -o depot ./cmd/depot

FROM busybox:musl

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /home/depot/depot /home/depot

USER depot
WORKDIR /home
EXPOSE 8081

ENTRYPOINT ["/home/depot"]
