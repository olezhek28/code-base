package repository

import (
	"context"

	"github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/model"
)

type NoteRepository interface {
	Create(ctx context.Context, uuid string, info model.NoteInfo) error
	Get(ctx context.Context, uuid string) (model.Note, error)
}
