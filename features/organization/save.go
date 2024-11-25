package organization

import (
	"context"

	"RouteHub.Service.Dashboard/ent"
)

func CreateOrganization(ctx context.Context, client *ent.Client, p *ent.Person, name string) (o *ent.Organization, err error) {
	// Create the organization with the given name.
	o, err = client.Organization.Create().
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// Add the newly created organization to the person's list of organizations.
	err = client.Person.UpdateOne(p).
		AddOrganizations(o). // This is the correct function name for adding a single organization
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return o, nil
}
