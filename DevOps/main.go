package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Ruta para servir la pagina principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
            <html>
                <head>
                    <title>DevOPS TFA</title>
                </head>
                <body style="font-family: Arial; text-align: left; margin-top: 50px; color: blue;">
                    <h1>Hola soy alumno de la UOC</h1>
                    <img src="/static/Logotipo_UOC.jpg" alt="Mi foto" width="300">
                </body>
            </html>
        `)
	})

	// Servir archivos estaticos, como imagenes
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Iniciar servidor
	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
