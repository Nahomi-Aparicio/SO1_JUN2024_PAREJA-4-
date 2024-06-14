package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"time"
)

type RamInfo struct {
	Total      uint64  `json:"totalRam"`
	EnUso      uint64  `json:"memoriaEnUso"`
	Porcentaje float32 `json:"porcentaje"`
}

type InfoSistema struct {
	CPUPorcentaje float32 `json:"cpu_porcentaje"`
}

func main() {
	// Crear un ticker para ejecutar el comando cada segundo
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	for range ticker.C {
		ramInfo, err := ObtenerRam()

		if err != nil {
			log.Println("Error al obtener la RAM:", err)
			continue
		}

		infoSistema, err := obtenerCPUPorcentaje()

		if err != nil {
			log.Println("Error al obtener el porcentaje de CPU:", err)
			continue
		}

		// Convertir campos a enteros
		total := int(ramInfo.Total)
		libre := int(ramInfo.EnUso)
		porcentaje := int(ramInfo.Porcentaje)
		enUso := total - libre
		porcentajeR := (float64(libre) * 100) / float64(total)
		porcentajeR = float64(int(porcentajeR*100)) / 100 // Redondear a dos decimales
		fmt.Println("------------------RAM---------------")
		// Imprimir la información de la RAM en la consola
		fmt.Printf("Total de Ram: %d kB\n", total)
		fmt.Printf("Memoria Ram uso: %d kB\n", enUso)
		fmt.Printf("Porcentaje libre: %d%%\n", porcentaje)
		fmt.Printf("libre : %d%%\n", libre)
		fmt.Println(porcentajeR)

		fmt.Println("------------------CPU---------------")
		fmt.Println(infoSistema.CPUPorcentaje)
		// Imprimir el porcentaje de CPU

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

func obtenerCPUPorcentaje() (*InfoSistema, error) {
	// Ejecutar el comando para leer el archivo
	cmd := exec.Command("sh", "-c", "cat /proc/cpu_so1_jun2024")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	// Decodificar el JSON
	infoSistema, err := parsearSalidaCPU(string(out))
	if err != nil {
		return nil, fmt.Errorf("error al decodificar JSON: %v", err)
	}

	return infoSistema, nil
}

func parsearSalidaCPU(output string) (*InfoSistema, error) {
	var infoSistema InfoSistema

	var data map[string]interface{}
	err := json.Unmarshal([]byte(output), &data)
	if err != nil {
		return nil, err
	}

	// Verificar si el campo "cpu_porcentaje" existe en el JSON
	porcentaje, ok := data["cpu_porcentaje"].(float64)
	if !ok {
		return nil, fmt.Errorf("campo 'cpu_porcentaje' no encontrado en el JSON")
	}

	infoSistema.CPUPorcentaje = float32(porcentaje)

	return &infoSistema, nil
}
