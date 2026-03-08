package repository

import (
	"time"

	"aerowatch.com/api/repository"
	"aerowatch.com/api/users"
	"github.com/TLD01/tld_constants/types"
)	

type UserEntity struct {
	repository.DBEntity
	Sub           string       `json:"sub" bson:"sub"`
	Name          string       `json:"name" bson:"name"`
	GivenName     string       `json:"givenName" bson:"givenName"`
	FamilyName    string       `json:"familyName" bson:"familyName"`
	Email         string       `json:"email" bson:"email"`
	EmailVerified bool         `json:"emailVerified" bson:"emailVerified"`
	Picture       string       `json:"picture" bson:"picture"`
	AuthProvider  users.AuthProvider `json:"authProvider" bson:"authProvider"`
	LastSignOn    time.Time        `json:"lastSignOn" bson:"lastSignOn"` // Store as Unix timestamp for easier BSON handling
}

func Create(u *users.User) *UserEntity {
	return &UserEntity{
		DBEntity:      repository.Create(u.Persisted),
		Sub:           u.Sub,
		Name:          u.Name,
		GivenName:     u.GivenName,
		FamilyName:    u.FamilyName,
		Email:         u.Email,
		EmailVerified: u.EmailVerified,
		Picture:       u.Picture,
		AuthProvider:  u.AuthProvider,
		LastSignOn:    time.Time(u.LastSignOn), // Convert from ISO8601Time to time.Time
	}		
}



func (e *UserEntity) ToUser() *users.User {
	return &users.User{
		Persisted:     e.DBEntity.ToPersisted(),
		Sub:           e.Sub,
		Name:          e.Name,
		GivenName:     e.GivenName,
		FamilyName:    e.FamilyName,
		Email:         e.Email,
		EmailVerified: e.EmailVerified,
		Picture:       e.Picture,
		AuthProvider:  e.AuthProvider,
		LastSignOn:    types.ISO8601Time(e.LastSignOn),
	}
}