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
	Name  string `json:"Name"`
	State string `json:"State"`
	Padre string `json:"Padre"`
	Rss   string `json:"Rss"`
	Uid   string `json:"Uid"`
}

type Prueba struct {
	ID      string `json:"id" bson:"_id,omitempty"`
	Percent string `json:"percent"`
}
