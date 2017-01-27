package developers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mux"
)

type Developer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Language  string `json:"language"`
	Floor     int    `json:"floor"`
}

type Response struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

func GetDeveloper(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// db, err := sql.Open("mymysql", "tcp:127.0.0.1:3306*webservice/root/")
	db, err := sql.Open("sqlite3", "./developers.db")
	if err != nil {
		log.Fatal(err)
	}

	var query = "SELECT first_name, last_name, age, language," +
		"floor from developers WHERE id = ?"

	rows, err := db.Query(query, id)
	if err != nil {
		log.Fatal(err)
	}
	if rows.Next() {
		var developer Developer
		for rows.Next() {
			if err := rows.Scan(&developer.FirstName,
				&developer.LastName,
				&developer.Age,
				&developer.Language,
				&developer.Floor,
			); err != nil {
				log.Fatal(err)
			}
			fmt.Println(developer)
			developerString := fmt.Sprintf("Name : %s, Age: %d , Developing Language: %s, Floor: %d", developer.FirstName+
				" "+developer.LastName,
				developer.Age,
				developer.Language,
				developer.Floor,
			)
			createResponse(200, developerString, w)

		}
	} else {
		createResponse(404, "Requested developer not found", w)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
func createResponse(statusCode int, message string, w http.ResponseWriter) {
	res := Response{
		Code:   statusCode,
		Result: fmt.Sprintf("%s", message),
	}

	jsonMsg, err := json.Marshal(res)
	if err != nil {
		log.Fatal("Unable to marshal the json")
	}
	w.Header().Set("application/type", "json")
	w.WriteHeader(statusCode)
	w.Write(jsonMsg)
}
