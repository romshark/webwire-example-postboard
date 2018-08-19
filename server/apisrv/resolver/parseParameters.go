package resolver

import (
	"encoding/json"
	"fmt"

	wwr "github.com/qbeon/webwire-go"
)

// helps parsing incomming parameters
func parseParameters(messagePayload wwr.Payload, parseTo interface{}) error {
	if err := json.Unmarshal(messagePayload.Data(), parseTo); err != nil {
		return wwr.ReqErr{
			Code:    "INVALID_PARAMS",
			Message: fmt.Sprintf("parameter unmarshalling failed: %s", err),
		}
	}
	return nil
}
