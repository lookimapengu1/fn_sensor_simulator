package main

import (
	"context"
	"encoding/json"
	"io"
	"math/rand"

	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))
}

type Patient struct {
	PatientID string `json:"patientID"`
	Heartrate int `json:"heartrate"`
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {
	p := &Patient{PatientID: "001", Heartrate: (rand.Intn(60)+60)}
	json.NewEncoder(out).Encode(&p)
}
