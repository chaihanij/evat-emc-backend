package models

import (
	"time"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	ID               primitive.ObjectID `bson:"_id"`
	UUID             string             `bson:"uuid"`
	OriginalFileName string             `bson:"original_filename"`
	FileName         string             `bson:"filename"`
	FileExtension    string             `bson:"file_extension"`
	FilePath         string             `bson:"file_path"`
	FileFullPath     string             `bson:"file_full_path"`
	CreatedAt        time.Time          `bson:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at"`
}

func NewFile(originalFilename, fileName, fileExtension, filePath, fileFullPath string) *File {
	return &File{
		ID:               primitive.NewObjectID(),
		UUID:             uuid.NewString(),
		OriginalFileName: originalFilename,
		FileName:         fileName,
		FileExtension:    fileExtension,
		FilePath:         filePath,
		FileFullPath:     fileFullPath,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

func (file *File) ToEntity() *entities.File {
	id := file.ID.String()
	return &entities.File{
		ID:               pointer.ToString(id),
		UUID:             file.UUID,
		OriginalFileName: file.OriginalFileName,
		FileName:         file.FileName,
		FileExtension:    file.FileExtension,
		FilePath:         file.FilePath,
		FileFullPath:     file.FileFullPath,
		CreatedAt:        file.CreatedAt,
		UpdatedAt:        file.UpdatedAt,
	}
}

type Files []File

func (fs Files) ToEntity() []entities.File {
	var files []entities.File
	for _, v := range fs {
		file := v.ToEntity()
		files = append(files, *file)
	}
	return files
}
