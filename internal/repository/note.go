package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/almira-galeeva/note-service-api/internal/model"
	"github.com/almira-galeeva/note-service-api/internal/pkg/db"
	"github.com/almira-galeeva/note-service-api/internal/repository/table"
)

type NoteRepository interface {
	CreateNote(ctx context.Context, noteBody *model.NoteBody) (int64, error)
	GetNote(ctx context.Context, id int64) (*model.Note, error)
	GetListNote(ctx context.Context, ids []int64) ([]*model.Note, error)
	UpdateNote(ctx context.Context, UpdateNote *model.UpdateNoteInfo) (int64, error)
	DeleteNote(ctx context.Context, id int64) error
}

type repository struct {
	client db.Client
}

func NewNoteRepository(client db.Client) NoteRepository {
	return &repository{
		client: client,
	}
}

func (r *repository) CreateNote(ctx context.Context, noteBody *model.NoteBody) (int64, error) {
	builder := sq.Insert(table.Note).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author, email").
		Values(noteBody.Title, noteBody.Text, noteBody.Author, noteBody.Email).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "CreateNote",
		QueryRaw: query,
	}

	row, err := r.client.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) GetNote(ctx context.Context, id int64) (*model.Note, error) {
	builder := sq.Select("id, title, text, author, email, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.Note).
		Where(sq.Eq{"id": id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "GetNote",
		QueryRaw: query,
	}

	var note model.Note
	err = r.client.DB().GetContext(ctx, &note, q, args...)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (r *repository) GetListNote(ctx context.Context, ids []int64) ([]*model.Note, error) {
	builder := sq.Select("id, title, text, author, email, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.Note).
		Where(sq.Eq{"id": ids})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "GetListNote",
		QueryRaw: query,
	}

	var notes []*model.Note
	err = r.client.DB().SelectContext(ctx, &notes, q, args...)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *repository) UpdateNote(ctx context.Context, UpdateNote *model.UpdateNoteInfo) (int64, error) {
	builder := sq.Update(table.Note).
		PlaceholderFormat(sq.Dollar).
		Set("updated_at", time.Now())

	if UpdateNote.Title.Valid {
		builder = builder.Set("title", UpdateNote.Title.String)
	}
	if UpdateNote.Text.Valid {
		builder = builder.Set("text", UpdateNote.Text.String)
	}
	if UpdateNote.Author.Valid {
		builder = builder.Set("author", UpdateNote.Author.String)
	}
	if UpdateNote.Email.Valid {
		builder = builder.Set("email", UpdateNote.Email.String)
	}

	builder = builder.Where(sq.Eq{"id": UpdateNote.Id})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "UpdateNote",
		QueryRaw: query,
	}

	_, err = r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}

	return UpdateNote.Id, nil
}

func (r *repository) DeleteNote(ctx context.Context, id int64) error {
	builder := sq.Delete(table.Note).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "DeleteNote",
		QueryRaw: query,
	}

	_, err = r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
