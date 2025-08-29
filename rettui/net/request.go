package net

import (
	"bytes"
	"net/http"
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

