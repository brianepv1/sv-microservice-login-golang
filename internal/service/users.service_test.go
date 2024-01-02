package service

import (
	"context"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		FirstName     string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "RegisterUser__Success",
			Email:         "test@test.com",
			FirstName:     "User 1",
			Password:      "validPassword",
			ExpectedError: nil,
		}, {
			Name:          "RegisterUser__UserAlreadyExists",
			Email:         "test@exists.com",
			FirstName:     "User 1",
			Password:      "validPassword",
			ExpectedError: ErrUserAlreadyExists,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			// Assert expectations at the end of the test
			//defer repo.AssertExpectations(t)
			repo.Mock.Test(t)

			s := New(repo)

			err := s.RegisterUser(ctx, tc.Email, tc.FirstName, tc.Password)

			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "LoginUser_Success",
			Email:         "test@exists.com",
			Password:      "validPassword",
			ExpectedError: nil,
		},
		{
			Name:          "LoginUser_InvalidPassword",
			Email:         "test@exists.com",
			Password:      "invalidPassword",
			ExpectedError: PwsUserIsWrong,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			_, errr := s.LoginUser(ctx, tc.Email, tc.Password)

			if errr != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, errr)
			}

		})
	}
}
