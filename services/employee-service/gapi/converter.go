package gapi

import (
	"strconv"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/sse/eventdata"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/models"
)

func convertAction(action db.Action) *models.Action {
	return &models.Action{
		Id:          action.ID,
		Name:        action.Name,
		Description: action.Description.String,
		Enabled:     action.Enabled,
		CanBeDeleted: action.CanBeDeleted,
	}
}

func convertActionEvent(action db.Action) *eventdata.Action {
	return &eventdata.Action{
		Id:          strconv.FormatInt(action.ID, 10),
		Name:        action.Name,
		Description: action.Description.String,
		Enabled:     action.Enabled,
		CanBeDeleted: action.CanBeDeleted,
	}
}

func convertActions(actions []db.Action) []*models.Action {
	result := make([]*models.Action, len(actions))

	for i, v := range actions {
		result[i] = convertAction(v)
	}

	return result
}
