FROM golang:1.23.5 AS builder
ARG TARGETOS
ARG TARGETARCH

WORKDIR /workspace

COPY go.mod go.mod
COPY cmd/main.go cmd/main.go
COPY internal/ internal/
COPY config.go config.go

RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o kinko cmd/main.go

# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/kinko .
USER 65532:65532

ENTRYPOINT ["/kinko"]
