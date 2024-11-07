package person

import (
	"context"
	"errors"
	"log/slog"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/ent/person"
	"RouteHub.Service.Dashboard/features/organization"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

var (
	ErrUserInfoNil = errors.New("user info is nil")
	ErrUserSubject = errors.New("user subject is empty")
)

func GetPerson(userInfo *oidc.UserInfo, client *ent.Client, ctx context.Context, logger *slog.Logger) (p *ent.Person, err error) {
	if err = checkUserInfo(userInfo); err != nil {
		return
	}

	_person := ent.Person{}

	checkPersonExists, err := client.Person.Query().Where(person.SubjectID(userInfo.Subject)).Exist(ctx)
	if err != nil {
		logger.ErrorContext(ctx, "Error checking if person exists", "error", err)
		return nil, err
	}
	if !checkPersonExists {
		logger.Info("Person does not exist, creating new person")

		_person := client.Person.Create().
			SetSubjectID(userInfo.Subject).
			SetUserInfo(*userInfo)

		_, err = _person.Save(ctx)

		if err != nil {
			logger.ErrorContext(ctx, "Error creating person", "error", err)
			return nil, err
		}

		logger.Info("Person created", "data", _person)
	}

	if _person.Edges.Organization == nil {
		organization.CreateOrganization(client, &_person, "My Organization", ctx)
		return
	}

	return
}

func checkUserInfo(userInfo *oidc.UserInfo) error {
	if userInfo == nil {
		return ErrUserInfoNil
	}

	if userInfo.Subject == "" {
		return ErrUserSubject
	}
	return nil
}
