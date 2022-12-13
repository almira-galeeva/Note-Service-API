package note_v1

import (
	"context"

	"github.com/almira-galeeva/note-service-api/internal/converter"
	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	res, err := n.noteService.GetNote(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetNoteResponse{
		WholeNote: converter.ToDescWholeNote(res),
	}, nil
}
