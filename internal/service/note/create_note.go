package note

import (
	"context"

	"github.com/almira-galeeva/note-service-api/internal/model"
)

func (s *Service) CreateNote(ctx context.Context, noteBody *model.NoteBody) (int64, error) {
	id, err := s.noteRepository.CreateNote(ctx, noteBody)
	if err != nil {
		return 0, err
	}

	return id, nil
}
