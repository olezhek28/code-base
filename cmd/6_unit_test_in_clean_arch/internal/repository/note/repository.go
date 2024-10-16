package note

import (
	"context"
	"sync"
	"time"

	"github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/model"
	"github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/repository"
	"github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/repository/note/converter"
	modelRepo "github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/repository/note/model"
)

type repo struct {
	data map[string]modelRepo.Note
	m    sync.RWMutex
}

func NewRepository() repository.NoteRepository {
	return &repo{
		data: make(map[string]modelRepo.Note),
	}
}

func (r *repo) Create(_ context.Context, uuid string, info model.NoteInfo) error {
	r.m.Lock()
	defer r.m.Unlock()

	r.data[uuid] = modelRepo.Note{
		UUID:      uuid,
		Info:      converter.ToNoteInfoFromService(info),
		CreatedAt: time.Now(),
	}

	return nil
}

func (r *repo) Get(_ context.Context, uuid string) (model.Note, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	note, ok := r.data[uuid]
	if !ok {
		return model.Note{}, model.ErrorNoteNotFound
	}

	return converter.ToNoteFromRepo(note), nil
}
