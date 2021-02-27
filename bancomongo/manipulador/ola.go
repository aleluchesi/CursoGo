package manipulador

import (
	"time"
	"net/http"
	"cursodego/bancomongo/model"
)

func Ola(w http.ResponseWriter, r *http.Request)  {
	hora := time.Now().Format("15:04:05")
	pagina := model.Pagina{}
	pagina.Hora = hora
	if err := Modelos.ExecuteTemplate(w, "ola.html", pagina); err != nil {
		http.Error(w, "Houve um erro na reinderização da página!", http.StatusServiceUnavailable)
	}
}