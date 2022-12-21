package note_v1

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/almira-galeeva/note-service-api/internal/model"
	noteMocks "github.com/almira-galeeva/note-service-api/internal/repository/mocks"
	"github.com/almira-galeeva/note-service-api/internal/service/note"
	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGet(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id        = gofakeit.Int64()
		title     = gofakeit.BeerName()
		text      = gofakeit.BeerStyle()
		author    = gofakeit.Name()
		email     = gofakeit.Email()
		createdAt = gofakeit.Date()

		repoErrText = gofakeit.Phrase()

		req = &desc.GetNoteRequest{
			Id: id,
		}

		validRes = &desc.GetNoteResponse{
			WholeNote: &desc.Note{
				Id: id,
				NoteBody: &desc.NoteBody{
					Title:  title,
					Text:   text,
					Author: author,
					Email:  email,
				},
				CreatedAt: timestamppb.New(createdAt),
				UpdatedAt: nil,
			},
		}

		repoRes = &model.Note{
			Id: id,
			NoteBody: &model.NoteBody{
				Title:  title,
				Text:   text,
				Author: author,
				Email:  email,
			},
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{},
		}

		repoErr = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockNoteRepository(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().GetNote(ctx, id).Return(repoRes, nil),
		noteMock.EXPECT().GetNote(ctx, id).Return(nil, repoErr),
	)

	api := newMockNoteV1(Note{
		noteService: note.NewNoteMock(noteMock),
	})

	t.Run("success case", func(t *testing.T) {
		res, err := api.GetNote(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("note repo err", func(t *testing.T) {
		_, err := api.GetNote(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
