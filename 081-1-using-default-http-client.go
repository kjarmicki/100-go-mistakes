package main

import (
	"net"
	"net/http"
	"time"
)

/*
 * Problems with the default HTTP client:
 *
 * It doesn't specify any timeouts.
 * Steps involved in an HTTP request:
 * 1. Dial to establish TCP connection
 * 2. TLS handshake
 * 3. Send the request
 * 4. Read the response headers
 * 5. Read the response body
 */

// an HTTP client with timeouts explicitly set:
func HttpClientWithDefaults() {
	client := &http.Client{
		Timeout: 5 * time.Second, // entire request, steps 1 - 5
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second, // 1.
			}).DialContext,
			TLSHandshakeTimeout:   time.Second, // 2.
			ResponseHeaderTimeout: time.Second, // 4.
		},
	}

	_ = client
}

/*
 * Another thing to keep in mind is that the default HTTP client
 * does connection pooling. Idle connection is kept in the pool for 90 seconds
 * and the pool size is 100 connections, 2 per each host.
 */
