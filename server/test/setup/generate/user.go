package generate

import (
	"math"
	"time"

	"github.com/qbeon/webwire-messenger/server/apisrv/api"
)

// User generates a random user profile
func User(userType api.UserType) *api.User {
	return &api.User{
		Identifier:   api.NewIdentifier(),
		FirstName:    RandStr(8),
		LastName:     RandStr(8),
		Username:     RandStr(8),
		Registration: time.Now().UTC(),
		// Round random reputation to 1 decimal place
		Reputation: math.Round(RandNum(-10, 10)*float64(10)) / float64(10),
		Type:       userType,
	}
}
