package main

import (
	"net/http"
	"time"
)

/*
 * Problems with the default HTTP server:
 *
 * Similarly as the client, it doesn't specify any timeouts.
 * Steps involved in handling an HTTP request:
 * 1. Wait for the client to send the request
 * 2. TLS handshake
 * 3. Read the request headers
 * 4. Read the request body
 * 5. Write the response
 *
 * Timeouts should be provided because otherwise malicious clients
 * can create never ending requests and exhaust system resources.
 *
 */

func HttpServerWithDefaults() {
	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 500 * time.Millisecond, // 3.
		ReadTimeout:       500 * time.Millisecond, // 1. - 4.
		Handler: http.TimeoutHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // 5.
			// handle the request
		}), time.Second, "foo"),
	}

	_ = server
}
