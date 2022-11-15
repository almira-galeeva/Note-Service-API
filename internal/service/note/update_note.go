package note

import (
	"context"

	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
)

func (s *Service) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	id, err := s.noteRepository.UpdateNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.UpdateNoteResponse{
		Id: id,
	}, nil
}
