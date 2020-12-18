package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/juang77/Gabotor/bd"
)

//ObtenerBanner es la funcion para obtener en la BD el Banner de un usuario.
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID ", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "ocurrió un error al intentar buscar el registro "+err.Error(), http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uplads/banners/" + perfil.Banner)
	if err != nil {
		http.Error(w, "Imagen no encontrada ", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen ", http.StatusBadRequest)
	}

}
