package note_v1

import (
	"context"

	"github.com/almira-galeeva/note-service-api/internal/converter"
	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	res, err := n.noteService.UpdateNote(ctx, converter.ToUpdateNoteInfo(req))
	if err != nil {
		return nil, err
	}

	return res, nil
}
