package note_v1

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/almira-galeeva/note-service-api/internal/model"
	noteMocks "github.com/almira-galeeva/note-service-api/internal/repository/mocks"
	"github.com/almira-galeeva/note-service-api/internal/service/note"
	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id          = gofakeit.Int64()
		title       = gofakeit.BeerName()
		text        = gofakeit.BeerStyle()
		author      = gofakeit.Name()
		email       = gofakeit.Email()
		repoErrText = gofakeit.Phrase()

		req = &desc.CreateNoteRequest{
			NoteBody: &desc.NoteBody{
				Title:  title,
				Text:   text,
				Author: author,
				Email:  email,
			},
		}

		repoReq = &model.NoteBody{
			Title:  title,
			Text:   text,
			Author: author,
			Email:  email,
		}

		validRes = &desc.CreateNoteResponse{
			Id: id,
		}

		repoErr = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockNoteRepository(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().CreateNote(ctx, repoReq).Return(id, nil),
		noteMock.EXPECT().CreateNote(ctx, repoReq).Return(int64(0), repoErr),
	)

	api := newMockNoteV1(Note{
		noteService: note.NewNoteMock(noteMock),
	})

	t.Run("success case", func(t *testing.T) {
		fmt.Println(req.GetNoteBody())
		res, err := api.CreateNote(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("note repo err", func(t *testing.T) {
		fmt.Println(req.GetNoteBody())
		_, err := api.CreateNote(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
