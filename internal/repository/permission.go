package repository

import (
	"NextShortLink/internal/model"

	"xorm.io/xorm"
)

type PermissionRepository struct {
	session *xorm.Session
}

func NewPermissionRepository(session *xorm.Session) *PermissionRepository {
	return &PermissionRepository{session: session}
}

func (r *PermissionRepository) Check(appID int64, interfaceName string) (
	disableKey bool, disableToken bool,
	qps int64, qpm int64,
	blacklist []string, whitelist []string,
	err error,
) {
	permission := model.DatabasePermission{}
	has, err := r.session.Where("application = ?", appID).
		Where("interface = ?", interfaceName).Get(&permission)
	if err != nil {
		return true, true, 0, 0, []string{}, []string{}, err
	}
	if !has {
		return true, true, 0, 0, []string{}, []string{}, model.ErrPermissionDenied
	}

	return permission.DisableKey, permission.DisableToken,
		permission.QPS, permission.QPM,
		permission.Blacklist, permission.Whitelist,
		nil
}
