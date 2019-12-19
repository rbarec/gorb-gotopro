package pg01

import (
 "database/sql"
    _ "github.com/lib/pq"
    "log"
)

// go mod init github.com/rbarec/gorb-gotopro/persist01/pg01
// go test

type PersonaModel struct {
//	Id 		  bson.ObjectId `bson:"_id" json:"id" `
	Id		  uint16 `pg:"id" json:"id" `
	Cdn		  string `pg:"cdn"    json:"cdn"`
	IdenTyp  string `pg:"identyp"    json:"identTyp"`
	Iden     string `pg:"ident"    json:"Ident"`
	Names     string `pg:"nmes"   json:"Names"`
	SurNames  string `pg:"snames" json:"SurNames"`
//	Fnac      time.Time `pg:"Fnac" json:"Fnac"`
	
}

// GLOBAL_VAR
var db *sql.DB

func InitDB(dataSourceName string) {
	
	log.Println("InitDB ("+dataSourceName+"")
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



func AllPersonaModel() ([]*PersonaModel, error) {
	log.Println(db.Stats())
    rows, err := db.Query("SELECT id, cdn, nmes, snmes  FROM plt_pnj_tb")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    bks := make([]*PersonaModel, 0)
    for rows.Next() {
        row := new(PersonaModel)
        err := rows.Scan(&row.Id, &row.Cdn, &row.Names, &row.SurNames)
        if err != nil {
            return nil, err
        }
        bks = append(bks, row)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return bks, nil
}




func AllPersonaModel2() () {
	var strQuery = "SELECT id, cdn, nmes, snames  FROM plt_pnj_tb"
	log.Println("QUERY ",strQuery)
    rows, err := db.Query(strQuery)
    if err != nil {
    	log.Println(err)
        return 
    }
    defer rows.Close()
	log.Println("2.- MAKE Array")
    bks := make([]*PersonaModel, 0)
    for rows.Next() {
        row := new(PersonaModel)
        err := rows.Scan(&row.Id, &row.Cdn, &row.Names, &row.SurNames)
        if err != nil {
        	log.Println(err)
            return 
        }
        log.Println("OBJ:: ",row)
        bks = append(bks, row)
    }
    if err = rows.Err(); err != nil {
    	log.Println(err)
        return 
    }
    return 
}



