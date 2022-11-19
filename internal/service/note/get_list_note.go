package note

import (
	"context"

	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
)

func (s *Service) GetListNote(ctx context.Context, ids []int64) (*desc.GetListNoteResponse, error) {
	notes, err := s.noteRepository.GetListNote(ctx, ids)
	if err != nil {
		return &desc.GetListNoteResponse{}, err
	}

	var res []*desc.WholeNote
	for i := 0; i < len(notes); i++ {
		res = append(res, &desc.WholeNote{
			Id: notes[i].Id,
			NoteBody: &desc.NoteBody{
				Title:  notes[i].Title,
				Text:   notes[i].Text,
				Author: notes[i].Author,
				Email:  notes[i].Email,
			},
			CreatedAt: notes[i].TsCreatedAt,
			UpdatedAt: notes[i].TsUpdatedAt,
		})
	}

	return &desc.GetListNoteResponse{
		Results: res,
	}, nil
}
