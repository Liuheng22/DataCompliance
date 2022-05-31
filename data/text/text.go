package text

//测试数据
type Test struct {
	Data *Rowdata `json:"data"`
}

//一行数据
type Rowdata struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Id      string `json:"ID"`
}

//测试数据的返回数据
type Testres struct {
	Data *Rowres `json:"data"`
}

//测试数据的一行数据
type Rowres struct {
	Key      string    `json:"key"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Phone    string    `json:"phone"`
	Address  string    `json:"address"`
	Id       string    `json:"ID"`
	Problems []Problem `json:"problems"`
}

//测试数据的问题以及解决方案
type Problem struct {
	ID          string      `json:"id"`
	Col         string      `json:"col"`
	Seriousness string      `json:"seriousness"`
	Type        string      `json:"type"`
	Description string      `json:"description"`
	Fix         interface{} `json:"fix"`
}
