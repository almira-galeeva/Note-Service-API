package note

import "github.com/almira-galeeva/note-service-api/internal/repository"

type Service struct {
	noteRepository repository.NoteRepository
}

func NewService(noteRepository repository.NoteRepository) *Service {
	return &Service{
		noteRepository: noteRepository,
	}
}

func NewNoteMock(deps ...interface{}) *Service {
	is := Service{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.NoteRepository:
			is.noteRepository = s
		}
	}
	return &is
}
