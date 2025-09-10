package net

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Request struct {
    *http.Request
}

func (req *Request) Size() int {
    size := len(req.URL.String()) + len(req.Method) + len(req.Proto) + 4 // 4 for space and CRLF

    // Headers
    for key, values := range req.Header {
        size += len(key) + 2 // +2 for ": "
        for _, value := range values {
            size += len(value)
        }
        size += 2 // CRLF
    }
    size += 2 // CRLF at the end of headers

    // Body
    if req.Body != nil {
        bodyBytes, _ := req.GetBody()
        buf := new(bytes.Buffer)
        buf.ReadFrom(bodyBytes)
        size += buf.Len()
        req.Body, _ = req.GetBody() // Recovery body
    }

    return size
}

func NewRequest(method string, url *url.URL, headers []string, body string) *http.Request {
    req, err := http.NewRequest(method, url.String(), createBody(body))
    if err != nil {
        log.Fatalf("unable to create request: %v", err)
    }
    for _, h := range headers {
        k, v, ok := headerKeyValue(h)
        if !ok { continue }
        if strings.EqualFold(k, "host") {
            req.Host = v
            continue
        }
        req.Header.Add(k, v)
    }
    return req
}

func createBody(body string) io.Reader {
    if strings.HasPrefix(body, "@") {
        filename := body[1:]
        f, err := os.Open(filename)
        if err != nil {
            log.Fatalf("failed to open data file %s: %v", filename, err)
        }
        return f
    }
    return strings.NewReader(body)
}
