package setup

import (
	"github.com/qbeon/webwire-example-postboard/server/apisrv/api"
	"github.com/qbeon/webwire-example-postboard/server/client"
	"github.com/qbeon/webwire-example-postboard/server/test/setup/generate"
)

// CreateUserRand creates a random user profile with random fields
// expecting the operation to be successful, retrieves the profile
// and verifies whether all fields of the profile are correct
func (h *Helper) CreateUserRand(
	admin client.ApiClient,
	userType api.UserType,
) (*api.User, string, client.ApiClient) {
	randomProfile := generate.User(userType)
	randomPassword := generate.RandStr(10)

	profile, clt := h.CreateUser(
		admin,
		api.CreateUserParams{
			FirstName: randomProfile.FirstName,
			LastName:  randomProfile.LastName,
			Username:  randomProfile.Username,
			Password:  randomPassword,
			Type:      userType,
		},
	)

	return profile, randomPassword, clt
}
