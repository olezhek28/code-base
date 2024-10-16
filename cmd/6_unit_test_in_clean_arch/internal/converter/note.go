package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/model"
	desc "github.com/olezhek28/code-base/6_unit_test_in_clean_arch/pkg/note_v1"
)

func ToNoteFromService(note model.Note) *desc.Note {
	var updatedAt *timestamppb.Timestamp
	if note.UpdatedAt.Valid {
		updatedAt = timestamppb.New(note.UpdatedAt.Time)
	}

	return &desc.Note{
		Uuid:      note.UUID,
		Info:      ToNoteInfoFromService(note.Info),
		CreatedAt: timestamppb.New(note.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToNoteInfoFromService(info model.NoteInfo) *desc.NoteInfo {
	return &desc.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}

func ToNoteInfoFromDesc(info *desc.NoteInfo) model.NoteInfo {
	return model.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}
