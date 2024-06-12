package Model

type Data struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Total string `json:"Total"`
	Libre string `json:"Libre"`
	EnUso string `json:"Uso"`
	Porc  string `json:"Porcentaje"`
}

type CPU struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Pid   string `json:"Pid"`
	Name  string `json:"Pid"`
	State string `json:"Pid"`
	Rss   string `json:"Pid"`
	Uid   string `json:"Pid"`
}
