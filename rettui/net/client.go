package net

import (
	"net/http"
	"time"
)

type Timming struct {
    DNSLookup           time.Duration
    TCPConnection       time.Duration
    TLSHandshake        time.Duration
    ServerProcessing    time.Duration
    ContentTransfer     time.Duration
}

type HttpClient struct {
    *http.Client
}

func (c *HttpClient) DoReq(method, url, body string) (Response, error) {
    return Response{}, nil
}

func NewHttpClient() *HttpClient {
    tp := http.DefaultTransport.(*http.Transport).Clone()
    // tp.TLSClientConfig = &tls.Config{
    //     ServerName:         host,
    //     InsecureSkipVerify: insecure,
    //     Certificates:       readClientCert(clientCertFile),
    //     MinVersion:         tls.VersionTLS12,
    // }
    return &HttpClient{Client: &http.Client{
        Transport: tp,
        Timeout: 10 * time.Second,
    }}
}

// func NewRequest(method, url string) error {
//     req, err := http.NewRequest(method, url, nil)
//     if err != nil {
//         return err
//     }
//
//     var start, connect, dns, tlsHandshake time.Time
//
//     trace := &httptrace.ClientTrace{
//         DNSStart: func(dsi httptrace.DNSStartInfo) { dns = time.Now() },
//         DNSDone: func(ddi httptrace.DNSDoneInfo) {
//             fmt.Printf("DNS Done: %v\n", time.Since(dns))
//         },
//
//         TLSHandshakeStart: func() { tlsHandshake = time.Now() },
//         TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
//             fmt.Printf("TLS Handshake: %v\n", time.Since(tlsHandshake))
//         },
//
//         ConnectStart: func(network, addr string) { connect = time.Now() },
//         ConnectDone: func(network, addr string, err error) {
//             fmt.Printf("Connect time: %v\n", time.Since(connect))
//         },
//
//         GotFirstResponseByte: func() {
//             fmt.Printf("Time from start to first byte: %v\n", time.Since(start))
//         },
//     }
//
//     req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
//     start = time.Now()
//     if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
//         log.Fatal(err)
//     }
//     fmt.Printf("Total time: %v\n", time.Since(start))
//     return nil
// }
