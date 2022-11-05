package note_v1

import (
	"context"
	"fmt"

	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	fmt.Printf("Note With Id %d Was Updated\n", req.GetId())
	fmt.Println("New Title:", req.GetTitle())
	fmt.Println("New Text:", req.GetText())
	fmt.Printf("New Author: %s\n\n", req.GetAuthor())

	return &desc.UpdateNoteResponse{
		Id:  req.GetId(),
		Res: 0,
	}, nil
}
