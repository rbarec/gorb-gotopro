package pg02

import (
	"testing"
    "fmt"
)

func TestInfo(t *testing.T) {
	fmt.Println( "TESTING pg01" )
    want := "1"
    fmt.Println( Info() )
    if got := Info(); got != want {
        t.Errorf("Proverb() = %q, want %q", got, want)
    }
}

//func TestSelect(t *testing.T) {
//	fmt.Println( "\n\n\n TESTING PG01 TestSelect" )
////	strcon := "postgresql://postgres:rbblrnoonm@localhost:5432/plataform"
//	strcon := "postgresql://postgres:rbblrnoonm@localhost:5432/plataform?sslmode=disable"
//	want := "1"
//	InitDB(strcon)
//	AllPersonaModel2()
//    if got := Info(); got != want {
//        t.Errorf("Proverb() = %q, want %q", got, want)
//    }
//}


func TestSelectPointer(t *testing.T) {
	fmt.Println( "\n\n\n TESTING PG01 TestSelect Pointer" )
//	strcon := "postgresql://postgres:rbblrnoonm@localhost:5432/plataform"
	strcon := "postgresql://postgres:rbblrnoonm@localhost:5432/plataform?sslmode=disable"
	want := "1"
	InitDB(strcon)
	Dao_getAll(want)
    if got := Info(); got != want {
        t.Errorf("Proverb() = %q, want %q", got, want)
    }
}