package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	log.Println("Soy un log")
	globals.ClientConfig = utils.IniciarConfiguracion("/Users/adriangil/tp0-golang/client/config.json")
	// validar que la config este cargada correctamente
	if globals.ClientConfig == nil {
		log.Println("Error al cargar la configuración")
		return
	}

	// loggeamos el valor de la config
	log.Println(globals.ClientConfig.Ip)

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config

	//utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)

	utils.GenerarYEnviarPaquete()
	// leer de la consola el mensaje
	// utils.LeerConsola()

	// generamos un paquete y lo enviamos al servidor
	//utils.GenerarYEnviarPaquete()

}
