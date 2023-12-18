package errors

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	internali18n = multilang{
		Uz: "Serverda xatolik",
		Ru: "Ошибка сервера",
		En: "Internal server error",
	}

	badRequesti18n = multilang{
		Uz: "Xato so'rov",
		Ru: "Неверная ошибка запроса",
		En: "Bad request error",
	}

	notFoundi18n = multilang{
		Uz: "Ma'lumot topilmadi",
		Ru: "Ошибка не найдена",
		En: "Not found error",
	}

	forbiddeni18n = multilang{
		Uz: "Taqiqlangan",
		Ru: "Запрещенный",
		En: "Forbidden",
	}

	unauthorizedi18n = multilang{
		Uz: "No'malum foydalanuvchi",
		Ru: "Неизвестный пользователь",
		En: "Unauthorized",
	}
)

type multilang struct {
	Uz string
	Ru string
	En string
}

type Error struct {
	i18nmsg multilang
	msg     string
	code    int
}

func (e Error) Error() string {
	return e.msg
}

func (e Error) Code() int {
	return e.code
}

func (e Error) ErrorI18nMsg(lang string) string {
	switch lang {
	case "uz":
		return e.i18nmsg.Uz
	case "ru":
		return e.i18nmsg.Ru
	case "en":
		return e.i18nmsg.En
	default:
		return e.i18nmsg.Uz
	}
}

func new(msg string, code int, i18n multilang) error {
	return Error{
		msg:     msg,
		code:    code,
		i18nmsg: i18n,
	}
}

func Wrap(msg string, err error) error {
	myErr, ok := err.(Error)

	if !ok {
		return Error{
			msg: fmt.Sprintf("%s: %s", msg, err.Error()),
		}
	}

	myErr.msg = fmt.Sprintf("%s: %s", msg, myErr.msg)
	return myErr
}

func IsWrappedWith(err, target error) bool {
	if err == nil {
		return false
	}

	if target == nil {
		return false
	}
	return strings.Contains(err.Error(), target.Error())
}

///
/// common errors
///

func NewInternalServerError(err error) error {
	return new(err.Error(), http.StatusInternalServerError, internali18n)
}

func NewInternalServerErrorf(format string, args ...interface{}) error {
	return new(fmt.Sprintf(format, args...), http.StatusInternalServerError, internali18n)
}

func NewInternalServerErrorw(msg string, err error) error {
	return Wrap(msg, new(err.Error(), http.StatusInternalServerError, internali18n))
}

func NewBadRequestError(err error) error {
	return new(err.Error(), http.StatusBadRequest, badRequesti18n)
}

func NewBadRequestErrorf(format string, args ...interface{}) error {
	return new(fmt.Sprintf(format, args...), http.StatusBadRequest, badRequesti18n)
}

func NewBadRequestErrorw(msg string, err error) error {
	return Wrap(msg, new(err.Error(), http.StatusBadRequest, badRequesti18n))
}

func NewNotFoundError(err error) error {
	return new(err.Error(), http.StatusNotFound, notFoundi18n)
}

func NewNotFoundErrorf(format string, args ...interface{}) error {
	return new(fmt.Sprintf(format, args...), http.StatusNotFound, notFoundi18n)
}

func NewNotFoundErrorw(msg string, err error) error {
	return Wrap(msg, new(err.Error(), http.StatusNotFound, notFoundi18n))
}

func NewForbiddenError(err error) error {
	return new(err.Error(), http.StatusForbidden, forbiddeni18n)
}

func NewForbiddenErrorf(format string, args ...interface{}) error {
	return new(fmt.Sprintf(format, args...), http.StatusForbidden, forbiddeni18n)
}

func NewForbiddenErrorw(msg string, err error) error {
	return Wrap(msg, new(err.Error(), http.StatusForbidden, forbiddeni18n))
}

func NewUnauthorizedError(err error) error {
	return new(err.Error(), http.StatusUnauthorized, unauthorizedi18n)
}

func NewUnauthorizedErrorf(format string, args ...interface{}) error {
	return new(fmt.Sprintf(format, args...), http.StatusUnauthorized, unauthorizedi18n)
}

func NewUnauthorizerErrorw(msg string, err error) error {
	return Wrap(msg, new(err.Error(), http.StatusUnauthorized, unauthorizedi18n))
}
