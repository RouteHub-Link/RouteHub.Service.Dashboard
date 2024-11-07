package organization

import (
	"context"

	"RouteHub.Service.Dashboard/ent"
)

func CreateOrganization(client *ent.Client, p *ent.Person, name string, ctx context.Context) (o *ent.Organization, err error) {
	o = client.Organization.Create().
		SetName(name).
		SaveX(ctx)

	if err = p.Update().SetOrganization(o).Exec(ctx); err != nil {
		return
	}

	return

}
