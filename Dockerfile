# syntax=docker/dockerfile:1

FROM golang:1.24-bookworm AS build

WORKDIR /src

RUN apt-get update \
  && apt-get install -y --no-install-recommends ca-certificates build-essential wget tar \
  && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=1 \
  GOOS=linux \
  GOARCH=amd64

RUN go build -o /out/gopherai ./main.go

FROM debian:bookworm-slim AS runtime

ARG ORT_VERSION=1.20.1

WORKDIR /app

RUN apt-get update \
  && apt-get install -y --no-install-recommends ca-certificates wget tar \
  && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /opt/onnxruntime \
  && wget -q -O /tmp/onnxruntime.tgz "https://github.com/microsoft/onnxruntime/releases/download/v${ORT_VERSION}/onnxruntime-linux-x64-${ORT_VERSION}.tgz" \
  && tar -xzf /tmp/onnxruntime.tgz -C /opt/onnxruntime --strip-components=1 \
  && rm -f /tmp/onnxruntime.tgz

ENV LD_LIBRARY_PATH=/opt/onnxruntime/lib:${LD_LIBRARY_PATH} \
  ORT_DYLIB_PATH=/opt/onnxruntime/lib/libonnxruntime.so

COPY --from=build /out/gopherai /app/gopherai
COPY --from=build /src/config /app/config

EXPOSE 9090

CMD ["/app/gopherai"]
