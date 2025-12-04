package httpHandler

type Handlers struct {
	UserHandler *UserHandler
}

func NewHandler(userService UserService) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(userService),
	}
}
