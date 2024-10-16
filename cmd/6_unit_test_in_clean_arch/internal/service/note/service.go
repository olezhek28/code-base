package note

import (
	"github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/repository"
)

type service struct {
	noteRepository repository.NoteRepository
}

func NewService(
	noteRepository repository.NoteRepository,
) *service {
	return &service{
		noteRepository: noteRepository,
	}
}
