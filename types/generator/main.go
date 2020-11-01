package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type MyMsg struct {
	PackageOrig        string
	SubfolderOrig      string
	MsgNameOrig        string
	PackageCamelCase   string
	SubfolderCamelCase string
	MsgNameCamelCase   string
	CGoTypeForCast     string
	GoTypeForCast      string
}

type EquivalentTypes struct {
	Gotype string
	Ctype  string
}

func replaceUnderscoreAndCamelCase(input string) string {
	splittedString := strings.Split(input, "_")
	retValue := ""

	for _, splitted := range splittedString {
		if strings.Contains(splitted, "ui") {
			//Special case for UInt16 types...
			splittedChanged := strings.ReplaceAll(splitted, "uint", "UInt")
			retValue = retValue + splittedChanged

		} else {
			retValue = retValue + strings.Title(splitted)
		}

	}
	return retValue
}

func getGoTypeForCast(myMsg *MyMsg, typesTable []EquivalentTypes) {

	for _, mytype := range typesTable {

		if myMsg.MsgNameOrig == "char" {
			// myMsg.GoTypeForCast = "string"
			// myMsg.CGoTypeForCast = "C.GoString"

		} else {
			if mytype.Gotype == myMsg.MsgNameOrig {
				myMsg.GoTypeForCast = mytype.Gotype
				myMsg.CGoTypeForCast = mytype.Ctype
			}
		}

	}
}

func getFileAsString(myMsg *MyMsg) string {

	myMsg.PackageCamelCase = replaceUnderscoreAndCamelCase(myMsg.PackageOrig)
	myMsg.SubfolderCamelCase = replaceUnderscoreAndCamelCase(myMsg.SubfolderOrig)
	myMsg.MsgNameCamelCase = replaceUnderscoreAndCamelCase(myMsg.MsgNameOrig)

	retString :=
		`package types
/////////////////////////////////////////////////////
//// THE CONTENT OF THIS FILE HAS BEEN AUTOGENERATED
/////////////////////////////////////////////////////
// #cgo CFLAGS: -I/opt/ros/foxy/include
// #cgo LDFLAGS: -L/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -l{{.PackageOrig}}__rosidl_generator_c -l{{.PackageOrig}}__rosidl_typesupport_c
// #include "msg_types.h"
import "C"
import "unsafe"


type {{.PackageCamelCase}}{{.MsgNameCamelCase}} struct {
	data    *C.{{.PackageOrig}}__{{.SubfolderOrig}}__{{.MsgNameCamelCase}}
	MsgType MessageTypeSupport
}

func (msg *{{.PackageCamelCase}}{{.MsgNameCamelCase}}) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *{{.PackageCamelCase}}{{.MsgNameCamelCase}}) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *{{.PackageCamelCase}}{{.MsgNameCamelCase}}) InitMessage() {
	msg.data = C.init_{{.PackageOrig}}_{{.SubfolderOrig}}_{{.MsgNameCamelCase}}()
	msg.MsgType = GetMessageTypeFrom{{.PackageCamelCase}}{{.MsgNameCamelCase}}()
}

func (msg *{{.PackageCamelCase}}{{.MsgNameCamelCase}}) Set{{.MsgNameCamelCase}}(data {{.MsgNameOrig}}) {
	//TODO: to implement the setter
	msg.data.data = {{.CGoTypeForCast}}(data)
}


func (msg *{{.PackageCamelCase}}{{.MsgNameCamelCase}}) Get{{.MsgNameCamelCase}}() {{.MsgNameOrig}} {
	return {{.GoTypeForCast}}(msg.data.data)
}

func (msg *{{.PackageCamelCase}}{{.MsgNameCamelCase}}) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue:=""
	return myRetValue
}

func (msg *{{.PackageCamelCase}}{{.MsgNameCamelCase}}) DestroyMessage() {
	C.destroy_{{.PackageOrig}}_{{.SubfolderOrig}}_{{.MsgNameCamelCase}}(msg.data)
}

func GetMessageTypeFrom{{.PackageCamelCase}}{{.MsgNameCamelCase}}() MessageTypeSupport {
	return MessageTypeSupport{C.get_message_type_from_{{.PackageOrig}}_{{.SubfolderOrig}}_{{.MsgNameCamelCase}}()}
}
`
	retValue2 := strings.Trim(retString, "\t")
	return retValue2
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func main() {
	t := template.New("MyGoROS2Package")

	equivalentypes := []EquivalentTypes{
		{"int8", "C.schar"},
		{"int16", "C.short"},
		{"int32", "C.int"},
		{"int64", "C.long"},
		{"float32", "C.float"},
		{"float64", "C.double"},
		{"uint8", "C.uchar"},
		{"uint16", "C.ushort"},
		{"uint32", "C.uint"},
		{"uint64", "C.ulong"},
		{"char", "C.GoString"},
		{"bool", "C.bool"},
		{"byte", "C.uchar"},
	}

	msgs := []MyMsg{
		{"std_msgs", "msg", "int8", "", "", "", "", ""},
		{"std_msgs", "msg", "int16", "", "", "", "", ""},
		{"std_msgs", "msg", "int32", "", "", "", "", ""},
		{"std_msgs", "msg", "int64", "", "", "", "", ""},
		{"std_msgs", "msg", "uint8", "", "", "", "", ""},
		{"std_msgs", "msg", "uint16", "", "", "", "", ""},
		{"std_msgs", "msg", "uint32", "", "", "", "", ""},
		{"std_msgs", "msg", "uint64", "", "", "", "", ""},
		{"std_msgs", "msg", "float32", "", "", "", "", ""},
		{"std_msgs", "msg", "float64", "", "", "", "", ""},
		{"std_msgs", "msg", "bool", "", "", "", "", ""},
		// {"std_msgs", "msg", "char", "", "", "", "", ""},
		{"std_msgs", "msg", "byte", "", "", "", "", ""},
	}

	for _, msg := range msgs {

		path := "../msg_" + msg.PackageOrig + "_" + msg.MsgNameOrig + ".go"
		exist := Exists(path)

		if exist == false {

			getGoTypeForCast(&msg, equivalentypes)
			text := getFileAsString(&msg)

			t.Parse(text)
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
			if err != nil {
				fmt.Println(err)
				return
			}
			t.Execute(f, msg)

			defer f.Close()
		} else {
			fmt.Printf("File %v already exist, not overwritting\n " + path)
		}

	}
}
