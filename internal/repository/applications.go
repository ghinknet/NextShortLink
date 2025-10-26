package repository

import (
	"NextShortLink/internal/model"

	"xorm.io/xorm"
)

type ApplicationRepository struct {
	session *xorm.Session
}

func NewApplicationRepository(session *xorm.Session) *ApplicationRepository {
	return &ApplicationRepository{session: session}
}

func (r *ApplicationRepository) Get(secretID string, secretKey string) (id int64, err error) {
	application := &model.DatabaseApplication{}
	has, err := r.session.Where("secret_id = ?", secretID).
		Where("secret_key = ?", secretKey).
		Get(application)
	if err != nil {
		return 0, err
	}
	if !has {
		return 0, model.ErrApplicationNotFound
	}

	return application.ID, nil
}
