package note_v1

import (
	"context"
	"fmt"

	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.Empty, error) {
	fmt.Printf("Delete Note With Id: %d\n", req.GetId())

	return &desc.Empty{}, nil
}
