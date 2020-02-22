package input

// コントローラ、ユースケース、プレゼンターでバトンするために
// サードパーティのインターフェースをラップする
type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
	GetHeader(string) string
	GetQuery(string) (string, bool)
}