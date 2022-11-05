package note_v1

import (
	"context"
	"fmt"

	desc "github.com/almira-galeeva/Note-Service-API/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.DeleteNodeResponse, error) {
	fmt.Printf("Delete Note With Id: %d\n", req.GetId())

	return &desc.DeleteNodeResponse{}, nil
}
