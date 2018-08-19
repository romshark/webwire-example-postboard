package sessinfo

import (
	"github.com/pkg/errors"
	wwr "github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// ParseSessionInfo parses the session info from a variant map
func ParseSessionInfo(data map[string]interface{}) wwr.SessionInfo {
	// Parse identifier
	var userIdentifier api.Identifier

	switch val := data["id"].(type) {
	case api.Identifier:
		userIdentifier = val
	case string:
		var ident api.Identifier
		if err := ident.FromString(val); err != nil {
			panic(errors.Wrap(err, "couldn't parse user identifier"))
		}
		userIdentifier = ident
	}

	// Parse client type
	var userType api.UserType

	switch val := data["type"].(type) {
	case api.UserType:
		userType = val
	case string:
		if err := userType.FromString(val); err != nil {
			panic(errors.Wrap(
				err,
				"couldn't parse UserType from session info",
			))
		}
	}

	return &SessionInfo{
		UserIdentifier: userIdentifier,
		UserType:       userType,
	}
}
