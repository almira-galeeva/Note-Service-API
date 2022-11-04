package note_v1

import (
	"context"
	"fmt"

	desc "github.com/almira-galeeva/testGrpc/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.DeleteNodeResponse, error) {
	fmt.Println("Delete Note With Id:", req.GetId())
	fmt.Println()

	return &desc.DeleteNodeResponse{
		Id:  req.GetId(),
		Res: 0,
	}, nil
}
