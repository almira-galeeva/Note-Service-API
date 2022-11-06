package note_v1

import (
	"context"
	"fmt"

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

	query := `UPDATE note SET title = $1, text = $2, author = $3 WHERE id = $4`

	row, err := db.QueryContext(ctx, query, req.GetTitle(), req.GetText(), req.GetAuthor(), req.GetId())
	if err != nil {
		return nil, err
	}
	defer row.Close()

	return &desc.UpdateNoteResponse{
		Id:  req.GetId(),
		Res: 0,
	}, nil
}
