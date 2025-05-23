# Note: This Dockerfile assumes that the "bundle-crds" taskfile target has been run already
ARG VERSION_FLAGS

# Build the manager binary
FROM golang:1.24 AS builder
ARG VERSION_FLAGS

WORKDIR /workspace/
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY . ./

# Build
RUN CGO_ENABLED=0 go build -ldflags "${VERSION_FLAGS}" -o ./bin/aso-controller ./cmd/controller/

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/bin/aso-controller .
COPY ./out/crds ./crds
USER nonroot:nonroot
ENTRYPOINT ["/aso-controller"]
