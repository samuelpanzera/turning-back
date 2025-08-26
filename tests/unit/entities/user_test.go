package entities_test

import (
	"testing"

	"github.com/samuelpanzera/turning-back/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestUser_FullName(t *testing.T) {
	tests := []struct {
		name      string
		user      entities.User
		expected  string
	}{
		{
			name: "should return full name when both first and last names are provided",
			user: entities.User{
				FirstName: "John",
				LastName:  "Doe",
			},
			expected: "John Doe",
		},
		{
			name: "should return full name with empty first name",
			user: entities.User{
				FirstName: "",
				LastName:  "Doe",
			},
			expected: " Doe",
		},
		{
			name: "should return full name with empty last name",
			user: entities.User{
				FirstName: "John",
				LastName:  "",
			},
			expected: "John ",
		},
		{
			name: "should return empty string when both names are empty",
			user: entities.User{
				FirstName: "",
				LastName:  "",
			},
			expected: " ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.user.FullName()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestUser_TableName(t *testing.T) {
	user := entities.User{}
	assert.Equal(t, "users", user.TableName())
}