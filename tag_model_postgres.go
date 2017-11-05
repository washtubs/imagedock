package imagedock

import (
	postgres "github.com/go-pg/pg"
)

type PostgresTagModel struct {
}

func (p *PostgresTagModel) CreateTag(tagId TagId, term Term, imageId ContentId) (TagId, TagError) {
	panic("not implemented")
	postgres.Connect(&postgres.Options{
		Network: "tcp",
	})
	return TagId(1), nil
}

func (p *PostgresTagModel) AssignCanonTag(tagId TagId, imageId ContentId) {
	panic("not implemented")
}

func (p *PostgresTagModel) UnassignCanonTag(tagId TagId, imageId ContentId) {
	panic("not implemented")
}

func (p *PostgresTagModel) AssignPrivateTag(tagId TagId, userId UserId, imageId ContentId) {
	panic("not implemented")
}

func (p *PostgresTagModel) UnassignPrivateTag(tagId TagId, userId UserId, imageId ContentId) {
	panic("not implemented")
}
