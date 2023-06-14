package interfaces

import (
	"testing"

	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/domain/entity"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		entity.User
	}{
		{
			desc: "Testing User ID",
			User: entity.User{
				ID:        0,
				FirstName: "John",
				LastName:  "Doe",
				Email:     "johndoe@example.com",
				Password:  "password123",
			},
		},
		{
			desc: "Testing Entire Struct",
			User: entity.User{
				ID:        1,
				FirstName: "John",
				LastName:  "Doe",
				Email:     "johndoe@example.com",
				Password:  "password123",
			},
		},
		{
			desc: "Testing First Name",
			User: entity.User{
				ID:        2,
				FirstName: "j",
				LastName:  "Doe",
				Email:     "johndoe@example.com",
				Password:  "password123",
			},
		},
		{
			desc: "Testing Last Name",
			User: entity.User{
				ID:        3,
				FirstName: "John",
				LastName:  "d",
				Email:     "johndoe@example.com",
				Password:  "password123",
			},
		},
		{
			desc: "Testing Email Address",
			User: entity.User{
				ID:        4,
				FirstName: "John",
				LastName:  "Doe",
				Email:     "johndoe#example.com",
				Password:  "password123",
			},
		},
		{
			desc: "Testing Password",
			User: entity.User{
				ID:        5,
				FirstName: "John",
				LastName:  "Doe",
				Email:     "johndoe@example.com",
				Password:  "passw",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if err := validate.Struct(&tC); err != nil {
				t.Errorf("Struct validation error : %v", err.Error())
			}
		})
	}
}
