package main

import (
	"io"
	"log"
	"net/http"
)

/*
 * Resources that implement the io.Closer interface should be closed in order to avoid resource leaks.
 * GC can't reclaim such resources if they're not closed and therefore memory will be needlessly allocated.
 */

type httpHandler struct {
	client http.Client
	url    string
}

func (h httpHandler) getBody() (string, error) {
	resp, err := h.client.Get(h.url)
	if err != nil {
		return "", err
	}
	// response MUST be closed because otherwise there will be a resource leak
	// note that on the server side there's no need to close the request body because server does that automatically
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("failed to close response body: %v'n", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
