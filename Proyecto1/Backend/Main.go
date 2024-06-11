package main

import (
	"Backend/Controller"
	"Backend/Database"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type RamInfo struct {
	Total      uint64  `json:"totalRam"`
	EnUso      uint64  `json:"memoriaEnUso"`
	Porcentaje float32 `json:"porcentaje"`
}

func main() {
	app := fiber.New()

	if err := Database.Connect(); err != nil {
		log.Fatal("Error en", err)
	}

	getMem()

	if err := app.Listen(":8000"); err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 500)
}

func getMem() {
	for range time.Tick(time.Second * 2) {

		ramInfo, err := ObtenerRam()

		if err != nil {
			log.Println("Error al obtener la RAM:", err)
			continue
		}
		total := int(ramInfo.Total)
		libre := int(ramInfo.EnUso)
		porcentaje := int(ramInfo.Porcentaje)
		porcentajeR := (float64(libre) * 100) / float64(total)
		porcentajeR = float64(int(porcentajeR*100)) / 100 // Redondear a dos decimales
		enUso := total - libre

		fmt.Println("------------------RAM---------------")
		// Imprimir la información de la RAM en la consola
		fmt.Printf("Total de Ram: %d kB\n", total)
		fmt.Printf("Memoria Ram uso: %d kB\n", enUso)
		fmt.Printf("Porcentaje libre: %d%%\n", porcentaje)
		fmt.Printf("libre : %d%%\n", libre)
		fmt.Println(porcentajeR)

		DBtotal := strconv.Itoa(total)
		DBlibre := strconv.Itoa(libre)
		DBPorLibre := strconv.FormatFloat(porcentajeR, 'f', 2, 64)
		DBUso := strconv.Itoa(enUso)

		Controller.InsertData("datosRam", string(DBtotal), string(DBlibre), string(DBPorLibre), string(DBUso))

	}
}

func ObtenerRam() (*RamInfo, error) {
	cmd := exec.Command("sh", "-c", "cat /proc/ram_so1_jun2024")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	ramInfo, err := parsearSalidaComando(string(out))
	if err != nil {
		return nil, err
	}
	return ramInfo, nil
}

func parsearSalidaComando(output string) (*RamInfo, error) {
	var ramInfo RamInfo
	err := json.Unmarshal([]byte(output), &ramInfo)
	if err != nil {
		return nil, err
	}
	return &ramInfo, nil
}
