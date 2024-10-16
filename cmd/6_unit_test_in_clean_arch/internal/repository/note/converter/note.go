package converter

import (
	"github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/model"
	modelRepo "github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/repository/note/model"
)

func ToNoteInfoFromService(info model.NoteInfo) modelRepo.NoteInfo {
	return modelRepo.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}

func ToNoteFromRepo(note modelRepo.Note) model.Note {
	return model.Note{
		UUID:      note.UUID,
		Info:      ToNoteInfoFromRepo(note.Info),
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
}

func ToNoteInfoFromRepo(info modelRepo.NoteInfo) model.NoteInfo {
	return model.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}
