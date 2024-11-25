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

const (
	errUserInfoNil          = "user info is nil"
	errUserSubject          = "user subject is empty"
	errCheckingUserInfo     = "checking user info"
	errCreatingOrganization = "creating new organization"
	errCreatingNewPerson    = "creating new person"
	errCheckingPersonExists = "checking if person exists"
	errRetrievingPerson     = "retrieving person"

	logErrorCheckingPerson   = "Error checking if person exists"
	logErrorCreatingPerson   = "Error creating new person"
	logErrorRetrievingPerson = "Error retrieving person"
	logPersonDoesNotExist    = "Person does not exist, creating new person"
)

var (
	ErrUserInfoNil          = errors.New(errUserInfoNil)
	ErrUserSubject          = errors.New(errUserSubject)
	ErrCheckingUserInfo     = errors.New(errCheckingUserInfo)
	ErrCreatingOrganization = errors.New(errCreatingOrganization)
	ErrCheckingPersonExists = errors.New(errCheckingPersonExists)
	ErrCreatingNewPerson    = errors.New(errCreatingNewPerson)
	ErrRetrievingPerson     = errors.New(errRetrievingPerson)
)

func GetPerson(ctx context.Context, userInfo *oidc.UserInfo, client *ent.Client, logger *slog.Logger) (*ent.Person, error) {
	if err := checkUserInfo(userInfo); err != nil {
		return nil, ErrCheckingUserInfo
	}

	exists, err := client.Person.Query().Where(person.SubjectID(userInfo.Subject)).Exist(ctx)
	if err != nil {
		logger.ErrorContext(ctx, errCheckingPersonExists, "error", err, "userInfo", userInfo)
		return nil, ErrCheckingPersonExists
	}

	var p *ent.Person
	if !exists {
		logger.InfoContext(ctx, logPersonDoesNotExist)
		p, err = createNewPerson(ctx, userInfo, client, logger)
		if err != nil {
			return nil, err
		}
	} else {
		p, err = retrieveExistingPerson(ctx, userInfo, client, logger)
		if err != nil {
			return nil, err
		}
	}

	p, err = checkOrganization(ctx, client, p, logger)
	if err != nil {
		return nil, err
	}

	return p, nil
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

func createNewPerson(ctx context.Context, userInfo *oidc.UserInfo, client *ent.Client, logger *slog.Logger) (*ent.Person, error) {
	newPerson, err := client.Person.Create().
		SetSubjectID(userInfo.Subject).
		SetUserInfo(*userInfo).
		Save(ctx)

	if err != nil {
		logger.ErrorContext(ctx, errCreatingNewPerson, "error", err, "userInfo", userInfo)
		return nil, ErrCreatingNewPerson
	}

	return newPerson, nil
}

func retrieveExistingPerson(ctx context.Context, userInfo *oidc.UserInfo, client *ent.Client, logger *slog.Logger) (*ent.Person, error) {
	person, err := client.Person.Query().Where(person.SubjectID(userInfo.Subject)).WithOrganizations().Only(ctx)
	if err != nil {
		logger.ErrorContext(ctx, errRetrievingPerson, "error", err)
		return nil, ErrRetrievingPerson
	}

	return person, nil
}

func checkOrganization(ctx context.Context, client *ent.Client, person *ent.Person, logger *slog.Logger) (*ent.Person, error) {
	if person.Edges.Organizations == nil || len(person.Edges.Organizations) == 0 {
		return createOrganization(ctx, client, person, logger)
	}

	return person, nil
}

func createOrganization(ctx context.Context, client *ent.Client, person *ent.Person, logger *slog.Logger) (*ent.Person, error) {
	_, err := organization.CreateOrganization(ctx, client, person, person.UserInfo.Name+"'s Organization")
	if err != nil {
		logger.ErrorContext(ctx, errCreatingOrganization, "error", err)
		return nil, ErrCreatingOrganization
	}

	return person, nil
}
