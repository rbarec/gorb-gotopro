package pg02

import (
	"log"
)

type PersonaModel struct {
	//	Id 		  bson.ObjectId `bson:"_id" json:"id" `
	Id       uint16 `pg:"id" json:"id" `
	Cdn      string `pg:"cdn"    json:"cdn"`
	IdenTyp  string `pg:"identyp"    json:"identTyp"`
	Iden     string `pg:"ident"    json:"Ident"`
	Names    string `pg:"nmes"   json:"Names"`
	SurNames string `pg:"snames" json:"SurNames"`
	//	Fnac      time.Time `pg:"Fnac" json:"Fnac"`

}

func InfoModel() string {
	log.Println("INFO")
	return "gotoPRO -> persist01 -> pg01"
}


func (obj *PersonaModel) toStringName() string {
    return obj.Names + " "+obj.SurNames 
}
