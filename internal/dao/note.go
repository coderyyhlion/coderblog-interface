// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"coderblog-interface/internal/dao/internal"
)

// internalNoteDao is internal type for wrapping internal DAO implements.
type internalNoteDao = *internal.NoteDao

// noteDao is the data access object for table note.
// You can define custom methods on it to extend its functionality as you wish.
type noteDao struct {
	internalNoteDao
}

var (
	// Note is globally public accessible object for table note operations.
	Note = noteDao{
		internal.NewNoteDao(),
	}
)

// Fill with you ideas below.
