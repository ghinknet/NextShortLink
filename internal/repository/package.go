package repository

import (
	"NextShortLink/internal/database"
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
	err := r.session.
		Where("application = ?", appID).
		And("interface = ?", interfaceName).
		Asc("priority").
		Find(&packages)
	if err != nil {
		database.Rollback(r.session)
		return err
	}

	for _, pkg := range packages {
		if pkg.Unlimit && pkg.AvailableFrom < time.Now().Unix() && pkg.AvailableTo > time.Now().Unix() {
			database.Rollback(r.session)
			return nil
		}
	}

	for _, pkg := range packages {
		if pkg.Used < pkg.Total && pkg.AvailableFrom < time.Now().Unix() && pkg.AvailableTo > time.Now().Unix() {
			_, err = r.session.
				Table("package").
				Where("id = ?", pkg.ID).
				Update(map[string]any{
					"used": pkg.Used + 1,
				})
			if err != nil {
				database.Rollback(r.session)
				return err
			} else {
				if err = r.session.Commit(); err != nil {
					return err
				}
				return nil
			}
		}
	}

	if err = r.session.Commit(); err != nil {
		return err
	}
	return model.ErrNoPackageAvailable
}
