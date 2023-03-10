# A sampler for zipkin

[![Go Report Card](https://goreportcard.com/badge/github.com/eleven26/zipkin_sampler)](https://goreportcard.com/report/github.com/eleven26/zipkin_sampler)
[![Go](https://github.com/eleven26/zipkin_sampler/actions/workflows/go.yml/badge.svg)](https://github.com/eleven26/zipkin_sampler/actions/workflows/go.yml)
[![codecov](https://codecov.io/github/eleven26/zipkin_sampler/branch/main/graph/badge.svg?token=WJNVVWZALZ)](https://codecov.io/github/eleven26/zipkin_sampler)
[![GitHub license](https://img.shields.io/github/license/eleven26/zipkin_sampler)](https://github.com/eleven26/zipkin_sampler/blob/main/LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/eleven26/zipkin_sampler)

## Description

This is a simple sampler for zipkin. It listens on a port and waits for spans to be sent to it. It collects the spans and sends them to the endpoint if the time of trace is greater than the time specified in the flag, otherwise it discards the trace.

## Installation

```bash
go install github.com/eleven26/zipkin_sampler@v0.0.1
```

## Usage

```bash
zipkin_sampler --port=9422 --endpoint=http://localhost:9411/api/v2/spans --time=5000
```

Flags:

* `--port`: The port to listen on (default 9422)
* `--endpoint`: Zipkin server endpoint
* `--time`: The minimum time of trace to send to zipkin server (default 5000, in milliseconds)
