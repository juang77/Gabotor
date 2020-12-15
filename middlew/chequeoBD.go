package middlew

import (
	"net/http"

	"github.com/juang77/Gabotor/bd"
)

//ChequeoDB es el meedleware que me permite conocer el estado de la BD.
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión Perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
