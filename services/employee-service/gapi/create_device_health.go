package gapi

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/device_health"
	"github.com/rs/zerolog/log"
)

func (server *Server) CreateDeviceHealth(ctx context.Context, req *device_health.CreateDeviceHealthRequest) (*device_health.CreateDeviceHealthResponse, error) {
	log.Info().Str("method", "CreateDeviceHealth").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing CreateDeviceHealth request")

	authPayloadDevice, err := server.authenticateDevice(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	fmt.Println("authPayloadDevice", authPayloadDevice)

	arg := db.CreateDeviceHealthParams{
		DeviceID:        req.GetId(),
		ConnectionType:  req.GetConnectionType(),
		FreeMemory:      pgtype.Float8{Float64: req.GetFreeRam(), Valid: true},
		FreeStorage:     pgtype.Float8{Float64: req.GetFreeStorage(), Valid: true},
		BatteryLevel:    pgtype.Float8{Float64: req.GetBatteryLevel(), Valid: true},
		BatterySaveMode: req.GetBatterySaveMode(),
	}

	_, err = server.store.CreateDeviceHealth(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_UNIQUE, fmt.Sprintln("Failed to create device health due to unique constraint violation"), constraintName)
		}
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to create device health due to foreign key constraint violation"), constraintName)
		}
		return nil, internalError(fmt.Sprintln("Failed to create device health", err))
	}

	actions, err := server.store.GetActionsEnabledByDeviceId(ctx, authPayloadDevice.DeviceID)

	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to get actions enabled by device id", err))
	}

	rsp := &device_health.CreateDeviceHealthResponse{
		Actions: convertActions(actions),
	}

	return rsp, nil
}
