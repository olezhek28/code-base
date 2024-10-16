package note

import (
	"context"

	"github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/model"
)

func (s *service) Get(ctx context.Context, uuid string) (model.Note, error) {
	note, err := s.noteRepository.Get(ctx, uuid)
	if err != nil {
		return model.Note{}, err
	}

	return note, nil
}
