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
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUpdate(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id     = gofakeit.Int64()
		title  = gofakeit.BeerName()
		text   = gofakeit.BeerStyle()
		author = gofakeit.Name()
		email  = gofakeit.Email()

		repoErrText = gofakeit.Phrase()

		req = &desc.UpdateNoteRequest{
			Id: id,
			NoteBody: &desc.UpdateNoteInfo{
				Title: &wrapperspb.StringValue{
					Value: title,
				},
				Text: &wrapperspb.StringValue{
					Value: text,
				},
				Author: &wrapperspb.StringValue{
					Value: author,
				},
				Email: &wrapperspb.StringValue{
					Value: email,
				},
			},
		}

		repoReq = &model.UpdateNoteInfo{
			Id: id,
			Title: sql.NullString{
				String: title,
				Valid:  true,
			},
			Text: sql.NullString{
				String: text,
				Valid:  true,
			},
			Author: sql.NullString{
				String: author,
				Valid:  true,
			},
			Email: sql.NullString{
				String: email,
				Valid:  true,
			},
		}

		validRes = &desc.UpdateNoteResponse{
			Id: id,
		}

		repoErr = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockNoteRepository(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().UpdateNote(ctx, repoReq).Return(id, nil),
		noteMock.EXPECT().UpdateNote(ctx, repoReq).Return(int64(0), repoErr),
	)

	api := newMockNoteV1(Note{
		noteService: note.NewNoteMock(noteMock),
	})

	t.Run("success case", func(t *testing.T) {
		fmt.Println(req)
		res, err := api.UpdateNote(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("note repo err", func(t *testing.T) {
		fmt.Println(req)
		_, err := api.UpdateNote(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
