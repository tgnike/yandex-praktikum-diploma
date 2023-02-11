package server

import "errors"

var ErrUploadedByOtherUser = errors.New("номер заказа уже был загружен другим пользователем")
var ErrUploadedByUser = errors.New("номер заказа уже был загружен этим пользователем")
var ErrOrderNumberFormat = errors.New("неверный формат номера заказа")
