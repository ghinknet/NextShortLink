package repository

import (
	"NextShortLink/internal/model"
	"errors"

	"xorm.io/xorm"
)

type ConfigRepository struct {
	session *xorm.Session
}

func NewConfigRepository(session *xorm.Session) *ConfigRepository {
	return &ConfigRepository{session: session}
}

func (r *ConfigRepository) Update(config *model.DatabaseConfig) error {
	_, err := r.session.Where("key = ?", 0).Update(config)
	return err
}

func (r *ConfigRepository) Get() (*model.DatabaseConfig, error) {
	config := new(model.DatabaseConfig)
	has, err := r.session.Where("key = ?", 0).Get(config)
	if !has {
		return nil, model.ErrConfigError
	}
	return config, err
}

func (r *ConfigRepository) Init() error {
	_, err := r.Get()
	if err != nil {
		if !errors.Is(err, model.ErrConfigError) {
			return err
		} else {
			_, err = r.session.Insert(model.DatabaseConfig{
				Value: model.DefaultConfig,
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
