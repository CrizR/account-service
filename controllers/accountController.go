package controllers

type AccountController interface {
	Stop()
}

func NewAccountController() AccountController {

	return nil
}
