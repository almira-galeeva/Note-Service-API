package note_v1

import (
	"github.com/almira-galeeva/note-service-api/internal/service/note"
	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
)

type Note struct {
	desc.UnimplementedNoteV1Server

	noteService *note.Service
}

func NewNote(noteService *note.Service) *Note {
	return &Note{
		noteService: noteService,
	}
}

func newMockNoteV1(n Note) *Note {
	return &Note{
		desc.UnimplementedNoteV1Server{},

		n.noteService,
	}
}
