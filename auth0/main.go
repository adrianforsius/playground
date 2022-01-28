package main

import (
	"log"

	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
)

const client_id = "4u1wwfRwA963ifLbXpkbhGja0QCVJdUj"
const client_secret = "Ze_L6GbldL-DPS5NwV0DLl9PNp9woqHmJ1fSqNee8Qw89mUOPN2UWBmhAlIjijZh"

//const client_id = "JqloGhgICYcmrd4e0CD5ROBp2MC4dpvF"
//const client_secret = "1NSqQ-mzfPwVpEER5Ux1iAC-IWEZWxclPAz5vnLuWDdK-WHRkQADRUXmPglRddkO"

const domain = "dd3.us.auth0.com"

func main() {
	m, err := management.New(
		domain,
		management.WithClientCredentials(
			client_id,
			client_secret,
		),
	)
	if err != nil {
		log.Fatalf("failed to create manager: %s", err)
	}
	err = m.User.Create(&management.User{
		Connection: auth0.String("Username-Password-Authentication"),
		Email:      auth0.String("hello@test.com"),
		Password:   auth0.String("aqaEsW75YSQK6T"),
	})
	if err != nil {
		log.Fatalf("failed to create user: %s", err)
	}
}
