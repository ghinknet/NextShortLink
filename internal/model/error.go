package model

import "errors"

var ErrConfigError = errors.New("config error")
var ErrApplicationNotFound = errors.New("application not found")
var ErrPermissionDenied = errors.New("permission denied")
var ErrNoPackageAvailable = errors.New("no available package")
var ErrLinkNotExist = errors.New("link does not exist")
