package note_v1

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
)

func (n *Note) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {
	var idsString strings.Builder
	for i := 0; i < len(req.GetIds()); i++ {
		idsString.WriteString(strconv.FormatInt(req.GetIds()[i], 10))
		if i != len(req.GetIds())-1 {
			idsString.WriteString(", ")
		}
	}
	fmt.Printf("Got Notes With Ids: %s\n\n", idsString.String())

	one := &desc.GetListNoteResponse_Result{
		Id:     req.GetIds()[0],
		Title:  "Note 1",
		Text:   "Such a lonely day",
		Author: "SOAD",
	}
	two := &desc.GetListNoteResponse_Result{
		Id:     req.GetIds()[1],
		Title:  "Note 2",
		Text:   "Somewhere in the end of all this hate",
		Author: "TPR",
	}

	return &desc.GetListNoteResponse{
		Results: []*desc.GetListNoteResponse_Result{one, two},
	}, nil
}
