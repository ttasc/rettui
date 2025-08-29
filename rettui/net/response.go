package net

import (
	"net/http"
	"strconv"
	"time"
)

type Response struct {
    *http.Response
    Timing ResponseTime
    BodyBytes []byte
}

type ResponseTime struct {
    DNSLookup    time.Duration
    Connect      time.Duration
    TLSHandshake time.Duration
    FirstByte    time.Duration
}

func (res *Response) Size() int {
    // Format: "HTTP/1.1 200 OK\r\n"
    size := len(res.Proto) + 1 + len(strconv.Itoa(res.StatusCode)) + 1 + len(http.StatusText(res.StatusCode)) + 2

    // Headers
    for key, values := range res.Header {
        size += len(key) + 2 // +2 for ": "
        for _, value := range values {
            size += len(value)
        }
        size += 2 // CRLF
    }
    size += 2 // CRLF at the end of headers

    // Body
    size += len(res.BodyBytes)

    return size
}
