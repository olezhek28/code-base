package note

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/repository/mocks"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	noteRepository *mocks.NoteRepository

	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.noteRepository = mocks.NewNoteRepository(s.T())

	s.service = NewService(
		s.noteRepository,
	)
}

func (s *ServiceSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
