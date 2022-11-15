package note

import (
	"context"

	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
)

func (s *Service) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	id, err := s.noteRepository.CreateNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.CreateNoteResponse{
		Id: id,
	}, nil
}
