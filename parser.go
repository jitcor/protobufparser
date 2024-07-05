package protobufparser

//refs: https://protobuf.dev/programming-guides/encoding/
import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
)

type protobufParser struct {
	originalBin []byte
	bin         []byte
	offset      int
}

func newProtobufParser(bin []byte) *protobufParser {
	return &protobufParser{originalBin: bin, bin: bin, offset: 0}
}
func (that *protobufParser) remaining() []byte {
	return that.bin[that.offset:]
}
func (that *protobufParser) query(fieldNumbers ...uint) ([]interface{}, error) {
	that.bin = that.originalBin
	that.offset = 0
	fieldBuffer := make([]interface{}, 0)
	for i := 0; i < len(fieldNumbers); {
		fn, value, err := that.readField()
		if err != nil {
			if err == io.EOF {
				return fieldBuffer, nil
			}
			return nil, err
		}
		if fn == fieldNumbers[i] {
			if i == len(fieldNumbers)-1 {
				fieldBuffer = append(fieldBuffer, value)
				continue
			}
			if v, ok := value.([]byte); ok {
				that.bin = v
				that.offset = 0
				i++
			} else {
				return nil, fmt.Errorf("unexpected value type for field %d: %T", fn, value)
			}
		}
	}
	return nil, fmt.Errorf("not found field")
}
func (that *protobufParser) readField() (uint, interface{}, error) {
	if len(that.bin) == that.offset {
		return 0, nil, io.EOF
	}
	tag, rLen := proto.DecodeVarint(that.bin[that.offset:])
	if rLen == 0 {
		return 0, nil, io.EOF
	}
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

type object struct {
	value       any
	members     []*object
	isPrimitive bool
}

func PQuery(bin []byte, fieldNumbers ...uint) ([]any, error) {
	rootObject := &object{
		value:       nil,
		members:     make([]*object, 0),
		isPrimitive: false,
	}
	if len(fieldNumbers) == 0 {
		return nil, fmt.Errorf("not found field")
	}
	err := queryFields1(bin, rootObject, fieldNumbers...)
	if err != nil {
		return nil, err
	}
	members := getRootMembers(rootObject)
	results := make([]any, 0)
	for _, member := range members {
		if member.isPrimitive {
			results = append(results, member.value)
		}
	}
	return results, nil
}
func getRootMembers(root *object) []*object {
	var roots []*object
	for _, member := range root.members {
		if member != nil && member.isPrimitive {
			roots = append(roots, member)
		} else {
			roots = append(roots, getRootMembers(member)...)
		}
	}
	return roots
}
func queryFields1(bin []byte, obj *object, fieldNumbers ...uint) error {
	tmpObject := obj
	result, err := queryFields0(bin, fieldNumbers[0])
	if err != nil {
		return err
	}
	tmpObject.members = append(tmpObject.members, result...)
	if len(fieldNumbers) == 1 {
		for _, member := range tmpObject.members {
			member.isPrimitive = true
		}
		return nil
	}
	for _, member := range tmpObject.members {
		err := queryFields1(member.value.([]byte), member, fieldNumbers[1:]...)
		if err != nil {
			return err
		}
	}
	return nil
}
func queryFields0(bin []byte, fieldNumber uint) ([]*object, error) {
	query, err := newProtobufParser(bin).query(fieldNumber)
	if err != nil {
		return nil, err
	}
	result := make([]*object, 0)
	for _, q := range query {
		result = append(result, &object{value: q})
	}
	return result, nil
}
