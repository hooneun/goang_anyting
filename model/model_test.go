package model

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// create user test id
var createUserID int64

func TestCreateUser(t *testing.T) {
	var user = User{
		Name:     "Test",
		Email:    "tt32@tt.com",
		Password: "12341234",
	}
	user, err := CreateUser(user)
	if err != nil {
		log.Fatal(err)
	}

	createUserID = user.ID
}

func TestDestroyUserByID(t *testing.T) {
	n, err := DestroyUserByID(createUserID)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, n, int64(1))
}
