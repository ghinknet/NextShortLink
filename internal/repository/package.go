package repository

import (
	"NextShortLink/internal/infra/database"
	"NextShortLink/internal/model"
	"time"

	"xorm.io/xorm"
)

type PackageRepository struct {
	session *xorm.Session
}

func NewPackageRepository(session *xorm.Session) *PackageRepository {
	return &PackageRepository{session: session}
}

func (r *PackageRepository) Take(appID int64, interfaceName string) error {
	if err := r.session.Begin(); err != nil {
		return err
	}

	var packages []model.DatabasePackage
	if err := r.session.
		Where("application = ?", appID).
		And("interface = ?", interfaceName).
		Asc("priority").
		Find(&packages); err != nil {
		return database.RollbackError(r.session, err)
	}

	for _, pkg := range packages {
		if pkg.Unlimit && pkg.AvailableFrom < time.Now().Unix() && pkg.AvailableTo > time.Now().Unix() {
			return database.RollbackError(r.session, nil)
		}
	}

	for _, pkg := range packages {
		if pkg.Used < pkg.Total && pkg.AvailableFrom < time.Now().Unix() && pkg.AvailableTo > time.Now().Unix() {
			if _, err := r.session.
				Where("id = ?", pkg.ID).
				Incr("used", 1).
				Update(new(model.DatabasePackage)); err != nil {
				return database.RollbackError(r.session, err)
			}

			return r.session.Commit()
		}
	}

	if err := r.session.Commit(); err != nil {
		return err
	}
	return model.ErrNoPackageAvailable
}
