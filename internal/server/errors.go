package server

type UploadedByUser struct {
	Status int
}

func NewUploadedByUser() error {
	return &UploadedByUser{Status: 200}
}

func (e *UploadedByUser) Error() string {
	return "номер заказа уже был загружен этим пользователем"
}

type UploadedByOtherUser struct {
	Status int
}

func NewUploadedByOtherUser() error {
	return &UploadedByOtherUser{Status: 409}
}

func (e *UploadedByOtherUser) Error() string {
	return "номер заказа уже был загружен другим пользователем"
}

var ErrUploadedByOtherUser = NewUploadedByOtherUser()
var ErrUploadedByUser = NewUploadedByUser()
