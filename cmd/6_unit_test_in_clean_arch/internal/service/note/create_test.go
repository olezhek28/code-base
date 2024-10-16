package note

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/model"
)

func (s *ServiceSuite) TestCreateSuccess() {
	var (
		title   = gofakeit.BookTitle()
		content = gofakeit.Paragraph(3, 5, 5, " ")

		noteInfo = model.NoteInfo{
			Title:   title,
			Content: content,
		}
	)

	s.noteRepository.On("Create", s.ctx, mock.Anything, noteInfo).Return(nil)

	noteUUID, err := s.service.Create(s.ctx, noteInfo)
	s.Require().NoError(err)

	_, err = uuid.Parse(noteUUID)
	s.Require().NoError(err)
}

func (s *ServiceSuite) TestCreateRepoError() {
	var (
		repoErr = gofakeit.Error()
		title   = gofakeit.BookTitle()
		content = gofakeit.Paragraph(3, 5, 5, " ")

		noteInfo = model.NoteInfo{
			Title:   title,
			Content: content,
		}
	)

	s.noteRepository.On("Create", s.ctx, mock.Anything, noteInfo).Return(repoErr)

	noteUUID, err := s.service.Create(s.ctx, noteInfo)
	s.Require().Error(err)
	s.Require().ErrorIs(err, repoErr)
	s.Require().Empty(noteUUID)
}
