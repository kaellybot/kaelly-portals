package mappers

import (
	"time"

	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-portals/models"
	"github.com/kaellybot/kaelly-portals/models/constants"
	"github.com/kaellybot/kaelly-portals/payloads/dofusportals"
	"github.com/kaellybot/kaelly-portals/services/areas"
	"github.com/kaellybot/kaelly-portals/services/dimensions"
	"github.com/kaellybot/kaelly-portals/services/servers"
	"github.com/kaellybot/kaelly-portals/services/subareas"
	"github.com/kaellybot/kaelly-portals/services/transports"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapPortal(portal dofusportals.Portal, serverService servers.Service,
	dimensionService dimensions.Service, areaService areas.Service,
	Service subareas.Service, transportService transports.Service,
) *amqp.PortalPositionAnswer_PortalPosition {
	var remainingUses int32 = 0
	if portal.RemainingUses != nil {
		remainingUses = int32(*portal.RemainingUses)
	}

	return &amqp.PortalPositionAnswer_PortalPosition{
		ServerId:      getInternalServerID(portal.Server, serverService),
		DimensionId:   getInternalDimensionID(portal.Dimension, dimensionService),
		Position:      mapPosition(portal.Position, areaService, Service, transportService),
		RemainingUses: remainingUses,
		CreatedBy:     mapUser(portal.CreatedBy),
		CreatedAt:     mapTimestamp(portal.CreatedAt),
		UpdatedBy:     mapUser(portal.UpdatedBy),
		UpdatedAt:     mapTimestamp(portal.UpdatedAt),
		Source:        mapSource(models.SourceDofusPortals),
	}
}

func mapPosition(position *dofusportals.Position, areaService areas.Service,
	Service subareas.Service, transportService transports.Service,
) *amqp.PortalPositionAnswer_PortalPosition_Position {
	if position == nil {
		return nil
	}

	isInCanopy := false
	if position.IsInCanopy != nil && *position.IsInCanopy {
		isInCanopy = true
	}

	return &amqp.PortalPositionAnswer_PortalPosition_Position{
		X:                    int32(position.X),
		Y:                    int32(position.Y),
		IsInCanopy:           isInCanopy,
		Transport:            mapTransport(position.Transport, areaService, Service, transportService),
		ConditionalTransport: mapTransport(position.ConditionalTransport, areaService, Service, transportService),
	}
}

func mapTransport(transport *dofusportals.Transport, areaService areas.Service,
	Service subareas.Service, transportService transports.Service,
) *amqp.PortalPositionAnswer_PortalPosition_Position_Transport {
	if transport == nil {
		return nil
	}

	return &amqp.PortalPositionAnswer_PortalPosition_Position_Transport{
		AreaId:    getInternalAreaID(transport.Area, areaService),
		SubAreaId: getInternalSubAreaID(transport.SubArea, Service),
		TypeId:    getInternalTransportTypeID(string(transport.Type), transportService),
		X:         int32(transport.X),
		Y:         int32(transport.Y),
	}
}

func mapUser(user *dofusportals.User) string {
	if user == nil {
		return ""
	}

	return user.Name
}

func mapTimestamp(timestamp *time.Time) *timestamppb.Timestamp {
	if timestamp == nil {
		return nil
	}

	return timestamppb.New(*timestamp)
}

func mapSource(source models.Source) *amqp.PortalPositionAnswer_PortalPosition_Source {
	return &amqp.PortalPositionAnswer_PortalPosition_Source{
		Name: source.Name,
		Icon: source.Icon,
		Url:  source.Url,
	}
}

func getInternalServerID(dofusPortalsID string, serverService servers.Service) string {
	server, found := serverService.FindServerByDofusPortalsID(dofusPortalsID)
	if found {
		return server.ID
	}

	log.Warn().Str(constants.LogServerID, dofusPortalsID).
		Msgf("Server not found with following dofusPortalsID, using it as internal one")
	return dofusPortalsID
}

func getInternalDimensionID(dofusPortalsID string, dimensionService dimensions.Service) string {
	dimension, found := dimensionService.FindDimensionByDofusPortalsID(dofusPortalsID)
	if found {
		return dimension.ID
	}

	log.Warn().Str(constants.LogDimensionID, dofusPortalsID).
		Msgf("Dimension not found with following dofusPortalsID, using it as internal one")
	return dofusPortalsID
}

func getInternalAreaID(dofusPortalsID string, areaService areas.Service) string {
	area, found := areaService.FindAreaByDofusPortalsID(dofusPortalsID)
	if found {
		return area.ID
	}

	log.Warn().Str(constants.LogAreaID, dofusPortalsID).
		Msgf("Area not found with following dofusPortalsID, using it as internal one")
	return dofusPortalsID
}

func getInternalSubAreaID(dofusPortalsID string, Service subareas.Service) string {
	subArea, found := Service.FindSubAreaByDofusPortalsID(dofusPortalsID)
	if found {
		return subArea.ID
	}

	log.Warn().Str(constants.LogSubAreaID, dofusPortalsID).
		Msgf("SubArea not found with following dofusPortalsID, using it as internal one")
	return dofusPortalsID
}

func getInternalTransportTypeID(dofusPortalsID string, transportService transports.Service) string {
	transportType, found := transportService.FindTransportTypeByDofusPortalsID(dofusPortalsID)
	if found {
		return transportType.ID
	}

	log.Warn().Str(constants.LogTransportTypeID, dofusPortalsID).
		Msgf("TransportType not found with following dofusPortalsID, using it as internal one")
	return dofusPortalsID
}
