package note_v1

import (
	"context"
	"fmt"
	"strconv"

	desc "github.com/almira-galeeva/testGrpc/pkg/note_v1"
)

func (n *Note) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {

	idsString := ""
	for i := 0; i < len(req.GetIds()); i++ {
		idsString += strconv.FormatInt(req.GetIds()[i], 10)
		if i != len(req.GetIds())-1 {
			idsString += ", "
		}
	}
	fmt.Println("Got Notes With Ids:", idsString)
	fmt.Println()

	one := &desc.GetListNoteResponse_Result{
		Title:  "Note 1",
		Text:   "Such a lonely day",
		Author: "SOAD",
	}
	two := &desc.GetListNoteResponse_Result{
		Title:  "Note 2",
		Text:   "Somewhere in the end of all this hate",
		Author: "TPR",
	}

	return &desc.GetListNoteResponse{
		Results: []*desc.GetListNoteResponse_Result{one, two},
	}, nil
}
