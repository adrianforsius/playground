package main

import (
	"log"

	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
)

const client_id = "client_id"
const client_secret = "client_secret"


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
