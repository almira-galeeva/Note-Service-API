package note_v1

import (
	"context"
	"errors"
	"fmt"
	"testing"

	noteMocks "github.com/almira-galeeva/note-service-api/internal/repository/mocks"
	"github.com/almira-galeeva/note-service-api/internal/service/note"
	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestDelete(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id          = gofakeit.Int64()
		repoErrText = gofakeit.Phrase()

		req = &desc.DeleteNoteRequest{
			Id: id,
		}

		validRes = &emptypb.Empty{}

		repoErr = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockNoteRepository(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().DeleteNote(ctx, id).Return(nil),
		noteMock.EXPECT().DeleteNote(ctx, id).Return(repoErr),
	)

	api := newMockNoteV1(Note{
		noteService: note.NewNoteMock(noteMock),
	})

	t.Run("success case", func(t *testing.T) {
		fmt.Println(req)
		res, err := api.DeleteNote(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("note repo err", func(t *testing.T) {
		fmt.Println(req)
		_, err := api.DeleteNote(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
