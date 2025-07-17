package water

type StatusCode int

const (
	StatusCodeOK        StatusCode = 200
	StatusCodeCreated   StatusCode = 201
	StatusCodeNoContent StatusCode = 204

	StatusCodeBadRequest   StatusCode = 400
	StatusCodeUnauthorized StatusCode = 401
	StatusCodeForbidden    StatusCode = 403
	StatusCodeNotFound     StatusCode = 404
	StatusCodeMethodNot    StatusCode = 405
	StatusCodeConflict     StatusCode = 409

	StatusCodeInternalServerError StatusCode = 500
)
