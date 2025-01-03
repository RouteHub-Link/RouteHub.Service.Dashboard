package layouts

import (
	"RouteHub.Service.Dashboard/ent"
	"github.com/a-h/templ"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

type PageDescription struct {
	AdditionalHead   *templ.Component
	MainContent      templ.Component
	AdditionalFooter *templ.Component
	UserInfo         *oidc.UserInfo
	Hub              *ent.Hub
}

func (pd PageDescription) GetAdditionalFooter() templ.Component {
	return *pd.AdditionalFooter
}

func (pd PageDescription) GetAdditionalHead() templ.Component {
	return *pd.AdditionalHead
}
