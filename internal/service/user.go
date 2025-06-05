// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
)

type (
	IUser interface {
		// Create creates user account.
		Create(ctx context.Context, in model.UserCreateInput) (err error)
		// IsSignedIn checks and returns whether current user is already signed-in.
		IsSignedIn(ctx context.Context) bool
		// SignIn creates session for given user account.
		SignIn(ctx context.Context, in model.UserSignInInput) (err error)
		// SignOut removes the session for current signed-in user.
		SignOut(ctx context.Context) error
		// IsPassportAvailable checks and returns given passport is available for signing up.
		IsPassportAvailable(ctx context.Context, passport string) (bool, error)
		// IsNicknameAvailable checks and returns given nickname is available for signing up.
		IsNicknameAvailable(ctx context.Context, nickname string) (bool, error)
		// GetProfile retrieves and returns current user info in session.
		GetProfile(ctx context.Context) *entity.User
		// Recharge adds funds to the user's account by writing a log entry to the recharge_log table.
		// It does not update the user.balance field.
		// Returns the new balance after recharge, or an error if the operation fails.
		Recharge(ctx context.Context, userId int, amount float64) (float64, error)
		// GetBalance returns the current account balance for the given user.
		// The balance is calculated as the sum of all recharge amounts in the recharge_log table.
		GetBalance(ctx context.Context, userId int) (float64, error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
