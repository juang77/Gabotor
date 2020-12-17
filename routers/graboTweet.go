package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/juang77/Gabotor/bd"
	"github.com/juang77/Gabotor/models"
)

//GraboTweet es la funcion para crear en la BD el registro de un Tweet.
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro. Reintente nuevamente "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el Tweet ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
