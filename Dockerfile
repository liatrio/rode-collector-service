# syntax = docker/dockerfile:experimental
FROM golang:1.15 as builder

ENV GRPC_HEALTH_PROBE_VERSION="v0.3.1"
RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.sum /workspace/
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY proto/ proto/
COPY protodeps/ protodeps/
COPY server/ server/
COPY config/ config/
COPY auth/ auth/

RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o rode main.go

# ---------------

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/rode .
COPY --from=builder /bin/grpc_health_probe .
USER nonroot:nonroot

ENTRYPOINT ["/rode"]
