package backend

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func PostData(url string, data any) ([]byte, int, error) {
	marshalled, err := json.Marshal(data)
	if err != nil {
		return nil, 0, errors.Wrap(err, "could not marshal data")
	}

	request, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer(marshalled),
	)

	if err != nil {
		return nil, 0, errors.Wrap(err, "Could not build POST request")
	}

	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, 0, errors.Wrap(err, "Could not make a POST request")
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, 0, errors.Wrap(err, "Could not read response body")
	}

	return responseBytes, response.StatusCode, nil
}

func BuildResponseError(resData []byte) error {
	var res struct {
		Message string `json:"message"`
	}

	err := json.Unmarshal(resData, &res)
	if err != nil {
		return errors.Wrap(err, "could not unmarshal choices confirmation error")
	}

	return errors.New(res.Message)
}
