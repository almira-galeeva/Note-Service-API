package note_v1

import (
	"context"
	"database/sql"
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
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGetList(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		ids       = []int64{gofakeit.Int64(), gofakeit.Int64()}
		title     = gofakeit.BeerName()
		text      = gofakeit.BeerStyle()
		author    = gofakeit.Name()
		email     = gofakeit.Email()
		createdAt = gofakeit.Date()

		repoErrText = gofakeit.Phrase()

		req = &desc.GetListNoteRequest{
			Ids: ids,
		}

		validRes = &desc.GetListNoteResponse{
			Results: []*desc.Note{
				{
					Id: ids[0],
					NoteBody: &desc.NoteBody{
						Title:  title,
						Text:   text,
						Author: author,
						Email:  email,
					},
					CreatedAt: timestamppb.New(createdAt),
					UpdatedAt: nil,
				},
				{
					Id: ids[1],
					NoteBody: &desc.NoteBody{
						Title:  title,
						Text:   text,
						Author: author,
						Email:  email,
					},
					CreatedAt: timestamppb.New(createdAt),
					UpdatedAt: nil,
				},
			},
		}

		repoRes = []*model.Note{
			{
				Id: ids[0],
				NoteBody: &model.NoteBody{
					Title:  title,
					Text:   text,
					Author: author,
					Email:  email,
				},
				CreatedAt: createdAt,
				UpdatedAt: sql.NullTime{},
			},
			{
				Id: ids[1],
				NoteBody: &model.NoteBody{
					Title:  title,
					Text:   text,
					Author: author,
					Email:  email,
				},
				CreatedAt: createdAt,
				UpdatedAt: sql.NullTime{},
			},
		}

		repoErr = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockNoteRepository(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().GetListNote(ctx, ids).Return(repoRes, nil),
		noteMock.EXPECT().GetListNote(ctx, ids).Return([]*model.Note{}, repoErr),
	)

	api := newMockNoteV1(Note{
		noteService: note.NewNoteMock(noteMock),
	})

	t.Run("success case", func(t *testing.T) {
		fmt.Println(req)
		res, err := api.GetListNote(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("note repo err", func(t *testing.T) {
		fmt.Println(req)
		_, err := api.GetListNote(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
