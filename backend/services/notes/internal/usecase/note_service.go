package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Ixorlive/tw_vk/backend/services/notes/internal/entity"
)

type BasicNoteService struct {
	repo NoteRepo
}

func NewNoteService(repo NoteRepo) NoteService {
	return &BasicNoteService{
		repo: repo,
	}
}

func (s *BasicNoteService) CreateNote(ctx context.Context, userID int, content string) (entity.Note, error) {
	note := entity.Note{
		UserID:  userID,
		Content: content,
	}
	return s.repo.CreateNote(ctx, note)
}

func (s *BasicNoteService) UpdateNote(ctx context.Context, noteID int, userID int, content string) (entity.Note, error) {
	note, err := s.repo.FindNoteByID(ctx, noteID)
	if err != nil {
		return entity.Note{}, err
	}
	if note.UserID != userID {
		return entity.Note{}, errors.New("unauthorized to update this note")
	}
	note.Content = content
	return s.repo.UpdateNote(ctx, note)
}

func (s *BasicNoteService) DeleteNote(ctx context.Context, noteID int, userID int) error {
	note, err := s.repo.FindNoteByID(ctx, noteID)
	if err != nil {
		return err
	}
	if note.UserID != userID {
		return errors.New("unauthorized to delete this note")
	}
	return s.repo.DeleteNote(ctx, noteID)
}

func (s *BasicNoteService) GetNoteByID(ctx context.Context, noteID int, userID int) (entity.Note, error) {
	return s.repo.FindNoteByID(ctx, noteID)
}

func (s *BasicNoteService) GetNotesByUserID(ctx context.Context, userID int) ([]entity.Note, error) {
	return s.repo.FindNotesByUserID(ctx, userID)
}

func (s *BasicNoteService) GetNotesByDateRange(ctx context.Context, userID int, start, end time.Time) ([]entity.Note, error) {
	return s.repo.FindNotesByDateRange(ctx, start, end)
}
