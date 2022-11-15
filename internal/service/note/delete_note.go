package note

import (
	"context"

	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
)

func (s *Service) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.Empty, error) {
	err := s.noteRepository.DeleteNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.Empty{}, nil
}
