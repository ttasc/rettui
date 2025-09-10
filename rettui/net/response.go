package net

import (
	"net/http"
	"strconv"
)

type Response struct {
    *http.Response
    Timing Timming
    body []byte
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
    size += len(res.body)

    return size
}

func (res *Response) GetBody() []byte {
    return res.body
}
