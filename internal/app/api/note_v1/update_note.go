package note_v1

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	fmt.Printf("Note With Id %d Was Updated\n", req.GetId())
	fmt.Println("New Title:", req.GetTitle())
	fmt.Println("New Text:", req.GetText())
	fmt.Printf("New Author: %s\n\n", req.GetAuthor())

	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Update(noteTable).
		PlaceholderFormat(sq.Dollar).
		SetMap(sq.Eq{"title": req.GetTitle(), "text": req.GetText(), "author": req.GetAuthor(), "updated_at": time.Now()}).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	return &desc.UpdateNoteResponse{
		Id:  req.GetId(),
		Res: 0,
	}, nil
}
