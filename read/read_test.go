package read

import (
	"testing"
)

func TestRtype_Read(t *testing.T) {

	r := Rtype{file: "../fixture/readFile.json"}
	result, err := r.Read()
	if err != nil {
		t.Fatalf("Could not read ../fixture/readFile.json")
	}

	if len(result) < 10 {
		t.FailNow()
	}

	r.file = "invalid"
	result, err = r.Read()
	if err == nil {
		t.FailNow()
	}
	if r.e == nil {
		t.FailNow()
	}

}
