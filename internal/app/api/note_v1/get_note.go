package note_v1

import (
	"context"
	"fmt"

	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	fmt.Printf("Got Note With Id %d\n\n", req.GetId())

	return &desc.GetNoteResponse{
		Id:     req.GetId(),
		Title:  "Beautiful Note",
		Text:   "Hello World",
		Author: "Almira",
	}, nil
}
