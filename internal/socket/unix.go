package socket

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

func HttpUnixSocket() http.Client {
	fmt.Println("Init Unix HTTP client")

	httpc := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/var/run/docker.sock")
			},
		},
	}

	return httpc
}
