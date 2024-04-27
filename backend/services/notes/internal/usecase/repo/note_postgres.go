package repo

import (
	"context"
	"time"

	"github.com/Ixorlive/tw_vk/backend/services/notes/internal/entity"
	"github.com/Ixorlive/tw_vk/backend/services/notes/internal/usecase"
	"github.com/Ixorlive/tw_vk/backend/services/notes/pkg/postgres"
)

type PGNoteRepo struct {
	*postgres.Postgres
}

func NewPGNoteRepo(pg *postgres.Postgres) usecase.NoteRepo {
	return &PGNoteRepo{pg}
}

func (repo *PGNoteRepo) CreateNote(ctx context.Context, note entity.Note) (entity.Note, error) {
	sql, args, err := repo.Builder.Insert("Notes").
		Columns("user_id", "content").
		Values(note.UserID, note.Content).
		Suffix("RETURNING id, created_at, updated_at").
		ToSql()

	if err != nil {
		return entity.Note{}, err
	}

	row := repo.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&note.ID, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return entity.Note{}, err
	}

	return note, nil
}

func (repo *PGNoteRepo) FindNoteByID(ctx context.Context, id int) (entity.Note, error) {
	sql, args, err := repo.Builder.
		Select("id", "user_id", "content", "created_at", "updated_at").
		From("Notes").
		Where("id = ?", id).
		ToSql()

	if err != nil {
		return entity.Note{}, err
	}

	var note entity.Note
	row := repo.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&note.ID, &note.UserID, &note.Content, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return entity.Note{}, err
	}

	return note, nil
}

func (repo *PGNoteRepo) UpdateNote(ctx context.Context, note entity.Note) (entity.Note, error) {
	sql, args, err := repo.Builder.Update("Notes").
		Set("content", note.Content).
		Set("updated_at", "CURRENT_TIMESTAMP").
		Where("id = ?", note.ID).
		Suffix("RETURNING id, user_id, content, created_at, updated_at").
		ToSql()

	if err != nil {
		return entity.Note{}, err
	}

	row := repo.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&note.ID, &note.UserID, &note.Content, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return entity.Note{}, err
	}

	return note, nil
}

func (repo *PGNoteRepo) DeleteNote(ctx context.Context, id int) error {
	sql, args, err := repo.Builder.Delete("Notes").
		Where("id = ?", id).
		ToSql()

	if err != nil {
		return err
	}

	_, err = repo.Pool.Exec(ctx, sql, args...)
	return err
}

func (repo *PGNoteRepo) FindNotesByUserID(ctx context.Context, userID int) ([]entity.Note, error) {
	sql, args, err := repo.Builder.
		Select("id", "user_id", "content", "created_at", "updated_at").
		From("Notes").
		Where("user_id = ?", userID).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := repo.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []entity.Note
	for rows.Next() {
		var note entity.Note
		if err := rows.Scan(&note.ID, &note.UserID, &note.Content, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (repo *PGNoteRepo) FindNotesByDateRange(ctx context.Context, start, end time.Time) ([]entity.Note, error) {
	sql, args, err := repo.Builder.
		Select("id", "user_id", "content", "created_at", "updated_at").
		From("Notes").
		Where("created_at >= ? AND updated_at <= ?", start, end).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := repo.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []entity.Note
	for rows.Next() {
		var note entity.Note
		if err := rows.Scan(&note.ID, &note.UserID, &note.Content, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}
