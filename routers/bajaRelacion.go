package routers

import (
	"net/http"

	"github.com/juang77/Gabotor/bd"
	"github.com/juang77/Gabotor/models"
)

//BajaRelacion is la funcion que permite eliminar relaciones entre usuarios
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar eliminar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado eliminar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
