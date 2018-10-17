package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"io"

	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))
}

type Patient struct {
	PatientID 			string 	`json:"patientID"`
	BloodPressureHigh	int		`json:"bloodPressureHigh"`
	BloodPressureLow	int		`json:"bloodPressureLow"`
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {
	p := &Patient{PatientID: "001", BloodPressureHigh: (rand.Intn(70)+80), BloodPressureLow: (rand.Intn(50)+50)}
	json.NewEncoder(out).Encode(&p)
}