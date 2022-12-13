package note_v1

import (
	"context"

	"github.com/almira-galeeva/note-service-api/internal/converter"
	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
)

func (n *Note) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {
	res, err := n.noteService.GetListNote(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}

	return &desc.GetListNoteResponse{
		Results: converter.ToDescListWholeNote(res),
	}, nil
}
