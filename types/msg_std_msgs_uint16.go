package types
/////////////////////////////////////////////////////
//// THE CONTENT OF THIS FILE HAS BEEN AUTOGENERATED
/////////////////////////////////////////////////////
// #cgo CFLAGS: -I/opt/ros/bouncy/include
// #cgo LDFLAGS: -L/opt/ros/bouncy/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
// #include "msg_types.h"
import "C"
import "unsafe"


type StdMsgsUInt16 struct {
	data    *C.std_msgs__msg__UInt16
	MsgType MessageTypeSupport
}

func (msg *StdMsgsUInt16) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsUInt16) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsUInt16) InitMessage() {
	msg.data = C.init_std_msgs_msg_UInt16()
	msg.MsgType = GetMessageTypeFromStdMsgsUInt16()
}

func (msg *StdMsgsUInt16) SetUInt16(data uint16) {
	//TODO: to implement the setter
	msg.data.data = C.ushort(data)
}

func (msg *StdMsgsUInt16) GetUInt16() uint16 {
	return uint16(msg.data.data)
}

func (msg *StdMsgsUInt16) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue:=""
	return myRetValue
}

func (msg *StdMsgsUInt16) DestroyMessage() {
	C.destroy_std_msgs_msg_UInt16(msg.data)
}

func GetMessageTypeFromStdMsgsUInt16() MessageTypeSupport {
	return MessageTypeSupport{C.get_message_type_from_std_msgs_msg_UInt16()}
}
