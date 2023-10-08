package protobufparser

//refs: https://protobuf.dev/programming-guides/encoding/
import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

type ProtobufParser struct {
	originalBin []byte
	bin         []byte
	offset      int
}

func NewProtobufParser(bin []byte) *ProtobufParser {
	return &ProtobufParser{originalBin: bin, bin: bin, offset: 0}
}

func (that *ProtobufParser) Query(fieldNumbers ...uint) (interface{}, error) {
	that.bin = that.originalBin
	that.offset = 0
	for i := 0; i < len(fieldNumbers); {
		fn, value, err := that.readField()
		if err != nil {
			return nil, err
		}
		if fn == fieldNumbers[i] {
			if i == len(fieldNumbers)-1 {
				return value, nil
			}
			that.bin = value.([]byte)
			that.offset = 0
			i++
		}
	}
	return nil, fmt.Errorf("not found field")
}
func (that *ProtobufParser) readField() (uint, interface{}, error) {
	tag, rLen := proto.DecodeVarint(that.bin[that.offset:])
	that.offset += rLen
	fieldNumber := tag >> 3
	wireType := tag & 0x7
	switch wireType {
	case 0: //Varint
		value, vLen := proto.DecodeVarint(that.bin[that.offset:])
		that.offset += vLen
		return uint(fieldNumber), value, nil
	case 2: //string|object|bytes
		sLen, vLen := proto.DecodeVarint(that.bin[that.offset:])
		that.offset += vLen
		value := that.bin[that.offset : that.offset+int(sLen)]
		that.offset += int(sLen)
		return uint(fieldNumber), value, nil
	default:
		return uint(fieldNumber), nil, fmt.Errorf("not supported wireType:%d", wireType)
	}
}
func (that *ProtobufParser) Put(index int, value *BaseValue) error {
	panic("not implement Put")
}
