package base64

import (
	"strings"
	"errors"
)

const caracteres string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func Codifica(original string) (resultado string) {
	f := []byte{3, 240, 15, 192}
	original = strings.TrimSpace(original)
	base := []byte(original)
	resultado = ""
	total := len(base) 
	extras := total%3
	seccion := 0
	
	if extras == 1 {
		base = append([]byte(original), 0,0)
		total+=2
	}
	if extras == 2 {
		base = append([]byte(original), 0)
		total+=1
	}
	for seccion < total {
		b1:= base[seccion:seccion+1]
		b2:= base[seccion+1:seccion+2]
		b3:= base[seccion+2:seccion+3]
		s1 := b1[0]>>2
		s2 := b1[0]&f[0]<<4+b2[0]&f[1]>>4
		s3 := b2[0]&f[2]<<2+b3[0]&f[3]>>6
		s4 := (b3[0]<<2)>>2
		resultado += string(caracteres[s1]) + string(caracteres[s2]) + string(caracteres[s3]) + string(caracteres[s4])
		seccion+=3
	}
	if extras == 1 {
		resultado = resultado[:len(resultado)-2]+"=="
	}
	if extras == 2 {
		resultado = resultado[:len(resultado)-1]+"="
	}
	return 
}

func Decodifica(codificada string) (resultado string, err error){
	err = errors.New("El tamaño de los datos de entrada no concuerda con la codificación.")
	f2 := []int{48, 15, 60, 3}
	resultado = ""
	var indices = make([]uint8, 0)
	codificada = strings.TrimSpace(codificada)
	total := len(codificada)
	if total%4 != 0{
		return
	}
	extras := strings.Count(codificada, "=")
	actual:=0
	for _, value := range codificada{
		indices = append(indices, uint8(strings.Index(caracteres, string(value))))
	}
	if extras > 0{
		if indices[len(indices)-1] == 255{
			indices[len(indices)-1] = 0
		}
		if indices[len(indices)-2] == 255{
			indices[len(indices)-2] = 0
		}
	}
	for actual<total{
		s1 := int(indices[actual])
		s2 := int(indices[actual+1])
		s3 := int(indices[actual+2])
		s4 := int(indices[actual+3])
		r1 := s1<<2+s2&f2[0]>>4
		r2 := (s2&f2[1])<<4+(s3&f2[2])>>2
		r3 := (s3&f2[3])<<6+s4
		resultado += string(r1)+string(r2)+string(r3)
		actual+=4
	}
	if extras > 0{
		if extras == 2{
			resultado = resultado[:len(resultado)-2]
		}else{
			resultado = resultado[:len(resultado)-1]
		}
	}
	err = nil
	return
}