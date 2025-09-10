package net

import (
	"crypto/tls"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func formatSize(size int) string {
    units := []string{"B", "KB", "MB", "GB"}
    value := float32(size)
    index := 0
    for value >= 1024 && index < len(units)-1 {
        value /= 1024
        index++
    }
    return fmt.Sprintf("%.2f %s", value, units[index])
}

func parseURL(uri string) (*url.URL, error) {
    if !strings.Contains(uri, "://") && !strings.HasPrefix(uri, "//") {
        uri = "//" + uri
    }

    url, err := url.Parse(uri)
    if err != nil {
        return nil, err
    }

    if url.Scheme == "" {
        url.Scheme = "http"
        if !strings.HasSuffix(url.Host, ":80") {
            url.Scheme += "s"
        }
    }
    return url, nil
}

func readClientCert(filename string) ([]tls.Certificate, error) {
    if filename == "" {
        return nil, nil
    }
    var (
        pkeyPem []byte
        certPem []byte
    )

    // read client certificate file (must include client private key and certificate)
    certFileBytes, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    for {
        block, rest := pem.Decode(certFileBytes)
        if block == nil {
            break
        }
        certFileBytes = rest

        if strings.HasSuffix(block.Type, "PRIVATE KEY") {
            pkeyPem = pem.EncodeToMemory(block)
        }
        if strings.HasSuffix(block.Type, "CERTIFICATE") {
            certPem = pem.EncodeToMemory(block)
        }
    }

    cert, err := tls.X509KeyPair(certPem, pkeyPem)
    if err != nil {
        return nil, err
    }
    return []tls.Certificate{cert}, nil
}

func headerKeyValue(h string) (string, string, bool) {
    i := strings.Index(h, ":")
    if i == -1 {
        return "", "", false
    }
    return strings.TrimRight(h[:i], " "), strings.TrimLeft(h[i:], " :"), true
}

func isRedirect(resp *http.Response) bool {
    return resp.StatusCode > 299 && resp.StatusCode < 400
}

// getFilenameFromHeaders tries to automatically determine the output filename,
// when saving to disk, based on the Content-Disposition header.
// If the header is not present, or it does not contain enough information to
// determine which filename to use, this function returns "".
func getFilenameFromHeaders(headers http.Header) string {
    // if the Content-Disposition header is set parse it
    if hdr := headers.Get("Content-Disposition"); hdr != "" {
        // pull the media type, and subsequent params, from
        // the body of the header field
        mt, params, err := mime.ParseMediaType(hdr)

        // if there was no error and the media type is attachment
        if err == nil && mt == "attachment" {
            if filename := params["filename"]; filename != "" {
                return filename
            }
        }
    }

    // return an empty string if we were unable to determine the filename
    return ""
}

// readResponseBody consumes the body of the response.
// readResponseBody returns an informational message about the
// disposition of the response body's contents.
func readResponseBody(req *http.Request, resp *http.Response, outputFile string) error {
    if isRedirect(resp) || req.Method == http.MethodHead {
        return errors.New("refusing to download a redirect or HEAD response body")
    }

    var w io.Writer

    if outputFile == "" {
        w = io.Discard
    } else {
        filename := outputFile

        f, err := os.Create(filename)
        if err != nil {
            return err
        }
        defer f.Close()
        w = f
    }

    if _, err := io.Copy(w, resp.Body); err != nil && w != io.Discard {
        return err
    }

    return nil
}
