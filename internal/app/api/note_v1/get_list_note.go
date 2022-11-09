package note_v1

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	builder := sq.Select("id, title, text, author, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(noteTable).
		Where(sq.Eq{"id": req.GetIds()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	res := make([]*desc.GetListNoteResponse_Result, 0)
	for row.Next() {

		type getNote struct {
			id        int64
			title     string
			text      string
			author    string
			createdAt time.Time
			updatedAt *time.Time
		}

		note := new(getNote)

		row.Scan(&note.id, &note.title, &note.text, &note.author, &note.createdAt, &note.updatedAt)

		if err != nil {
			return nil, err
		}

		val := new(desc.GetListNoteResponse_Result)

		tsCreatedAt := timestamppb.New(note.createdAt)

		tsUpdatedAt := new(timestamppb.Timestamp)
		if note.updatedAt != nil {
			tsUpdatedAt = timestamppb.New(*note.updatedAt)
		}
		val = &desc.GetListNoteResponse_Result{
			Id:        note.id,
			Title:     note.title,
			Text:      note.text,
			Author:    note.author,
			CreatedAt: tsCreatedAt,
			UpdatedAt: tsUpdatedAt,
		}

		res = append(res, val)
	}

	return &desc.GetListNoteResponse{
		Results: res,
	}, nil
}
