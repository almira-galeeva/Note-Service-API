package note

import (
	"context"
)

func (s *Service) DeleteNote(ctx context.Context, id int64) error {
	err := s.noteRepository.DeleteNote(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
