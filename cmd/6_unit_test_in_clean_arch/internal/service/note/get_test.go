package note

import (
	"github.com/brianvoe/gofakeit/v6"

	"github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/model"
)

func (s *ServiceSuite) TestGetSuccess() {
	var (
		noteUUID  = gofakeit.UUID()
		title     = gofakeit.BookTitle()
		content   = gofakeit.Paragraph(3, 5, 5, " ")
		createdAt = gofakeit.Date()

		note = model.Note{
			UUID: noteUUID,
			Info: model.NoteInfo{
				Title:   title,
				Content: content,
			},
			CreatedAt: createdAt,
		}
	)

	s.noteRepository.On("Get", s.ctx, noteUUID).Return(note, nil)

	res, err := s.service.Get(s.ctx, noteUUID)
	s.NoError(err)
	s.Equal(note, res)
}

func (s *ServiceSuite) TestGetRepoError() {
	var (
		repoErr  = gofakeit.Error()
		noteUUID = gofakeit.UUID()
	)

	s.noteRepository.On("Get", s.ctx, noteUUID).Return(model.Note{}, repoErr)

	res, err := s.service.Get(s.ctx, noteUUID)
	s.Error(err)
	s.ErrorIs(err, repoErr)
	s.Empty(res)
}
