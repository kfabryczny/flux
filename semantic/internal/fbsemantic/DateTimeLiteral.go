// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbsemantic

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DateTimeLiteral struct {
	_tab flatbuffers.Table
}

func GetRootAsDateTimeLiteral(buf []byte, offset flatbuffers.UOffsetT) *DateTimeLiteral {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DateTimeLiteral{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *DateTimeLiteral) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DateTimeLiteral) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DateTimeLiteral) Loc(obj *SourceLocation) *SourceLocation {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SourceLocation)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *DateTimeLiteral) Value(obj *Time) *Time {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Time)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *DateTimeLiteral) TypType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DateTimeLiteral) MutateTypType(n byte) bool {
	return rcv._tab.MutateByteSlot(8, n)
}

func (rcv *DateTimeLiteral) Typ(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func DateTimeLiteralStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func DateTimeLiteralAddLoc(builder *flatbuffers.Builder, loc flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(loc), 0)
}
func DateTimeLiteralAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(value), 0)
}
func DateTimeLiteralAddTypType(builder *flatbuffers.Builder, typType byte) {
	builder.PrependByteSlot(2, typType, 0)
}
func DateTimeLiteralAddTyp(builder *flatbuffers.Builder, typ flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(typ), 0)
}
func DateTimeLiteralEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}