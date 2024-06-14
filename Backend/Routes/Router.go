package Routes

import (
	"Backend/Controller"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RamInfo struct {
	Total      uint64  `json:"totalRam"`
	EnUso      uint64  `json:"memoriaEnUso"`
	Porcentaje float32 `json:"porcentaje"`
}

type DatosRam struct {
	Total       int     `json:"total"`
	Libre       int     `json:"libre"`
	PorcentajeR float64 `json:"porcentajeR"`
	EnUso       int     `json:"enUso"`
}
type InfoSistema struct {
	CPUPorcentaje float32   `json:"cpu_porcentaje"`
	Procesos      []Proceso `json:"procesos"`
}

type Proceso struct {
	Pid      int    `json:"pid"`
	Name     string `json:"name"`
	State    int    `json:"state"`
	Rss      int    `json:"rss"`
	Uid      int    `json:"uid"`
	Children []Hijo `json:"Hijo"`
}

type Hijo struct {
	Pid   int    `json:"pid"`
	Name  string `json:"name"`
	State int    `json:"state"`
	Padre int    `json:"padre"`
	Rss   int    `json:"rss"`
	Uid   int    `json:"uid"`
}

func Setup(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("Hello World")
	})

	//app.Get("/Prueba", func(ctx *fiber.Ctx) error {
	//		fmt.Println("estoy tratando de enviar algo")
	//		dataParam := strconv.Itoa(rand.Intn(100))
	//
	//		Controller.InsertData1("ramprueba", string(dataParam))
	//
	//		return ctx.Status(201).JSON(dataParam)
	//	})

	app.Get("/insertRam", func(ctx *fiber.Ctx) error {
		ramInfo := getMem()
		log.Println(ramInfo.EnUso)
		return ctx.JSON(fiber.Map{
			"status":  200,
			"uso":     ramInfo.EnUso,
			"libre":   ramInfo.Libre,
			"porcent": ramInfo.PorcentajeR,
			"total":   ramInfo.Total,
		})
	})

	app.Get("/InsertCPU", func(ctx *fiber.Ctx) error {
		infoSistema, err := getcpu()
		if err != nil {
			log.Fatal("Error al obtener información de la CPU:", err)
		}
		guardarCPUmongo(infoSistema)

		return ctx.JSON(fiber.Map{
			"status":   200,
			"cpu":      infoSistema.CPUPorcentaje,
			"procesos": infoSistema.Procesos,
		})
	})

	app.Get("/insertProcess", func(ctx *fiber.Ctx) error {
		log.Println("Insertando proceso")

		cmd := exec.Command("sleep", "infinity")
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
			"pid":     cmd.Process.Pid,
		})
	})

	app.Get("/killProcess", func(ctx *fiber.Ctx) error {
		pid := ctx.Query("pid")
		log.Println(pid)
		pidInt, err := strconv.Atoi(pid)
		if err != nil {
			log.Fatal(err)
		}

		cmd := exec.Command("kill", "-9", strconv.Itoa(pidInt))
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
		})
	})

}

func getMem() DatosRam {

	ramInfo, err := ObtenerRam()
	if err != nil {
		log.Println("Error al obtener la RAM:", err)

	}

	total := int(ramInfo.Total)
	libre := int(ramInfo.EnUso)
	//porcentaje := int(ramInfo.Porcentaje)
	porcentajeR := (float64(libre) * 100) / float64(total)
	porcentajeR = float64(int(porcentajeR*100)) / 100 // Redondear a dos decimales
	enUso := total - libre
	DBtotal := strconv.Itoa(total)
	DBlibre := strconv.Itoa(libre)
	DBPorLibre := strconv.FormatFloat(porcentajeR, 'f', 2, 64)
	DBUso := strconv.Itoa(enUso)

	Controller.InsertData("datosRam", string(DBtotal), string(DBlibre), string(DBPorLibre), string(DBUso))

	return DatosRam{
		Total:       total,
		Libre:       libre,
		PorcentajeR: porcentajeR,
		EnUso:       enUso,
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

func getcpu() (*InfoSistema, error) {
	infoSistema, err := obtenerCPUPorcentaje()

	if err != nil {
		log.Println("Error al obtener el porcentaje de CPU:", err)

	}

	return infoSistema, err
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

	procesos, err := parsearProcesos(data)
	if err != nil {
		return nil, err
	}
	infoSistema.Procesos = procesos

	return &infoSistema, nil
}

// Procesar el JSON para obtener los procesos y sus hijos
func parsearProcesos(data map[string]interface{}) ([]Proceso, error) {
	var procesos []Proceso

	// Verificar si el campo "Procesos_existentes" existe y es una lista
	if procesosData, ok := data["Procesos_existentes"].([]interface{}); ok {
		for _, p := range procesosData {
			if procesoData, ok := p.(map[string]interface{}); ok {
				proceso, err := parsearProceso(procesoData)
				if err != nil {
					return nil, err
				}
				procesos = append(procesos, proceso)
			}
		}
	} else {
		return nil, fmt.Errorf("campo 'Procesos_existentes' no encontrado o no es una lista")
	}

	return procesos, nil
}

// Procesar un proceso individual y sus hijos
func parsearProceso(data map[string]interface{}) (Proceso, error) {
	var proceso Proceso

	// Convertir y asignar cada campo del proceso
	if pid, ok := data["pid"].(float64); ok {
		proceso.Pid = int(pid)
	}
	if name, ok := data["name"].(string); ok {
		proceso.Name = name
	}
	if state, ok := data["state"].(float64); ok {
		proceso.State = int(state)
	}
	if rss, ok := data["rss"].(float64); ok {
		proceso.Rss = int(rss)
	}
	if uid, ok := data["uid"].(float64); ok {
		proceso.Uid = int(uid)
	}
	if childrenData, ok := data["children"].([]interface{}); ok {
		for _, h := range childrenData {
			if childData, ok := h.(map[string]interface{}); ok {
				child, err := parsearHijo(childData)
				if err != nil {
					return proceso, err
				}
				proceso.Children = append(proceso.Children, child)
			}
		}
	}

	return proceso, nil
}

// Procesar un hijo individual
func parsearHijo(data map[string]interface{}) (Hijo, error) {
	var hijo Hijo

	// Convertir y asignar cada campo del hijo
	if pid, ok := data["pid"].(float64); ok {
		hijo.Pid = int(pid)
	}
	if name, ok := data["name"].(string); ok {
		hijo.Name = name
	}
	if state, ok := data["state"].(float64); ok {
		hijo.State = int(state)
	}
	if padre, ok := data["pidPadre"].(float64); ok {
		hijo.Padre = int(padre)
	}
	if rss, ok := data["rss"].(float64); ok {
		hijo.Rss = int(rss)
	}
	if uid, ok := data["uid"].(float64); ok {
		hijo.Uid = int(uid)
	}

	return hijo, nil
}

func guardarCPUmongo(infoSistema *InfoSistema) error {
	Controller.InsertData1("datoscpu")

	for _, proceso := range infoSistema.Procesos {

		//fmt.Printf("PID: %d, Name: %s, State: %d, RSS: %d, UID: %d\n", proceso.Pid, proceso.Name, proceso.State, proceso.Rss, proceso.Uid)

		pp := strconv.Itoa(proceso.Pid)
		ss := strconv.Itoa(proceso.State)
		rr := strconv.Itoa(proceso.Rss)
		uu := strconv.Itoa(proceso.Uid)
		Controller.InsertData2("datoscpu", pp, string(proceso.Name), string(ss), string(""), string(rr), string(uu))
		// Imprimir también la información de los hijos del proceso

		for _, hijo := range proceso.Children {
			hijopp := strconv.Itoa(hijo.Pid)
			hijoss := strconv.Itoa(hijo.State)
			hijopd := strconv.Itoa(hijo.Padre)
			hijorr := strconv.Itoa(hijo.Rss)
			hijouu := strconv.Itoa(hijo.Uid)
			//fmt.Printf("\tHijo PID: %d, Name: %s, State: %d, RSS: %d, UID: %d\n", hijo.Pid, hijo.Name, hijo.State, hijo.Rss, hijo.Uid)
			Controller.InsertData2("datoscpu", string(hijopp), string(hijo.Name), string(hijoss), string(hijopd), string(hijorr), string(hijouu))
		}
	}
	fmt.Printf("en poceso de guardar datos")
	return nil
}
