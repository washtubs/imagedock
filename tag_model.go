package imagedock

import "errors"

type ContentId int64
type TagId int64
type UserId int64

type Term string

type TagError error

//type TagError struct {
//tagId TagId
//}

var (
	ErrTagExists TagError = errors.New("Tag already exists")
)

type TagModel interface {

	// Attempts to create a tag for the term, returns ErrTagExists if it already exists
	CreateTag(term Term) (TagId, TagError)

	AssignCanonTag(tagId TagId, imageId ContentId)
	UnassignCanonTag(tagId TagId, imageId ContentId)
	AssignPrivateTag(tagId TagId, userId UserId, imageId ContentId)
	UnassignPrivateTag(tagId TagId, userId UserId, imageId ContentId)
}
