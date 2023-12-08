FROM golang:1.18.10 AS builder
ENV GOPROXY="https://goproxy.cn,direct"

WORKDIR /bin
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w --extldflags "-static -fpic"' -a -installsuffix nocgo -o /wlog .

FROM scratch
COPY --from=builder /wlog bin/wlog
ENTRYPOINT ["/bin/wlog"]