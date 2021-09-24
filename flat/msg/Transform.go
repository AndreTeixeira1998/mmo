// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package msg

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Transform struct {
	_tab flatbuffers.Struct
}

func (rcv *Transform) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Transform) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Transform) X() float64 {
	return rcv._tab.GetFloat64(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Transform) MutateX(n float64) bool {
	return rcv._tab.MutateFloat64(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Transform) Y() float64 {
	return rcv._tab.GetFloat64(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}
func (rcv *Transform) MutateY(n float64) bool {
	return rcv._tab.MutateFloat64(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

func CreateTransform(builder *flatbuffers.Builder, x float64, y float64) flatbuffers.UOffsetT {
	builder.Prep(8, 16)
	builder.PrependFloat64(y)
	builder.PrependFloat64(x)
	return builder.Offset()
}