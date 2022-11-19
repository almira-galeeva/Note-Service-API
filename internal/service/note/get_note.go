package note

import (
	"context"

	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
)

func (s *Service) GetNote(ctx context.Context, id int64) (*desc.GetNoteResponse, error) {
	note, err := s.noteRepository.GetNote(ctx, id)
	if err != nil {
		return &desc.GetNoteResponse{}, err
	}

	return &desc.GetNoteResponse{
		WholeNote: &desc.WholeNote{
			Id: note.Id,
			NoteBody: &desc.NoteBody{
				Title:  note.Title,
				Text:   note.Text,
				Author: note.Author,
				Email:  note.Email,
			},
			CreatedAt: note.TsCreatedAt,
			UpdatedAt: note.TsUpdatedAt,
		},
	}, nil
}
