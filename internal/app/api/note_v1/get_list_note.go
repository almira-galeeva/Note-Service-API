package note_v1

import (
	"context"

	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
)

func (n *Note) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {
	res, err := n.noteService.GetListNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
