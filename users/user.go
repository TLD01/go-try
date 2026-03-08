package users

import (
	"aerowatch.com/api/common"
	"github.com/TLD01/tld_constants/types"
)

type User struct {
	common.Persisted
	Sub           string       `json:"sub"`
	Name          string       `json:"name"`
	GivenName     string       `json:"givenName"`
	FamilyName    string       `json:"familyName"`
	Email         string       `json:"email"`
	EmailVerified bool         `json:"emailVerified"`
	Picture       string       `json:"picture"`
	AuthProvider  AuthProvider `json:"authProvider"`
	LastSignOn    types.ISO8601Time    `json:"lastSignOn"`
}
