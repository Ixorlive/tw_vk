package repo

import (
	"fmt"

	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/entity"
	"github.com/Ixorlive/tw_vk/backend/services/auth/pkg/postgres"
	"github.com/Masterminds/squirrel"
	"golang.org/x/net/context"
)

type UserRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (r *UserRepo) UserExists(ctx context.Context, user entity.User, onlyLogin bool) (bool, error) {
	builder := r.Builder.
		Select("1").
		From("User").
		Where(squirrel.Eq{"login": user.Login})

	if !onlyLogin {
		builder = builder.Where(squirrel.Eq{"password": user.Password})
	}
	query, _, err := builder.Limit(1).ToSql()

	if err != nil {
		return false, fmt.Errorf("error building query: %w", err)
	}
	rows, err := r.Pool.Query(ctx, query)
	if err != nil {
		return false, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	exists := rows.Next()
	if exists {
		return true, nil
	}
	if err := rows.Err(); err != nil {
		return false, fmt.Errorf("error reading query results: %w", err)
	}
	return false, nil
}

func (r *UserRepo) AddUser(ctx context.Context, user entity.User) (bool, error) {
	query, _, err := r.Builder.Insert("Users").Columns("login", "password").
		Values(user.Login, user.Password).ToSql()
	if err != nil {
		return false, fmt.Errorf("error building query: %w", err)
	}
	_, err = r.Pool.Query(ctx, query)
	if err != nil {
		return false, fmt.Errorf("error executing query: %w", err)
	}
	return true, nil
}
