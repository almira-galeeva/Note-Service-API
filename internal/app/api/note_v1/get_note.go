package note_v1

import (
	"context"
	"fmt"
	desc "github.com/almira-galeeva/testGrpc/pkg/note_v1"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	fmt.Println("Got Note With Id:", req.GetId())
	fmt.Println()

	return &desc.GetNoteResponse{
		Title:  "Beautiful Note",
		Text:   "Hello World",
		Author: "Almira",
	}, nil
}
