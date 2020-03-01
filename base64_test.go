package base64

import(
	"testing"
)

type Codificaciones struct{
	limpia string
	codificada string
}

var cadenas = []Codificaciones{
	{"",""},
	{"f","Zg=="},
	{"fo","Zm8="},
	{"foo","Zm9v"},
	{"foob","Zm9vYg=="},
	{"fooba","Zm9vYmE="},
	{"foobar","Zm9vYmFy"},
    {"any carnal pleasure.","YW55IGNhcm5hbCBwbGVhc3VyZS4="},
    {"any carnal pleasure","YW55IGNhcm5hbCBwbGVhc3VyZQ=="},
    {"any carnal pleasur","YW55IGNhcm5hbCBwbGVhc3Vy"},
}

func TestCodifica(t *testing.T){
	for _, v := range cadenas{
		got := Codifica(v.limpia)
	 	if got != v.codificada {
			t.Fatalf("Expected %q, got %q", v.codificada, got)
	    }
	}
}

func TestDecodifica(t *testing.T){
	for _, v := range cadenas{
		got, _ := Decodifica(v.codificada)
	 	if got != v.limpia {
			t.Fatalf("Expected %q, got %q", v.limpia, got)
	    }
	}
}