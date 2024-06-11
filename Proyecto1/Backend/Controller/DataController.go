package Controller

import (
	"Backend/Instance"
	"Backend/Model"
	"context"
	"log"
)

func InsertData(nameCol string, dataParam string, dataParam1 string, dataParam2 string, dataParam3 string) {
	collection := Instance.Mg.Db.Collection(nameCol)
	doc := Model.Data{Total: dataParam, Libre: dataParam1, EnUso: dataParam3, Porc: dataParam2}

	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
}
