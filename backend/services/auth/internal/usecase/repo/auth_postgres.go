package repo

import (
	"fmt"

	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/entity"
	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/usecase"
	"github.com/Ixorlive/tw_vk/backend/services/auth/pkg/postgres"
	"github.com/Masterminds/squirrel"
	"golang.org/x/net/context"
)

type PGUserRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) usecase.UserRepo {
	return &PGUserRepo{pg}
}

func (r *PGUserRepo) FindByLogin(ctx context.Context, login string) (entity.User, error) {
	query, _, err := r.Builder.
		Select("id", "login", "password").
		From("Users").
		Where(squirrel.Eq{"login": login}).
		ToSql()

	var user entity.User

	if err != nil {
		return user, fmt.Errorf("error building query: %w", err)
	}
	rows, err := r.Pool.Query(ctx, query)
	if err != nil {
		return user, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	// user can be not found - return empty User
	for rows.Next() {
		var lg, password string
		var id uint64
		err = rows.Scan(&id, &lg, &password)
		if err != nil {
			return user, fmt.Errorf("error scanning row: %w", err)
		}
		user.Id = id
		user.Login = lg
		user.Password = password
		break
	}
	if err := rows.Err(); err != nil {
		return user, fmt.Errorf("error reading rows: %w", err)
	}

	return user, nil
}

func (r *PGUserRepo) AddUser(ctx context.Context, user entity.User) (bool, error) {
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
