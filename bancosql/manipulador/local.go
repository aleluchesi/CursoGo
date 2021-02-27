package manipulador

import (
	"cursogo/bancosql/model"
	"cursogo/bancosql/repro"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Local(w http.ResponseWriter, r *http.Request) {

	local := model.Local{}
	codigotelefone, err := strconv.Atoi(r.URL.Path[7:])
	fmt.Println("Codigo ", codigotelefone)
	if err != nil {
		http.Error(w, "Houve um erro na converção de numero!", http.StatusBadRequest)
		fmt.Println("Erro no mumero enviado", err.Error())
		return
	}

	sql := "SELECT country, city, telcod FROM Local WHERE telcod = ?"
	linha, err := repro.Db.Queryx(sql, codigotelefone)
	if err != nil {
		http.Error(w, "Não foi possivel encontrra o número!", http.StatusInternalServerError)
		fmt.Println("Erro no mumero enviado", err.Error())
		return
	}

	for linha.Next() {
		err = linha.StructScan(&local)
		if err != nil {
			http.Error(w, "erro ao tratar a struct!", http.StatusInternalServerError)
			fmt.Println("erro no struct", err.Error())
			return
		}

		if err := ModeloLocal.ExecuteTemplate(w, "local.html", local); err != nil {
			http.Error(w, "Houve um erro na reinderização da página!", http.StatusServiceUnavailable)
		}

		sql = "INSERT INTO logquery (daterequest) values (?)"
		resultado, err := repro.Db.Exec(sql, time.Now().Format("02/01/2016 15:04:05"))
		if err != nil {
			fmt.Println("erro na inclusão do log", err.Error())
		}
		linhaafetadas, err := resultado.RowsAffected()
		if err != nil {
			fmt.Println("erro ao pegar o número de linhas pelo comando", err.Error())
		}
		fmt.Println("sucesso ", linhaafetadas)
	}

}
