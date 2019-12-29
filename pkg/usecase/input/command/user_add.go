package command

type UserAddInputData struct {
	Name string `json:"name" binding:"required,min=1,max=256"`
	UID  string `json:"name" binding:"required"`
}
