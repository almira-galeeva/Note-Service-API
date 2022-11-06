package note_v1

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
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

	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `SELECT title, text, author FROM note WHERE id = $1`

	//res := []*desc.GetListNoteResponse_Result{}
	res := make([]*desc.GetListNoteResponse_Result, 0)

	for i := 0; i < len(req.GetIds()); i++ {
		row, err := db.QueryContext(ctx, query, req.GetIds()[i])
		if err != nil {
			return nil, err
		}
		defer row.Close()

		row.Next()
		var title, text, author string
		err = row.Scan(&title, &text, &author)
		if err != nil {
			return nil, err
		}

		val := &desc.GetListNoteResponse_Result{
			Id:     req.GetIds()[i],
			Title:  title,
			Text:   text,
			Author: author,
		}
		res = append(res, val)
	}

	return &desc.GetListNoteResponse{
		Results: res,
	}, nil
}
