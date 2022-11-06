package note_v1

import (
	"context"
	"fmt"

	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	fmt.Printf("Got Note With Id %d\n\n", req.GetId())

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

	row, err := db.QueryContext(ctx, query, req.GetId())
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

	return &desc.GetNoteResponse{
		Id:     req.GetId(),
		Title:  title,
		Text:   text,
		Author: author,
	}, nil
}
