package business

type Business struct {
	Message    string       `json:"message"`
	Code       int          `json:"code"`
	Data       BusinessData `json:"data"`
	Request_id string       `json:"request_id"`
	Result     bool         `json:"result"`
}

type BusinessData struct {
	Count int     `json:"count"`
	Info  []Binfo `json:"info"`
}

type Binfo struct {
	Bk_biz_id           int    `json:"bk_biz_id"`
	Default             int    `json:"default"`
	Language            string `json:"language"`
	Life_cycle          string `json:"life_cycle"`
	Bk_biz_developer    string `json:"bk_biz_developer"`
	Bk_biz_maintainer   string `json:"bk_biz_maintainer"`
	Bk_biz_tester       string `json:"bk_biz_tester"`
	Time_zone           string `json:"time_zone"`
	App_manager_org     string `json:"app_manager_org"`
	Biz_type            string `json:"biz_type"`
	Create_time         string `json:"create_time"`
	Biz_manager         string `json:"biz_manager"`
	Bk_supplier_account string `json:"bk_supplier_account"`
	Operator            string `json:"operator"`
	Bk_biz_productor    string `json:"bk_biz_productor"`
	Bk_biz_name         string `json:"bk_biz_name"`
	Last_time           string `json:"last_time"`
	Bk_supplier_id      int    `json:"bk_supplier_id"`
}
