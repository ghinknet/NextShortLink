package repository

import (
	"NextShortLink/internal/model"
	"time"

	"xorm.io/xorm"
)

type LinkRepository struct {
	session *xorm.Session
}

func NewLinkRepository(session *xorm.Session) *LinkRepository {
	return &LinkRepository{session: session}
}

func (r *LinkRepository) Read(linkID int64) (string, *int64, error) {
	// Offset
	linkID -= 4000

	link := new(model.DatabaseLink)
	has, err := r.session.ID(linkID).Get(link)
	if err != nil {
		return "", nil, err
	}
	if !has {
		return "", nil, model.ErrLinkNotExist
	}

	if link.Validity != nil && *link.Validity < time.Now().Unix() {
		_, err = r.session.ID(linkID).Delete(link)
		if err != nil {
			return "", nil, err
		}
		return "", nil, model.ErrLinkNotExist
	}

	return link.Link, link.Validity, nil
}

func (r *LinkRepository) Insert(link *model.DatabaseLink) error {
	_, err := r.session.Insert(link)
	if err != nil {
		return err
	}

	// Offset
	link.ID += 4000
	
	return nil
}
