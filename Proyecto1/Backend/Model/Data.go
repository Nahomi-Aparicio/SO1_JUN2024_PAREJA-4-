package Model

type Data struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Total string `json:"Total"`
	Libre string `json:"Libre"`
	EnUso string `json:"Uso"`
	Porc  string `json:"Porcentaje"`
}
