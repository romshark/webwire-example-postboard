package engine

import (
	"github.com/pkg/errors"
	wwr "github.com/qbeon/webwire-go"
)

// ParseSessionInfo parses the session info from a variant map
func ParseSessionInfo(data map[string]interface{}) wwr.SessionInfo {
	// Parse identifier
	var userIdentifier Identifier

	switch val := data["id"].(type) {
	case Identifier:
		userIdentifier = val
	case string:
		var ident Identifier
		if err := ident.FromString(val); err != nil {
			panic(errors.Wrap(err, "couldn't parse user identifier"))
		}
		userIdentifier = ident
	}

	// Parse client type
	var userType UserType

	switch val := data["type"].(type) {
	case UserType:
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
