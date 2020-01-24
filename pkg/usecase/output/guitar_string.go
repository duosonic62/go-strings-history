package output

type GuitarStringOutput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Maker       string `json:"maker"`
	ThinGauge   int    `json:"thinGauge"`
	ThickGauge  int    `json:"thickGauge"`
	Url         string `json:"url"`
}
