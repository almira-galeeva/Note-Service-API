package note_v1

import (
	"context"
	"fmt"

	desc "github.com/almira-galeeva/testGrpc/pkg/note_v1"
)

func (n *Note) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	fmt.Println("Note Was Created.")
	fmt.Println("Title:", req.GetTitle())
	fmt.Println("Text:", req.GetText())
	fmt.Println("Author:", req.GetAuthor())
	fmt.Println()

	return &desc.CreateNoteResponse{
		Id: 1,
	}, nil
}
