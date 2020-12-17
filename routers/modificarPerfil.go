package routers

import (
	"encoding/json"
	"net/http"

	"github.com/juang77/Gabotor/bd"
	"github.com/juang77/Gabotor/models"
)

//ModificarPerfil is la funcion que permite la modificaciond e un perfil de usuario
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool

	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el Registro. Reintente nuevamente "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
