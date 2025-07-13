package gapi

import (
	"strconv"
	"time"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/sse/eventdata"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/models"
	"google.golang.org/protobuf/types/known/durationpb"
)

func convertAction(action db.Action) *models.Action {
	return &models.Action{
		Id:           action.ID,
		Name:         action.Name,
		Description:  action.Description.String,
		Enabled:      action.Enabled,
		CanBeDeleted: action.CanBeDeleted,
	}
}

func convertActionEvent(action db.Action) *eventdata.Action {
	return &eventdata.Action{
		Id:           strconv.FormatInt(action.ID, 10),
		Name:         action.Name,
		Description:  action.Description.String,
		Enabled:      action.Enabled,
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

func convertFacilitiesGetRows(facilities []db.GetFacilitiesRow) []*models.Facility {
	result := make([]*models.Facility, len(facilities))

	for i, v := range facilities {
		result[i] = convertFacilitiesGetRow(v)
	}

	return result
}

func convertFacilitiesGetRow(facility db.GetFacilitiesRow) *models.Facility {
	return &models.Facility{
		Id:                facility.ID,
		Name:              facility.Name,
		Description:       facility.Description.String,
		OpenTime:          durationpb.New(time.Duration(facility.OpenTime.Microseconds)),
		CloseTime:         durationpb.New(time.Duration(facility.CloseTime.Microseconds)),
		AddressCountry:    facility.Country.String,
		AddressState:      facility.State.String,
		AddressSubState:   facility.SubState.String,
		AddressStreet:     facility.Street.String,
		AddressNumber:     facility.Number.String,
		AddressUnit:       facility.Unit.String,
		AddressPostalcode: facility.PostalCode.String,
		AddressLat:        facility.Lat.Float64,
		AddressLng:        facility.Lng.Float64,
	}
}

func convertFacilityGetRow(facility db.GetFacilityRow) *models.Facility {
	return &models.Facility{
		Id:                facility.ID,
		Name:              facility.Name,
		Description:       facility.Description.String,
		OpenTime:          durationpb.New(time.Duration(facility.OpenTime.Microseconds)),
		CloseTime:         durationpb.New(time.Duration(facility.CloseTime.Microseconds)),
		AddressCountry:    facility.Country.String,
		AddressState:      facility.State.String,
		AddressSubState:   facility.SubState.String,
		AddressStreet:     facility.Street.String,
		AddressNumber:     facility.Number.String,
		AddressUnit:       facility.Unit.String,
		AddressPostalcode: facility.PostalCode.String,
		AddressLat:        facility.Lat.Float64,
		AddressLng:        facility.Lng.Float64,
	}
}

func convertFacilityCreateTxResult(facility db.CreateFacilityTxResult) *models.Facility {
	return &models.Facility{
		Id:                facility.Facility.ID,
		Name:              facility.Facility.Name,
		Description:       facility.Facility.Description.String,
		OpenTime:          durationpb.New(time.Duration(facility.Facility.OpenTime.Microseconds)),
		CloseTime:         durationpb.New(time.Duration(facility.Facility.CloseTime.Microseconds)),
		AddressCountry:    facility.FacilityAddress.Country,
		AddressState:      facility.FacilityAddress.State,
		AddressSubState:   facility.FacilityAddress.SubState.String,
		AddressStreet:     facility.FacilityAddress.Street,
		AddressNumber:     facility.FacilityAddress.Number,
		AddressUnit:       facility.FacilityAddress.Unit.String,
		AddressPostalcode: facility.FacilityAddress.PostalCode,
		AddressLat:        facility.FacilityAddress.Lat.Float64,
		AddressLng:        facility.FacilityAddress.Lng.Float64,
	}
}

func convertFacilityUpdateTxResult(facility db.UpdateFacilityTxResult) *models.Facility {
	return &models.Facility{
		Id:                facility.Facility.ID,
		Name:              facility.Facility.Name,
		Description:       facility.Facility.Description.String,
		OpenTime:          durationpb.New(time.Duration(facility.Facility.OpenTime.Microseconds)),
		CloseTime:         durationpb.New(time.Duration(facility.Facility.CloseTime.Microseconds)),
		AddressCountry:    facility.FacilityAddress.Country,
		AddressState:      facility.FacilityAddress.State,
		AddressSubState:   facility.FacilityAddress.SubState.String,
		AddressStreet:     facility.FacilityAddress.Street,
		AddressNumber:     facility.FacilityAddress.Number,
		AddressUnit:       facility.FacilityAddress.Unit.String,
		AddressPostalcode: facility.FacilityAddress.PostalCode,
		AddressLat:        facility.FacilityAddress.Lat.Float64,
		AddressLng:        facility.FacilityAddress.Lng.Float64,
	}
}

func convertCreateFacilityTxResultEvent(facility db.CreateFacilityTxResult) *eventdata.Facility {
	return &eventdata.Facility{
		Id:                strconv.FormatInt(facility.Facility.ID, 10),
		Name:              facility.Facility.Name,
		Description:       facility.Facility.Description.String,
		OpenTime:          durationpb.New(time.Duration(facility.Facility.OpenTime.Microseconds)),
		CloseTime:         durationpb.New(time.Duration(facility.Facility.CloseTime.Microseconds)),
		AddressCountry:    facility.FacilityAddress.Country,
		AddressState:      facility.FacilityAddress.State,
		AddressSubState:   facility.FacilityAddress.SubState.String,
		AddressStreet:     facility.FacilityAddress.Street,
		AddressNumber:     facility.FacilityAddress.Number,
		AddressUnit:       facility.FacilityAddress.Unit.String,
		AddressPostalcode: facility.FacilityAddress.PostalCode,
		AddressLat:        facility.FacilityAddress.Lat.Float64,
		AddressLng:        facility.FacilityAddress.Lng.Float64,
	}
}

func convertUpdateFacilityTxResultEvent(facility db.UpdateFacilityTxResult) *eventdata.Facility {
	return &eventdata.Facility{
		Id:                strconv.FormatInt(facility.Facility.ID, 10),
		Name:              facility.Facility.Name,
		Description:       facility.Facility.Description.String,
		OpenTime:          durationpb.New(time.Duration(facility.Facility.OpenTime.Microseconds)),
		CloseTime:         durationpb.New(time.Duration(facility.Facility.CloseTime.Microseconds)),
		AddressCountry:    facility.FacilityAddress.Country,
		AddressState:      facility.FacilityAddress.State,
		AddressSubState:   facility.FacilityAddress.SubState.String,
		AddressStreet:     facility.FacilityAddress.Street,
		AddressNumber:     facility.FacilityAddress.Number,
		AddressUnit:       facility.FacilityAddress.Unit.String,
		AddressPostalcode: facility.FacilityAddress.PostalCode,
		AddressLat:        facility.FacilityAddress.Lat.Float64,
		AddressLng:        facility.FacilityAddress.Lng.Float64,
	}
}
