/*

CMD Commands:
 go mod init github.com/rbarec/gorb-gotopro/persist01/pg01
 go test

*/
package pg02


import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// go mod init github.com/rbarec/gorb-gotopro/persist01/pg01
// go test

// GLOBAL_VAR
var db *sql.DB

func InitDB(dataSourceName string) {
	log.Println("InitDB (" + dataSourceName + "")
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

func Info() string {
	log.Println("INFO")
	return "gotoPRO -> persist01 -> pg01"
}

func Dao_getAll(iden string) []*PersonaModel {
	fmt.Println("INVOKE::: Dao_getAll") //> rbfmt
	lPersonsPTR := []*PersonaModel{}
	fmt.Println(lPersonsPTR) //> rbfmt
	lPersonsPTR, _ = AllPersonaModel()

	//Iterar en un Slice
	// 	for _, per := range *lPersonsPTR {  .\pg02.go:52:22: invalid indirect of lPersonsPTR (type []*PersonaModel)
	for _, per := range lPersonsPTR {
		fmt.Printf("Nombre ::: %s  \n", per.toStringName() )
		fmt.Println("")
	}

	return lPersonsPTR
}

func AllPersonaModel() ([]*PersonaModel, error) {
	fmt.Println("INVOKE::: AllPersonaModel") //> rbfmt
	log.Println(db.Stats())
	rows, err := db.Query("SELECT id, cdn, nmes, snames  FROM plt_pnj_tb") ///>@query
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	log.Println("2.- MAKE Array")
	bks := make([]*PersonaModel, 0)
	for rows.Next() {
		row := new(PersonaModel)
		err := rows.Scan(&row.Id, &row.Cdn, &row.Names, &row.SurNames)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		log.Println(row)
		bks = append(bks, row)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return bks, nil
}

func AllPersonaModel2() {
	var strQuery = "SELECT id, cdn, nmes, snames  FROM plt_pnj_tb"
	log.Println("QUERY ", strQuery)
	rows, err := db.Query(strQuery)
	if err != nil {
		return
	}
	defer rows.Close()
	log.Println("2.- MAKE Array")
	bks := make([]*PersonaModel, 0)
	for rows.Next() {
		row := new(PersonaModel)
		err := rows.Scan(&row.Id, &row.Cdn, &row.Names, &row.SurNames)
		if err != nil {
			return
		}
		bks = append(bks, row)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return
	}
	return
}
