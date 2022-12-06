package converter

import (
	"database/sql"

	"github.com/almira-galeeva/note-service-api/internal/model"
	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToNoteBody(noteBody *desc.NoteBody) *model.NoteBody {
	return &model.NoteBody{
		Title:  noteBody.GetTitle(),
		Text:   noteBody.GetText(),
		Author: noteBody.GetAuthor(),
		Email:  noteBody.GetEmail(),
	}
}

func ToUpdateNoteInfo(noteInfo *desc.UpdateNoteRequest) *model.UpdateNoteInfo {
	return &model.UpdateNoteInfo{
		Id:     noteInfo.Id,
		Title:  sql.NullString{String: noteInfo.GetNoteBody().GetTitle().GetValue(), Valid: true},
		Text:   sql.NullString{String: noteInfo.GetNoteBody().GetText().GetValue(), Valid: true},
		Author: sql.NullString{String: noteInfo.GetNoteBody().GetAuthor().GetValue(), Valid: true},
		Email:  sql.NullString{String: noteInfo.GetNoteBody().GetEmail().GetValue(), Valid: true},
	}
}

func ToDescNoteBody(noteBody *model.NoteBody) *desc.NoteBody {
	return &desc.NoteBody{
		Title:  noteBody.Title,
		Text:   noteBody.Text,
		Author: noteBody.Author,
		Email:  noteBody.Email,
	}
}

func ToDescWholeNote(noteBody *model.WholeNote) *desc.WholeNote {
	var updatedAt *timestamppb.Timestamp
	if noteBody.UpdatedAt.Valid {
		updatedAt = timestamppb.New(noteBody.UpdatedAt.Time)
	}

	return &desc.WholeNote{
		Id:        noteBody.Id,
		NoteBody:  ToDescNoteBody(noteBody.NoteBody),
		CreatedAt: timestamppb.New(noteBody.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToDescListWholeNote(noteBodies *[]model.WholeNote) []*desc.WholeNote {
	var notes []*desc.WholeNote

	for _, note := range *noteBodies {
		notes = append(notes, ToDescWholeNote(&note))
	}
	return notes
}
