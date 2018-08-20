package resolver

import (
	"encoding/json"
	"fmt"

	wwr "github.com/qbeon/webwire-go"
	engiface "github.com/qbeon/webwire-messenger/server/apisrv/modules/engine"
)

// helps parsing incomming parameters
func parseParameters(message wwr.Message, parseTo interface{}) error {
	if err := json.Unmarshal(message.Payload().Data(), parseTo); err != nil {
		return wwr.ReqErr{
			Code:    engiface.ErrInvalidRequest.String(),
			Message: fmt.Sprintf("parameter unmarshalling failed: %s", err),
		}
	}
	return nil
}
