package services

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

func VigilarArchivo(ruta string) error {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	// Vigilar la carpeta
	err = watcher.Add(ruta)
	if err != nil {
		return err
	}

	fmt.Println("Vigilando carpeta:", ruta)

	var ultimoEvento time.Time

	for {

		select {

		case evento := <-watcher.Events:

			// Solo nos interesa actual.json
			if filepath.Base(evento.Name) != "actual.json" {
				continue
			}

			// Solo cuando el archivo fue escrito
			if evento.Op&fsnotify.Write == fsnotify.Write {

				// Evitar múltiples eventos por un mismo guardado
				if time.Since(ultimoEvento) < 500*time.Millisecond {
					continue
				}

				ultimoEvento = time.Now()

				fmt.Println("\nArchivo modificado:", evento.Name)

				err := EjecutarComparacion()
				if err != nil {
					fmt.Println("Error:", err)
				}
			}

		case err := <-watcher.Errors:

			fmt.Println("Error:", err)

		}

	}
}
