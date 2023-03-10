package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var spans = `[{
    "id": "3159df044ee05f27",
    "traceId": "b4a26e8e797152529a946d201e25c90b",
    "parentId": null,
    "localEndpoint": {
      "serviceName": "local-api"
    },
    "name": "GET /api/core/tracer",
    "timestamp": 1678366883749929,
    "duration": 5780570,
    "tags": {
      "op.status_code": "Ok",
      "op.status_description": "Not an error; returned on success.",
      "app.env": "qa",
      "request.start_time": "2023-03-09 21:01:23",
      "http.method": "GET",
      "http.client_ip": "127.0.0.1",
      "http.url": "http://localhost:5200/api/core/tracer",
      "http.status_code": "200",
      "request.end_time": "2023-03-09 21:01:29",
      "request.id": "814856409d8a3b5556368613018",
      "my.time": "205ms",
      "my.count": "10",
      "qu.time": "287ms",
      "qu.count": "10"
    }
  }]`

var (
	url         = "http://localhost:9433/traces"
	contentType = "application/json"
	server      *httptest.Server
)

func BenchmarkServeHTTP(b *testing.B) {
	if server == nil {
		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusAccepted)
		}))

		go serveHTTP("9433", server.URL, 5000)
	}

	b.ReportAllocs()
	// post => 9443 => server
	for i := 0; i < b.N; i++ {
		resp, err := http.Post(url, contentType, strings.NewReader(spans))
		if err != nil {
			b.Fatal(err)
		}
		if resp.StatusCode != http.StatusAccepted {
			b.Fatal("status code not 202")
		}
	}
}
