package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/almira-galeeva/note-service-api/internal/repository/table"
	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetNote struct {
	Id          int64
	Title       string
	Text        string
	Author      string
	Email       string
	TsCreatedAt *timestamppb.Timestamp
	TsUpdatedAt *timestamppb.Timestamp
}

type NoteRepository interface {
	CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (int64, error)
	GetNote(ctx context.Context, req *desc.GetNoteRequest) (*GetNote, error)
	GetListNote(ctx context.Context, req *desc.GetListNoteRequest) ([]*GetNote, error)
	UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (int64, error)
	DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) error
}

type repository struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) NoteRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (int64, error) {
	builder := sq.Insert(table.Note).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author, email").
		Values(req.GetTitle(), req.GetText(), req.GetAuthor(), req.GetEmail()).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
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

func (r *repository) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*GetNote, error) {
	builder := sq.Select("id, title, text, author, email, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.Note).
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
		email     string
		createdAt time.Time
		updatedAt *time.Time
	}

	note := new(getNote)

	err = r.db.QueryRow(query, args...).Scan(&note.id, &note.title, &note.text, &note.author, &note.email, &note.createdAt, &note.updatedAt)

	tsCreatedAt := timestamppb.New(note.createdAt)

	tsUpdatedAt := new(timestamppb.Timestamp)
	if note.updatedAt != nil {
		tsUpdatedAt = timestamppb.New(*note.updatedAt)
	}

	if err != nil {
		return nil, err
	}

	return &GetNote{
		Id:          note.id,
		Title:       note.title,
		Text:        note.text,
		Author:      note.author,
		Email:       note.email,
		TsCreatedAt: tsCreatedAt,
		TsUpdatedAt: tsUpdatedAt,
	}, nil
}

func (r *repository) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) ([]*GetNote, error) {
	builder := sq.Select("id, title, text, author, email, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.Note).
		Where(sq.Eq{"id": req.GetIds()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var res []*GetNote
	for row.Next() {

		type getNote struct {
			id        int64
			title     string
			text      string
			author    string
			email     string
			createdAt time.Time
			updatedAt *time.Time
		}

		note := new(getNote)

		err = row.Scan(&note.id, &note.title, &note.text, &note.author, &note.email, &note.createdAt, &note.updatedAt)

		if err != nil {
			return nil, err
		}

		tsCreatedAt := timestamppb.New(note.createdAt)

		tsUpdatedAt := new(timestamppb.Timestamp)
		if note.updatedAt != nil {
			tsUpdatedAt = timestamppb.New(*note.updatedAt)
		}
		val := &GetNote{
			Id:          note.id,
			Title:       note.title,
			Text:        note.text,
			Author:      note.author,
			Email:       note.email,
			TsCreatedAt: tsCreatedAt,
			TsUpdatedAt: tsUpdatedAt,
		}

		res = append(res, val)
	}

	return res, nil
}

func (r *repository) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (int64, error) {
	builder := sq.Update(table.Note).
		PlaceholderFormat(sq.Dollar).
		SetMap(sq.Eq{"title": req.GetTitle(), "text": req.GetText(), "author": req.GetAuthor(), "updated_at": time.Now()}).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	return req.GetId(), nil
}

func (r *repository) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) error {
	builder := sq.Delete(table.Note).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer row.Close()

	return nil
}
