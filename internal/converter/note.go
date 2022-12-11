package converter

import (
	"database/sql"

	"github.com/almira-galeeva/note-service-api/internal/model"
	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToNote(noteBody *desc.NoteBody) *model.NoteBody {
	return &model.NoteBody{
		Title:  noteBody.GetTitle(),
		Text:   noteBody.GetText(),
		Author: noteBody.GetAuthor(),
		Email:  noteBody.GetEmail(),
	}
}

func ToUpdateNoteInfo(noteInfo *desc.UpdateNoteRequest) *model.UpdateNoteInfo {
	var validTitle, validText, validAuthor, validEmail bool

	if noteInfo.GetNoteBody().GetTitle() != nil {
		validTitle = true
	}

	if noteInfo.GetNoteBody().GetText() != nil {
		validText = true
	}

	if noteInfo.GetNoteBody().GetAuthor() != nil {
		validAuthor = true
	}

	if noteInfo.GetNoteBody().GetEmail() != nil {
		validEmail = true
	}

	return &model.UpdateNoteInfo{
		Id:     noteInfo.GetId(),
		Title:  sql.NullString{String: noteInfo.GetNoteBody().GetTitle().GetValue(), Valid: validTitle},
		Text:   sql.NullString{String: noteInfo.GetNoteBody().GetText().GetValue(), Valid: validText},
		Author: sql.NullString{String: noteInfo.GetNoteBody().GetAuthor().GetValue(), Valid: validAuthor},
		Email:  sql.NullString{String: noteInfo.GetNoteBody().GetEmail().GetValue(), Valid: validEmail},
	}
}

func ToDescNote(noteBody *model.NoteBody) *desc.NoteBody {
	if noteBody == nil {
		return nil
	}

	return &desc.NoteBody{
		Title:  noteBody.Title,
		Text:   noteBody.Text,
		Author: noteBody.Author,
		Email:  noteBody.Email,
	}
}

func ToDescWholeNote(noteBody *model.Note) *desc.Note {
	if noteBody == nil {
		return nil
	}

	var updatedAt *timestamppb.Timestamp
	if noteBody.UpdatedAt.Valid {
		updatedAt = timestamppb.New(noteBody.UpdatedAt.Time)
	}

	return &desc.Note{
		Id:        noteBody.Id,
		NoteBody:  ToDescNote(noteBody.NoteBody),
		CreatedAt: timestamppb.New(noteBody.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToDescListWholeNote(noteBodies []*model.Note) []*desc.Note {
	notes := make([]*desc.Note, len(noteBodies))

	for ind, note := range noteBodies {
		notes[ind] = ToDescWholeNote(note)
	}
	return notes
}
