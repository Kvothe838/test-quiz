package backend

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func PostData(url string, data any) ([]byte, error) {
	marshalled, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal data")
	}

	request, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer(marshalled),
	)

	if err != nil {
		return nil, errors.Wrap(err, "Could not build POST request")
	}

	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "Could not make a POST request")
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Could not read response body")
	}

	return responseBytes, nil
}
