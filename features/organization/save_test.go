package organization

import (
	"context"
	"log"
	"testing"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/ent/enttest"
	"RouteHub.Service.Dashboard/ent/person"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

func TestSaveOrganizationFromUserOIDC(t *testing.T) {
	ctx := context.Background()

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	var userInfo *oidc.UserInfo
	userInfo = &oidc.UserInfo{
		Subject:         "292732476601229571",
		UserInfoEmail:   oidc.UserInfoEmail{Email: "test@tester.com", EmailVerified: true},
		UserInfoProfile: oidc.UserInfoProfile{Name: "Test Tester", GivenName: "Test", FamilyName: "Tester"},
	}

	_person := ent.Person{}

	checkPersonExists, _ := client.Person.Query().Where(person.SubjectID(userInfo.Subject)).Exist(ctx)
	if !checkPersonExists {
		t.Logf("Person does not exist, creating new person")

		_person = *client.Person.Create().
			SetSubjectID(userInfo.Subject).
			SetUserInfo(*userInfo).
			SaveX(ctx)

		t.Logf("Person created: %v", _person)
	}

	if _person.SubjectID == "" {
		_person = *client.Person.Query().Where(person.SubjectID(userInfo.Subject)).OnlyX(ctx)
	}

	_person = *client.Person.Query().
		Where(person.SubjectID(userInfo.Subject)).
		WithOrganization().
		OnlyX(ctx)

	if _person.Edges.Organization != nil {
		t.Logf("Person already has organization: %v", _person.Edges.Organization)
		return
	}

	organization := client.Organization.Create().
		SetName("MY ORGANIZATION").
		SaveX(ctx)

	_person.Update().SetOrganization(organization).SaveX(ctx)

	t.Logf("Organization created: %v", organization)

	personWithOrganization := client.Person.Query().
		Where(person.SubjectID(userInfo.Subject)).
		WithOrganization().
		OnlyX(ctx)

	t.Logf("Person with organization: %v", personWithOrganization)

}
