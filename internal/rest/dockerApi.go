package rest

import "net/http"

func Get(httpc http.Client) (*http.Response, error) {
	var response *http.Response
	response, err := httpc.Get("http://localhost/v1.41/containers/f709242fb73e91b4988d3faf8487eda1e0cfef7c4bdb456f91d86dc733158de3/logs?stdout=true")

	if err != nil {
		return nil, err
	}

	return response, nil
}
