package model

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// create user test id
var createUserID int64

func TestCreateUser(t *testing.T) {
	// password, _ := helper.MakeHash("12341234")
	// var user = User{
	// 	Name:     "Test",
	// 	Email:    "tt312@tt.com",
	// 	Password: password,
	// }
	// user, err := api.Handler(user)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// createUserID = user.ID

	// findUser, _ := GetUserByID(createUserID)

	// jsonUser, _ := json.Marshal(user)
	// jsonFindUser, _ := json.Marshal(findUser)

	// assert.JSONEq(t, string(jsonUser), string(jsonFindUser))
}

func TestDestroyUserByID(t *testing.T) {
	n, err := DestroyUserByID(createUserID)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, n, 1)
}
