package note

import (
	"context"

	"github.com/almira-galeeva/note-service-api/internal/model"
)

func (s *Service) GetNote(ctx context.Context, id int64) (*model.WholeNote, error) {
	note, err := s.noteRepository.GetNote(ctx, id)
	if err != nil {
		return &model.WholeNote{}, err
	}

	return note, nil
}
