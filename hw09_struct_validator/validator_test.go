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

	CheckStruct struct {
		StringSlice []string `validate:"len:3"`
		IntSlice    []int    `validate:"min:0|max:100"`
		EmailSlice  []string `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{App{Version: "12345"}, nil},
		{App{Version: "123456"}, errors.New("Version - len not equal target")},
		{User{
			ID:     "111111111111111111111111111111111111",
			Name:   "Ivan",
			Age:    32,
			Email:  "test@test.test",
			Role:   "admin",
			Phones: []string{"79999999999"},
			meta:   json.RawMessage{},
		}, nil},
		{Token{Header: []byte("1"), Payload: []byte("2"), Signature: []byte("3")}, nil},
		{Response{Code: 21, Body: "body"}, ErrValidationTag},
		{"string", ErrIsNotStruct},
		{User{
			ID:     "111111111111111111111111111111111111",
			Name:   "Ivan",
			Age:    32,
			Email:  "test@test.test",
			Role:   "admin",
			Phones: []string{"79999999999", "79", "79799555555555555999"},
			meta:   json.RawMessage{},
		}, errors.New("Phones - len not equal target value 79, Phones - len not equal target value 79799555555555555999")},
		{User{
			ID:     "111111111111111111111111111111111111",
			Name:   "Ivan",
			Age:    32,
			Email:  "test.test",
			Role:   "admin",
			Phones: []string{"79999999999"},
			meta:   json.RawMessage{},
		}, errors.New("Email - value not in regexp")},
		{User{
			ID:     "111111111111111111111111111111111111",
			Name:   "Ivan",
			Age:    132,
			Email:  "test@test.test",
			Role:   "DevOps",
			Phones: []string{"79999999999"},
			meta:   json.RawMessage{},
		}, errors.New("Age - value more max, Role - value not in range string")},
		{CheckStruct{
			StringSlice: []string{"qwe", "asd", "zxc"},
			IntSlice:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, nil},
		{CheckStruct{
			StringSlice: []string{"qwe", "asd", "zxc"},
			IntSlice:    []int{1001, 2001, 3, 4, 5, 6, 7, 8, 9},
		}, errors.New("IntSlice - value more max value 1001, IntSlice - value more max value 2001")},
		{CheckStruct{
			StringSlice: []string{"qw", "asd", "zxcv"},
			IntSlice:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, errors.New("StringSlice - len not equal target value qw, StringSlice - len not equal target value zxcv")},
		{CheckStruct{
			EmailSlice: []string{"test@test.test", "test.test", "t@tt@t"},
		}, errors.New("EmailSlice - value not in regexp value test.test, EmailSlice - value not in regexp value t@tt@t")},
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
