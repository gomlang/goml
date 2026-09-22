package native

import (
	"encoding/binary"
	"math"
)

func BinaryRecordBytes(enabled bool, sequence uint16, samples []int16, weightBits uint32, little bool) []byte {
	var order binary.ByteOrder = binary.BigEndian
	if little {
		order = binary.LittleEndian
	}
	value := struct {
		Enabled  bool
		Sequence uint16
		Samples  [3]int16
		Weight   float32
		_        [3]byte
	}{Enabled: enabled, Sequence: sequence, Samples: [3]int16{samples[0], samples[1], samples[2]}, Weight: math.Float32frombits(weightBits)}
	output, err := binary.Append(nil, order, value)
	if err != nil {
		panic(err)
	}
	return output
}
