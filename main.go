package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func main() {
	const msgpOut1 = "data_msgp_out1"
	const msgpOut2 = "data_msgp_out2"
	const protoOut1 = "data_proto_out1"
	const protoOut2 = "data_proto_out2"

	msgpData1, err := InnerStructA_inst.MarshalMsg(nil)
	if err != nil {
		fmt.Printf("error marshalling msgp1; %v\n", err)
		return
	}
	fmt.Printf("LEN of msgp 1: %d\n", len(msgpData1))

	msgpData2, err := Struct1A_inst.MarshalMsg(nil)
	if err != nil {
		fmt.Printf("error marshalling msgp2; %v\n", err)
		return
	}
	fmt.Printf("LEN of msgp 2: %d\n", len(msgpData2))

	protoData1, err := proto.Marshal(&InnerStructB_inst)
	if err != nil {
		fmt.Printf("error marshalling proto1; %v\n", err)
		return
	}
	fmt.Printf("LEN of proto 1: %d\n", len(protoData1))

	protoData2, err := proto.Marshal(&StructB_inst)
	if err != nil {
		fmt.Printf("error marshalling proto2; %v\n", err)
		return
	}
	fmt.Printf("LEN of proto 2: %d\n", len(protoData2))
}
