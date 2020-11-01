package types

/////////////////////////////////////////////////////
//// THE CONTENT OF THIS FILE HAS BEEN AUTOGENERATED
/////////////////////////////////////////////////////
// #cgo CFLAGS: -I/opt/ros/foxy/include
// #cgo LDFLAGS: -L/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
// #include "msg_types.h"
import "C"
import "unsafe"

type StdMsgsFloat64 struct {
	data    *C.std_msgs__msg__Float64
	MsgType MessageTypeSupport
}

func (msg *StdMsgsFloat64) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsFloat64) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsFloat64) InitMessage() {
	msg.data = C.init_std_msgs_msg_Float64()
	msg.MsgType = GetMessageTypeFromStdMsgsFloat64()
}

func (msg *StdMsgsFloat64) SetFloat64(data float64) {
	//TODO: to implement the setter
	msg.data.data = C.double(data)
}

func (msg *StdMsgsFloat64) GetFloat64() float64 {
	return float64(msg.data.data)
}

func (msg *StdMsgsFloat64) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue := ""
	return myRetValue
}

func (msg *StdMsgsFloat64) DestroyMessage() {
	C.destroy_std_msgs_msg_Float64(msg.data)
}

func GetMessageTypeFromStdMsgsFloat64() MessageTypeSupport {
	return MessageTypeSupport{C.get_message_type_from_std_msgs_msg_Float64()}
}
