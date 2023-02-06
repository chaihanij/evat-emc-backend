package entities

import "time"

type Album struct {
	ID        string
	UUID      string
	Title     string
	Images    interface{}
	Year      string
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string
	UpdatedBy string
}

type AlbumPartialUpdate struct {
	ID        *string
	UUID      *string
	Title     *string
	Images    interface{}
	Year      *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	CreatedBy *string
	UpdatedBy *string
}

type AlbumFilter struct {
	ID       *string
	UUID     *string
	Year     *string
	Page     *int64
	PageSize *int64
}
