package note_v1

import (
	"context"

	"github.com/almira-galeeva/note-service-api/internal/converter"
	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
)

func (n *Note) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	id, err := n.noteService.CreateNote(ctx, converter.ToNote(req.GetNoteBody()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateNoteResponse{
		Id: id,
	}, nil
}
