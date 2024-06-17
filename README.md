## Manual Técnico del Proyecto 1
 LABORATORIO SISTEMAS OPERATIVOS 1
 Sección: A
#
#
#

### **Autores:**
- Genesis Nahomi Aparicio Acan - Carnet: 202113293
- Richard Alexandro Marroquin Arana - Carnet: 202102894
#
#
#
### **Auxiliar:**
- Daniel Velásquez



### **Título:**

### Manual Técnico del Proyecto 1
#
#
#
---


## Índice
1. [Introducción](##introducción)
2. [Web UI](#web-ui)
3. [Backend](#base-de-datos)
4. [DataController](#módulos)
5. [Conn](#instalación-y-configuración)
6. [Instance](#guía-de-uso)
7. [Data](#conclusión)
8. [Router](#conclusión)
9. [Modulo CPU](#conclusión)
10. [Modulo RAM](#conclusión)
11. [Main](#conclusión) 
12. [Librerias](#conclusión)
13. [Comandos de Inicio](#conclusión)
#
#
---
#
#
## Introducción
Este proyecto tiene como objetivo principal implementar un sistema de monitoreo de recursos del sistema y gestión de procesos empleando varias tecnologías y lenguajes de programación. El
sistema resultante permitirá obtener información clave sobre el rendimiento del computador,
procesos en ejecución y su administración a través de una interfaz amigable.
#
#
#

---
## Web UI
## Importaciones:

El código comienza importando varios módulos y componentes

- **React**: Importa la biblioteca React.
- **useState** y **useEffect**: Hooks de React para manejar el estado y los efectos secundarios.
- **createProcess** y **deleteProcess**: Funciones que probablemente interactúan con una API para crear y eliminar procesos.
- **PASTEL**, **TablaProcesos** y **TablaPid**: Componentes personalizados definidos en otros archivos.

![Menu](https://i.ibb.co/vqJ1Vhq//image.png)


## Función App:

La función App es el componente principal de la aplicación. Renderiza la interfaz de usuario y maneja el estado.

- Dentro de App, se definen varios estados utilizando el hook `useState`. Estos estados incluyen `actionButton`, `pid`, `entradaTexto` y `procesos`.
- La función `createProc` se llama cuando se hace clic en el botón “Crear proceso”. Realiza una solicitud a la API para crear un proceso y actualiza el estado con el ID del proceso creado.
- La función `deleteProc` se llama cuando se hace clic en el botón “Eliminar proceso”. Realiza una solicitud a la API para eliminar un proceso y actualiza el estado de los procesos existentes.

## Renderizado de componentes:

La función `return` renderiza varios elementos:

-  “SO1-Proyecto1-2024”.
- El componente `PASTEL`.
- Un contenedor con un botón para crear procesos, un campo de entrada y un botón para eliminar procesos.
- Los componentes `TablaProcesos` y `TablaPid`.

## Componentes personalizados:

Los componentes `PASTEL`, `TablaProcesos` y `TablaPid` deben estar definidos en otros archivos. Sin ver su implementación, no puedo proporcionar detalles específicos sobre su funcionalidad.



### Características Principales
- **Diseño Responsivo**: La interfaz se adapta a diferentes tamaños de pantalla
- **Gestión de Datos**: Formularios y vistas para gestionar los datos almacenados en la base de datos.


#

---
#
#
#
# DataController
#
### Paquetes e Importaciones:

El código está organizado en un paquete llamado `Controller`.

Importa los paquetes necesarios, incluyendo `Backend/Instance` y `Backend/Model`.

### Funciones:

- **InsertData(nameCol string, dataParam ...):**
  Inserta datos en una colección específica (`nameCol`).
  Crea una estructura `Model.Data` con los campos `Total`, `Libre`, `EnUso` y `Porc`.
  Inserta la estructura en la colección.

- **InsertData2(nameCol string, dataParam1 ...):**
  Similar a `InsertData`, pero para una estructura de datos diferente (`Model.CPU`).
  Los campos incluyen `Pid`, `Name`, `State`, `Padre`, `Rss` y `Uid`.

- **InsertData1(nameCol string) error:**
  Elimina toda la colección especificada por `nameCol`.
  Devuelve un error si ocurre alguno.

- **InsertData22(nameCol string, dataParam string):**
  Inserta una estructura `Model.Prueba` con un campo `Percent` en la colección.

### Interacción con la Base de Datos:

El código se conecta a una base de datos  MongoDB utilizando el paquete `Instance`.
Realiza operaciones de inserción y eliminación en las colecciones especificadas.

### Instalaciones:

- Asegúrate de que los paquetes `Backend/Instance` y `Backend/Model` estén disponibles y contengan las definiciones necesarias.
- Verifica que la conexión a la base de datos esté configurada correctamente antes de utilizar este código.

![Menu](https://i.ibb.co/cvsRB6y/image.png)

### ----------------------------------------------------------------------------------------------



# Conn

## Paquetes e Importaciones:

El código importa los paquetes necesarios para trabajar con MongoDB, como `"Backend/Instance"`, `"context"`, `"log"`, `"os"`, `"time"`, y `"go.mongodb.org/mongo-driver/mongo"`.

## Función Connect()

Esta función se utiliza para establecer una conexión a la base de datos.


1. **Leer las variables de entorno:**
   - Dirección del servidor
   - Puerto
   - Nombre de la base de datos

   Estas variables se leen desde el archivo `.env`.

2. **Crear un cliente MongoDB:**
   - Utiliza la cadena de conexión especificada en las variables de entorno.

3. **Establecer un contexto:**
   - Define un tiempo de espera de 05 segundos para la conexión.

4. **Conectar al cliente a la base de datos:**
   - Conecta el cliente a la base de datos especificada.

5. **Asignar el cliente y la base de datos:**
   - Asigna el cliente y la base de datos a la instancia `Instance.Mg` (probablemente definida en otro lugar de tu código).

## Conectividad:

- Las variables de entorno (como `DB_HOST`, `DB_PORT` y `DB_NAME`) con la configuracion correctamente antes de ejecutar este código.
- Biblioteca `go.mongodb.org/mongo-driver` 

---
# Instance
![Menu](https://i.ibb.co/tH67yrX/image.png)

## Estructura MongoInstance:

Define una estructura llamada `MongoInstance` con dos campos:
- `Client` (cliente de MongoDB)
- `Db` (base de datos)

## Función Connect():

Esta función establece la conexión con la base de datos MongoDB:

1. Lee las variables de entorno (como la dirección del servidor, el puerto y el nombre de la base de datos) desde el archivo `.env`.
2. Crea un cliente MongoDB utilizando la URI generada a partir de las variables de entorno.
3. Establece un contexto con un tiempo de espera de 30 segundos para la conexión.
4. Conecta el cliente a la base de datos especificada.
5. Asigna la instancia de `MongoInstance` al paquete `Instance.Mg`.
6. Devuelve un error si algo sale mal durante la conexión.







---

---
# Data


**Descripción:** Esta estructura representa información relacionada con el uso de recursos .

**Campos:**
- **ID:** Identificador único (opcional).
- **Total:** Almacena el valor total de algún recurso.
- **Libre:** Representa la cantidad disponible o libre de ese recurso.
- **EnUso:** Indica la cantidad actual en uso.
- **Porc:** Contiene el porcentaje de uso en relación al total.

### CPU

**Descripción:** Esta estructura parece estar relacionada con información sobre procesos de la CPU.

**Campos:**
- **ID:** Identificador único (opcional).
- **Pid:** ID del proceso.
- **Name:** Nombre del proceso.
- **State:** Estado actual del proceso.
- **Padre:** ID del proceso padre (si existe).
- **Rss:** Tamaño de la memoria residente del proceso.
- **Uid:** Identificador único del usuario asociado al proceso.



**Campos:**
- **ID:** Identificador único.
- **Percent:** Almacena un valor de porcentaje (quizás para pruebas o cálculos).

![Menu](https://i.ibb.co/TcbYvHN/image.png)





---
# Router

#
#
#

# Descripción del Código
#
#
#

### Definiciones de Estructuras (struct)

El código define varios tipos de estructuras personalizadas, como `RamInfo`, `DatosRam`, `InfoSistema`, `Proceso` y `Hijo`. Estas estructuras representan diferentes estructuras de datos utilizadas en la aplicación.

### Función de Configuración (Setup)

La función `Setup` inicializa el servidor web utilizando el framework Fiber. Define varias rutas que manejan diferentes solicitudes HTTP.

### Manejadores de Rutas

![Menu](https://i.ibb.co/pZV4zhm/image.png)

### `/insertRam`

Obtiene información sobre la RAM (total, en uso, libre y porcentaje) y la devuelve como JSON.

### `/InsertCPU`

Obtiene información sobre la CPU (porcentaje de uso y detalles de procesos) y la devuelve como JSON.

### `/insertProcess`

Inicia un proceso en segundo plano (usando el comando `sleep`) y devuelve su ID de proceso (PID).

### `/delProcess`

Finaliza un proceso especificado por su PID.
##
#
# ------------------------------------------------------------------------
# Modulos 
#
### Modulos CPU

![Menu](https://i.ibb.co/YX9WDMt/image.png)
1.  **Encabezados y declaración del módulo:**
    
    -   Incluye los encabezados necesarios para el módulo del kernel.
    -   Define la licencia, autor y descripción del módulo.
2.  **Función  `calcularPorcentajeCpu`:**
    
    -   Abre el archivo  `/proc/stat`  en modo lectura.
    -   Lee la información de uso de la CPU desde el archivo.
    -   Calcula el porcentaje de uso de la CPU utilizando matemáticas enteras.
    -   Cierra el archivo.
3.  **Función  `escribir_a_proc`:**
    
    -   Calcula el porcentaje de uso de la CPU llamando a  `calcularPorcentajeCpu`.
    -   Escribe el porcentaje de uso de la CPU en el archivo de secuencia.
    -   Itera a través de todos los procesos en ejecución:
        -   Obtiene el uso de memoria (RSS) para cada proceso.
        -   Escribe información sobre el proceso, como PID, nombre, estado y UID.
        -   Si el proceso tiene hijos, también escribe información sobre ellos.
        -   
4.  **Función  `for_each_process(cpu)`**:
    
    -   Esta función itera a través de todos los procesos en el sistema.
    -   Para cada proceso, verifica si tiene un espacio de memoria (`mm`). Si lo tiene, calcula la cantidad de memoria residente (`rss`) en páginas.
    -   Luego, crea una estructura JSON con información relevante sobre el proceso, como su ID (`pid`), nombre (`name`), estado (`state`),  `rss`  y UID del usuario (`uid`).
    -   También incluye información sobre los procesos secundarios (hijos) de cada proceso.
5.  **Función  `abrir_aproc(struct inode *inode, struct file *file)`**:
    
    -   Esta función se utiliza para abrir el archivo  `/proc/cpu_so1_jun2024`.
    -   Específicamente, se llama cuando se accede a este archivo.
6.  **Estructura  `archivo_operaciones`**:
    
    -   Define las operaciones que se pueden realizar en el archivo  `/proc/cpu_so1_jun2024`.
    -   En este caso, especifica que al abrir el archivo, se debe llamar a la función  `abrir_aproc`, y al leer desde el archivo, se debe usar la función  `seq_read`.
7.  **Funciones  `modulo_init()`  y  `modulo_cleanup()`**:
    
    -   `modulo_init()`  se ejecuta cuando se carga el módulo del kernel.
    -   Crea el archivo  `/proc/cpu_so1_jun2024`  y registra las operaciones definidas en  `archivo_operaciones`.
    -   Imprime un mensaje de información en el registro del kernel.
    -   `modulo_cleanup()`  se ejecuta cuando se descarga el módulo del kernel.
    -   Elimina el archivo  `/proc/cpu_so1_jun2024`.
    -   Imprime otro mensaje de información en el registro del kernel. 
    
# ------------------------------------------------------------------------
### Modulos RAM

![Menu](https://i.ibb.co/dfLvHQL/image.png)
1.  **Encabezado y metadatos**:
    
    -   El código comienza con los encabezados estándar de un módulo del kernel.
    -   Se especifica la licencia, descripción y autores del módulo.
2.  **Función  `mostrar_info_ram`**:
    
    -   Esta función se llama cuando se accede al archivo  `/proc/ram_so1_jun2024`.
    -   Utiliza  `si_meminfo(&inf)`  para obtener información sobre la memoria del sistema.
    -   Calcula la cantidad total de RAM (`total`) y la cantidad utilizada (`usado`).
    -   También calcula el porcentaje de RAM utilizada (`porcentaje_usado`).
    -   Luego, formatea esta información en un objeto JSON y la escribe en el archivo  `/proc`.
3.  **Función  `abrir_info_ram`**:
    
    -   Esta función se llama cuando se abre el archivo  `/proc/ram_so1_jun2024`.
    -   Permite el acceso al archivo.
4.  **Estructura  `proc_ops_info_ram`**:
    
    -   Define las operaciones que se pueden realizar en el archivo  `/proc/ram_so1_jun2024`.
    -   Incluye las funciones de apertura, lectura, búsqueda y liberación.
5.  **Funciones  `modulo_init`  y  `modulo_cleanup`**:
    
    -   `modulo_init()`  se ejecuta cuando se carga el módulo del kernel.
    -   Crea el archivo  `/proc/ram_so1_jun2024`  y registra las operaciones definidas.
    -   Imprime un mensaje en el registro del kernel.
    -   `modulo_cleanup()`  se ejecuta cuando se descarga el módulo del kernel.
    -   Elimina el archivo  `/proc/ram_so1_jun2024`.
    -   Imprime otro mensaje en el registro del kernel.
# ------------------------------------------------------------------------

# Docker

Para la implementación de la aplicación, se utilizó Docker y Docker-Compose. Docker es una plataforma de contenedores que permite empaquetar y distribuir aplicaciones junto con todas sus dependencias, garantizando su ejecución de manera consistente en diferentes entornos.

## Docker-Compose

Docker-Compose es una herramienta que facilita la definición y ejecución de aplicaciones que utilizan múltiples contenedores de Docker como si fueran una sola aplicación. A continuación, se presenta el archivo docker-compose.yml utilizado para la aplicación

# ------------------------------------------------------------------------
# Funciones Auxiliares

### `getMem`

Obtiene información sobre la RAM, calcula el porcentaje libre y almacena los datos en una base de datos.

### `ObtenerRam`

Ejecuta un comando de shell (`cat /proc/ram_so1_jun2024`) para obtener detalles de la RAM.

### `parsearSalidaComando`

Analiza la salida del comando y convierte los datos en una estructura `RamInfo`.

##### Función `getcpu()`

Esta función llama a `obtenerCPUPorcentaje()` para obtener información sobre la CPU. Si hay un error, se registra en los registros.

##### Función `obtenerCPUPorcentaje()`

1. Ejecuta un comando de shell (`cat /proc/cpu_so1_jun2024`) para leer la información de la CPU.
2. Decodifica la salida JSON utilizando la función `parsearSalidaCPU()`.
3. Devuelve la información de la CPU o un error si algo sale mal.

##### Función `parsearSalidaCPU(output string) (*InfoSistema, error)`

1. Decodifica la salida JSON en una estructura `InfoSistema`.
2. Verifica si el campo `cpu_porcentaje` existe en el JSON y lo asigna al porcentaje de uso de la CPU.
3. Llama a `parsearProcesos(data)` para obtener detalles sobre los procesos en ejecución.

##### Función `parsearProcesos(data map[string]interface{}) ([]Proceso, error)`

1. Verifica si el campo `Procesos_existentes` existe y es una lista.
2. Itera sobre los procesos y llama a `parsearProceso(procesoData)` para obtener detalles de cada proceso.

##### Función `parsearProceso(data map[string]interface{}) (Proceso, error)`

1. Convierte y asigna los campos del proceso individual (como PID, nombre, estado, etc.).
2. También procesa los hijos de cada proceso llamando a `parsearHijo(childData)`.

##### Función `parsearHijo(data map[string]interface{}) (Hijo, error)`

Similar a `parsearProceso`, procesa los campos del hijo individual (como PID, nombre, estado, etc.).


---
## Main
# Importaciones

El código importa paquetes y módulos necesarios para la aplicación. 
Por ejemplo, "Backend/Database" y "Backend/Routes" son paquetes personalizados específicos de tu aplicación.

# Función main()

La función principal de la aplicación.

- Configura un servidor web utilizando Fiber.
- Conecta a la base de datos (presumiblemente utilizando el paquete Database).
- Define rutas y middleware (como CORS) para el servidor.
- Escucha en el puerto 8000.

#
#
#
#
---

#
#
#
#
#
#
# Librerias

```
Fiber: Un framework web rápido y ligero1.

encoding/json: Para codificar y decodificar datos en formato JSON.

fmt: Proporciona funciones para formatear y mostrar texto en la consola.

log: Se utiliza para registrar mensajes en la consola.

os/exec: Permite ejecutar comandos externos desde Go.
```
# Comandos de Inicio

#
#
```
sudo insmod file.ko  -- este es para incertar los ko de ram y cpu 

sudo rmmod file.ko     --- este es para eliminar el ko  se debe de hacer en el archivo

use DB                --nombre de la base de datos 

docker stop  DB      -- parar base de docker 

docker start DB     -- parar base de start 

df -h                --ver espacio en disco 

cd Proyecto/Proyecto1/Backend  -- ruta backend proyecto 

cd Proyecto SO1_JUN2024_PAREJA-4-/Proyecto1/Modulos/Ram

docker run -p 8000:8000 jdveloper/sol_backend_jun2024:lates
--descargar o corrrer la imagen de docker    
jdveloper/sol_backend_jun2024:latest  


sudo docker push  nahomiaparicio/so1_front_2024:2.0.0    -- subir imagen docker 

cd Docker-compose                        --carpeta docker compose

sudo docker compose up                   --correr docker compose     

docker run --it -p 5173:5173  nahomiaparicio/so1_front_2024          

sudo docker run --rm -it -p 80:80 naho

sudo docker run --rm -it -p 80:80 nahomiaparicio/front

udo docker run --rm -it -p 5173:5173 nahomiaparicio/pruebaf

sudo docker build -t nahomiaparicio/pruebaf2:5.0.0 .

sudo docker run --rm -it -p 5173:5173 nahomiaparicio/pruebaf2:9.3.0

ps aux | grep 2191      -- ver proceso 

kill -9  2191     -- matar poceso 

stress --vm 1 --vm-bytes 2G --timeout 30s     --estresar la ram 

sudo docker run -d -p 27017:27017 --name DB -v mongo-data:/data/db mongo   -jalar mongo 

stress --cpu $(nproc) --timeout 60                          ---estresar la cpu


```
## Conclusión
Este documento proporciona una visión general de los componentes y la estructura del proyecto. Para más detalles, consulta la documentación en cada módulo o contacta al equipo de desarrollo.

---

