package note_v1

import (
	"context"
	"fmt"
	desc "github.com/almira-galeeva/testGrpc/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	fmt.Printf("Note With Id %d Was Updated", req.GetId())
	fmt.Println()
	fmt.Println("New Title:", req.GetTitle())
	fmt.Println("New Text:", req.GetText())
	fmt.Println("New Author:", req.GetAuthor())
	fmt.Println()

	return &desc.UpdateNoteResponse{
		Id:  req.GetId(),
		Res: 0,
	}, nil
}
