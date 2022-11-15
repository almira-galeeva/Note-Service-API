package note

import (
	"context"

	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
)

func (s *Service) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {
	notes, err := s.noteRepository.GetListNote(ctx, req)
	if err != nil {
		return &desc.GetListNoteResponse{}, err
	}

	var res []*desc.GetListNoteResponse_Result
	for i := 0; i < len(notes); i++ {
		res = append(res, &desc.GetListNoteResponse_Result{
			Id:        notes[i].Id,
			Title:     notes[i].Title,
			Text:      notes[i].Text,
			Author:    notes[i].Author,
			Email:     notes[i].Email,
			CreatedAt: notes[i].TsCreatedAt,
			UpdatedAt: notes[i].TsUpdatedAt,
		})
	}

	return &desc.GetListNoteResponse{
		Results: res,
	}, nil
}
