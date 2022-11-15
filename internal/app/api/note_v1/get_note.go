package note_v1

import (
	"context"
	//"fmt"
	//"time"

	//sq "github.com/Masterminds/squirrel"
	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
	//"github.com/jmoiron/sqlx"
	//"google.golang.org/protobuf/types/known/timestamppb"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	res, err := n.noteService.GetNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
