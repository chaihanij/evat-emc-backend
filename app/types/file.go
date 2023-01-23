package types

type File struct {
	FileStream  string `bson:"file_stream" json:"file_stream"`
	FileName    string `bson:"file_name" json:"file_name"`
	ContentType string `bson:"content_type"  json:"content_type"`
}

type Files []File
