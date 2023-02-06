package models

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewFiles(input []entities.File) Files {
	var files Files
	now := time.Now()
	for _, value := range input {
		file := &File{
			ID:               primitive.NewObjectID(),
			UUID:             uuid.NewString(),
			OriginalFileName: value.OriginalFileName,
			FileName:         value.FileName,
			FileExtension:    value.FileExtension,
			FilePath:         value.FilePath,
			FileFullPath:     value.FileFullPath,
			CreatedAt:        now,
			UpdatedAt:        now,
		}
		files = append(files, *file)
	}
	return files
}
