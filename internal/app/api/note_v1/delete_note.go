package note_v1

import (
	"context"

	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*emptypb.Empty, error) {
	err := n.noteService.DeleteNote(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
