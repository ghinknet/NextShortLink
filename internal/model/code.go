package model

import "net/http"

const CodeOK = http.StatusOK

const CodeFound = http.StatusFound

const CodeBadRequest = http.StatusBadRequest
const CodeUnauthorized = http.StatusUnauthorized
const CodeForbidden = http.StatusForbidden
const CodeNotFound = http.StatusNotFound
const CodeMethodNotAllowed = http.StatusMethodNotAllowed
const CodeTeaPot = http.StatusTeapot
const CodeTooManyRequests = http.StatusTooManyRequests

const CodeInternalServerError = http.StatusInternalServerError

const CodeMissingParameter = 800
const CodePermissionDenied = 801
const CodeNoPackageAvailable = 802
const CodeApplicationNotFound = 803
const CodeLinkNotExist = 804
const CodeLinkInvalid = 805
const CodeValidityInvalid = 806
