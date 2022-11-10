package note_v1

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	builder := sq.Select("id, title, text, author, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(noteTable).
		Where(sq.Eq{"id": req.GetId()}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	type getNote struct {
		id        int64
		title     string
		text      string
		author    string
		createdAt time.Time
		updatedAt *time.Time
	}

	note := new(getNote)

	err = db.QueryRow(query, args...).Scan(&note.id, &note.title, &note.text, &note.author, &note.createdAt, &note.updatedAt)

	if err != nil {
		return nil, err
	}

	tsCreatedAt := timestamppb.New(note.createdAt)

	tsUpdatedAt := new(timestamppb.Timestamp)
	if note.updatedAt != nil {
		tsUpdatedAt = timestamppb.New(*note.updatedAt)
	}

	return &desc.GetNoteResponse{
		Id:        note.id,
		Title:     note.title,
		Text:      note.text,
		Author:    note.author,
		CreatedAt: tsCreatedAt,
		UpdatedAt: tsUpdatedAt,
	}, nil
}
