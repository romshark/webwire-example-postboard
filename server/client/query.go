package client

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
)

// Query performs an API request unmarshalls the incoming JSON result
// into 'result'
func (c *apiClient) Query(
	ctx context.Context,
	method api.Method,
	params interface{},
	result interface{},
) error {
	var err error
	var encoded []byte
	if params != nil {
		encoded, err = json.Marshal(params)
		if err != nil {
			return errors.Wrap(err, "marshalling failed")
		}
	}

	reply, err := c.connection.Request(
		ctx,
		string(method),
		wwr.NewPayload(
			wwr.EncodingUtf8,
			encoded,
		),
	)
	if err != nil {
		return err
	}

	if result != nil {
		if err := json.Unmarshal(reply.Data(), result); err != nil {
			return errors.Wrap(err, "unmarshalling failed")
		}
	}

	return nil
}
