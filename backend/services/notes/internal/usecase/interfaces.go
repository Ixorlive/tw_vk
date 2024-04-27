package usecase

import (
	"context"
	"time"

	"github.com/Ixorlive/tw_vk/backend/services/notes/internal/entity"
)

type (
	NoteRepo interface {
		CreateNote(ctx context.Context, note entity.Note) (entity.Note, error)
		UpdateNote(ctx context.Context, note entity.Note) (entity.Note, error)
		DeleteNote(ctx context.Context, id int) error
		FindNoteByID(ctx context.Context, id int) (entity.Note, error)
		FindNotesByUserID(ctx context.Context, userID int) ([]entity.Note, error)
		FindNotesByDateRange(ctx context.Context, start, end time.Time) ([]entity.Note, error)
	}

	NoteService interface {
		CreateNote(ctx context.Context, userID int, content string) (entity.Note, error)
		UpdateNote(ctx context.Context, noteID int, userID int, content string) (entity.Note, error)
		DeleteNote(ctx context.Context, noteID int, userID int) error
		GetNoteByID(ctx context.Context, noteID int, userID int) (entity.Note, error)
		GetNotesByUserID(ctx context.Context, userID int) ([]entity.Note, error)
		GetNotesByDateRange(ctx context.Context, userID int, start, end time.Time) ([]entity.Note, error)
	}
)
