package main

import (
	"encoding/json"
	"fmt"
	"os"
)
type Elem struct{
	Name string `json:"name,omitempty"`
	Type string `json:"type"`
	Out  []interface{} `json:"out,omitempty"`
}
type IOput []Elem
type Function struct{
	Name string `json:"name,omitempty"`
	Type string `json:"type"`
	Inputs IOput `json:"inputs,omitempty"`
	Outputs IOput `json:"outputs,omitempty"`
	Payable bool  `json:"payable"`
	Constant bool `json:"constant,omitempty"`
}

func JSON_UNMARSHAL_ERROR(err error)(error){
	return fmt.Errorf("json unmarshal error. %s",err)
}

func test() (error) {
	file, err := os.Open("./../examples/gasless_send/verified_contract_abis/BuyerFund.abi")
	if err!=nil {
		fmt.Print(err)
		return err
	}

	fileinfo,_:=file.Stat()
	Size :=fileinfo.Size()
	out := make([]byte,Size)
	_,err = file.Read(out)


	type Abi []*Function

	var abi = new(Abi)

	if err:=json.Unmarshal(out, abi);err!=nil{
		return JSON_UNMARSHAL_ERROR(err)
	}
	fmt.Print("hi")
	fmt.Print(*abi)

	return nil
}

func main() {
	test()
}
