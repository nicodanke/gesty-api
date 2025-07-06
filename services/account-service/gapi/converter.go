package gapi

import (
	"strconv"

	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/account-service/sse/eventdata"
	"github.com/nicodanke/gesty-api/shared/proto/account-service/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertPermission(permission db.Permission) *models.Permission {
	return &models.Permission{
		Id:       permission.ID,
		Code:     permission.Code,
		ParentId: permission.ParentID.Int64,
	}
}

func convertPermissions(permissions []db.Permission) []*models.Permission {
	result := make([]*models.Permission, len(permissions))

	for i, v := range permissions {
		result[i] = convertPermission(v)
	}

	return result
}

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
		RoleId:            user.RoleID,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}

func convertUserEvent(user db.User) *eventdata.User {
	return &eventdata.User{
		Id:                strconv.FormatInt(user.ID, 10),
		Username:          user.Username,
		Name:              user.Name,
		Lastname:          user.Lastname,
		Email:             user.Email,
		Phone:             user.Phone.String,
		Active:            user.Active,
		IsAdmin:           user.IsAdmin,
		RoleId:            strconv.FormatInt(user.RoleID, 10),
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
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

func convertRoleCreateEvent(role db.CreateRoleTxResult) *eventdata.Role {
	permissionIds := make([]string, 0)
	for _, v := range role.PermissionIDs {
		permissionIds = append(permissionIds, strconv.FormatInt(v, 10))
	}

	return &eventdata.Role{
		Id:            strconv.FormatInt(role.Role.ID, 10),
		Name:          role.Role.Name,
		Description:   role.Role.Description.String,
		PermissionIds: permissionIds,
	}
}

func convertRoleUpdateEvent(role db.UpdateRoleTxResult) *eventdata.Role {
	permissionIds := make([]string, 0)
	for _, v := range role.PermissionIDs {
		permissionIds = append(permissionIds, strconv.FormatInt(v, 10))
	}

	return &eventdata.Role{
		Id:            strconv.FormatInt(role.Role.ID, 10),
		Name:          role.Role.Name,
		Description:   role.Role.Description.String,
		PermissionIds: permissionIds,
	}
}

func convertRolesRowEach(role db.GetRolesRow) *models.Role {
	permissionIds := make([]int64, 0)
	for _, v := range role.PermissionIds.([]interface{}) {
		permissionIds = append(permissionIds, v.(int64))
	}

	return &models.Role{
		Id:            role.ID,
		Name:          role.Name,
		Description:   role.Description.String,
		PermissionIds: permissionIds,
	}
}

func convertRolesRow(roles []db.GetRolesRow) []*models.Role {
	result := make([]*models.Role, len(roles))

	for i, v := range roles {
		result[i] = convertRolesRowEach(v)
	}

	return result
}

func convertRoleRow(role db.GetRoleRow) *models.Role {
	permissionIds := make([]int64, 0)
	for _, v := range role.PermissionIds.([]interface{}) {
		permissionIds = append(permissionIds, v.(int64))
	}

	return &models.Role{
		Id:            role.ID,
		Name:          role.Name,
		Description:   role.Description.String,
		PermissionIds: permissionIds,
	}
}

func convertRoleCreate(role db.CreateRoleTxResult) *models.Role {
	return &models.Role{
		Id:            role.Role.ID,
		Name:          role.Role.Name,
		Description:   role.Role.Description.String,
		PermissionIds: role.PermissionIDs,
	}
}

func convertRoleUpdate(role db.UpdateRoleTxResult) *models.Role {
	return &models.Role{
		Id:            role.Role.ID,
		Name:          role.Role.Name,
		Description:   role.Role.Description.String,
		PermissionIds: role.PermissionIDs,
	}
}
