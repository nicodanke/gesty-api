package gapi

import (
	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/account-service/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *models.User {
	return &models.User{
		Id:                user.ID,
		Username:          user.Username,
		Name:              user.Name,
		Lastname:          user.Lastname,
		Email:             user.Email,
		Phone:             user.Phone.String,
		Active:            user.Active,
		IsAdmin:           user.IsAdmin,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}

func convertUsers(users []db.User) []*models.User {
	result := make([]*models.User, len(users))

	for i, v := range users {
		result[i] = convertUser(v)
	}

	return result
}

func convertAccount(user db.Account) *models.Account {
	return &models.Account{
		Id:          user.ID,
		Code:        user.Code,
		CompanyName: user.CompanyName,
		Phone:       user.Phone.String,
		Email:       user.Email,
		WebUrl:      user.WebUrl.String,
		Active:      user.Active,
		CreatedAt:   timestamppb.New(user.CreatedAt),
	}
}

func convertRole(role db.Role) *models.Role {
	return &models.Role{
		Id:   role.ID,
		Name: role.Name,
	}
}

func convertRoles(roles []db.Role) []*models.Role {
	result := make([]*models.Role, len(roles))

	for i, v := range roles {
		result[i] = convertRole(v)
	}

	return result
}
