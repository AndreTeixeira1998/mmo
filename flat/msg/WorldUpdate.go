// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package msg

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type WorldUpdate struct {
	_tab flatbuffers.Table
}

func GetRootAsWorldUpdate(buf []byte, offset flatbuffers.UOffsetT) *WorldUpdate {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &WorldUpdate{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsWorldUpdate(buf []byte, offset flatbuffers.UOffsetT) *WorldUpdate {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &WorldUpdate{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *WorldUpdate) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *WorldUpdate) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *WorldUpdate) UserId() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *WorldUpdate) MutateUserId(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *WorldUpdate) Entities(obj *Entity, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *WorldUpdate) EntitiesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func WorldUpdateStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func WorldUpdateAddUserId(builder *flatbuffers.Builder, userId uint64) {
	builder.PrependUint64Slot(0, userId, 0)
}
func WorldUpdateAddEntities(builder *flatbuffers.Builder, entities flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(entities), 0)
}
func WorldUpdateStartEntitiesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func WorldUpdateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}