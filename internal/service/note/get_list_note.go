package note

import (
	"context"

	"github.com/almira-galeeva/note-service-api/internal/model"
)

func (s *Service) GetListNote(ctx context.Context, ids []int64) ([]*model.Note, error) {
	notes, err := s.noteRepository.GetListNote(ctx, ids)
	if err != nil {
		return nil, err
	}

	return notes, nil
}
