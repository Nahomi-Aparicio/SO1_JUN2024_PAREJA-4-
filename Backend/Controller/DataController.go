package Controller

import (
	"Backend/Instance"
	"Backend/Model"
	"context"
	"fmt"
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

func InsertData2(nameCol string, dataParam1 string, dataParam2 string, dataParam3 string, Padre2 string, dataParam4 string, dataParam5 string) {

	collection := Instance.Mg.Db.Collection(nameCol)

	doc := Model.CPU{Pid: dataParam1, Name: dataParam2, State: dataParam3, Padre: Padre2, Rss: dataParam4, Uid: dataParam5}

	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertData1(nameCol string) error {
	collection := Instance.Mg.Db.Collection(nameCol)

	err := collection.Drop(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("en poceso de guardar datos")
	return nil
}

func InsertData22(nameCol string, dataParam string) {
	collection := Instance.Mg.Db.Collection(nameCol)
	doc := Model.Prueba{Percent: dataParam}

	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
}
