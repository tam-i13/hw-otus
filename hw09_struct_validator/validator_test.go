package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:nolintlint
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{User{
			ID:     "111111111111111111111111111111111111",
			Name:   "Ivan",
			Age:    32,
			Email:  "test@test.test",
			Role:   "admin",
			Phones: []string{"79999999999"},
			meta:   json.RawMessage{},
		}, nil},
		{App{Version: "12345"}, nil},
		{Token{Header: []byte("1"), Payload: []byte("2"), Signature: []byte("3")}, nil},
		{Response{Code: 21, Body: "body"}, ErrValidationTag},
		{"string", ErrIsNotStruct},
		{User{
			ID:     "1",
			Name:   "Ivan",
			Age:    140,
			Email:  "test@test.test",
			Role:   "admin",
			Phones: []string{"79999999999", "797999993999", "797999999999"},
			meta:   json.RawMessage{},
		}, errors.New("ID - len string value less, Age - value more max, Phones - len string value less")},
	}

	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()
			res := Validate(tt.in)
			if tt.expectedErr == nil {
				require.Equal(t, tt.expectedErr, res)
			} else {
				require.Equal(t, tt.expectedErr.Error(), res.Error())
			}
		})
	}
}
