package command

type UserEditInputData struct {
	Name string `json:"name" binding:"required,min=1,max=256"`
}
