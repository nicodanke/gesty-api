package gapi

import (
	"slices"

	"github.com/nicodanke/gesty-api/shared/token"
)

func (server *Server) authorizeUser(payload *token.Payload, permissions [][]string) bool {
	payloadPermissions := payload.Permissions

	if slices.Contains(payloadPermissions, "A") {
		return true
	}

	// For each permission group in permissions
	for _, permissionGroup := range permissions {
		// Check if at least one permission from the group exists in payload
		hasOnePermission := false
		for _, permission := range permissionGroup {
			if slices.Contains(payloadPermissions, permission) {
				hasOnePermission = true
				break
			}
		}
		
		// If no permission from this group was found, return false
		if !hasOnePermission {
			return false
		}
	}

	return true
}