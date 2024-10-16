package note

import (
	"context"

	"github.com/google/uuid"

	"github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/model"
)

func (s *service) Create(ctx context.Context, info model.NoteInfo) (string, error) {
	noteUUID := uuid.New().String()

	err := s.noteRepository.Create(ctx, noteUUID, info)
	if err != nil {
		return "", err
	}

	return noteUUID, nil
}
