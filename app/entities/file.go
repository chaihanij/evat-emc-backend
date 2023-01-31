package entities

import "time"

type File struct {
	ID               *string
	UUID             string
	OriginalFileName string
	FileName         string
	FileExtension    string
	FilePath         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Files []File

type FileFilter struct {
	ID               *string
	UUID             *string
	OriginalFileName *string
	FileName         *string
	FilePath         *string
}
