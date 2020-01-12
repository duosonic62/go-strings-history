package command

type StringAddInputData struct {
	Name        string `json:"name" binding:"required,min=1,max=64"`
	Description string `json:"name" binding:"required,min=1,max=1024"`
	Maker       string `json:"maker" binding:"required,min=1,max=64"`
	ThinGauge   int    `json:"thinGauge" binding:"required,min=1,max=100"`
	ThickGauge  int    `json:"thickGauge" binding:"required,min=1,max=100"`
	Url         string `json:"name" binding:"required,min=1,max=256"`
}
