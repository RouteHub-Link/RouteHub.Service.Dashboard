package components

import "RouteHub.Service.Dashboard/ent/schema/enums"

type EditLinkStatusPayload struct {
	Status   enums.StatusState `json:"link_status" form:"link_status"`
	HubSlug  string
	LinkPath string
}
