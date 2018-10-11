package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/hcl"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("unable to read from stdin: %s", err)
	}

	var v interface{}
	err = hcl.Unmarshal(input, &v)
	if err != nil {
		log.Fatalf("unable to parse HCL: %s", err)
	}

	json, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf("unable to marshal json: %s", err)
	}

	fmt.Println(string(json))
}
