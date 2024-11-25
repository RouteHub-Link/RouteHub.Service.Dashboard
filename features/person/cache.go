package person

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/features"
	"github.com/redis/go-redis/v9"
)

const (
	personCachePrefix       = "person"
	errGetCacheClientFailed = "failed to get cache client"
	errGetPersonFromCache   = "failed to get person from cache"
	errUnmarshalPerson      = "failed to unmarshal person"
)

var (
	ErrGetCacheClientFailed = errors.New(errGetCacheClientFailed)
	ErrGetPersonFromCache   = errors.New(errGetPersonFromCache)
	ErrUnmarshalPerson      = errors.New(errUnmarshalPerson)
)

func GetCachedPerson(ctx context.Context, subjectID string) (*ent.Person, bool, error) {
	rdb, err := features.GetCacheClient()
	if err != nil {
		return nil, false, ErrGetCacheClientFailed
	}

	personKey := strings.Join([]string{personCachePrefix, subjectID}, "-")

	val, err := rdb.Get(ctx, personKey).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, true, nil
		}
		return nil, false, ErrGetPersonFromCache
	}

	var person ent.Person
	if err = json.Unmarshal([]byte(val), &person); err != nil {
		return nil, false, ErrUnmarshalPerson
	}

	return &person, false, nil
}

func SetCachedPerson(ctx context.Context, p *ent.Person) error {
	rdb, err := features.GetCacheClient()
	if err != nil {
		return ErrGetCacheClientFailed
	}

	personKey := strings.Join([]string{personCachePrefix, p.SubjectID}, "-")

	personJSON, err := json.Marshal(p)
	if err != nil {
		return ErrUnmarshalPerson
	}

	rdb.Set(ctx, personKey, personJSON, 0).Err()

	return err
}
