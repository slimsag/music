package smf

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestVariableLengthQuantity(t *testing.T) {
	tests := []struct {
		input uint32
		want  []byte
	}{
		{
			input: 0x00000000,
			want:  []byte{0x00},
		},
		{
			input: 0x00000040,
			want:  []byte{0x40},
		},
		{
			input: 0x0000007F,
			want:  []byte{0x7F},
		},
		{
			input: 0x00000080,
			want:  []byte{0x81, 0x00},
		},
		{
			input: 0x00002000,
			want:  []byte{0xC0, 0x00},
		},
		{
			input: 0x00003FFF,
			want:  []byte{0xFF, 0x7F},
		},
		{
			input: 0x00004000,
			want:  []byte{0x81, 0x80, 0x00},
		},
		{
			input: 0x00100000,
			want:  []byte{0xC0, 0x80, 0x00},
		},
		{
			input: 0x001FFFFF,
			want:  []byte{0xFF, 0xFF, 0x7F},
		},
		{
			input: 0x00200000,
			want:  []byte{0x81, 0x80, 0x80, 0x00},
		},
		{
			input: 0x08000000,
			want:  []byte{0xC0, 0x80, 0x80, 0x00},
		},
		{
			input: 0x0FFFFFFF,
			want:  []byte{0xFF, 0xFF, 0xFF, 0x7F},
		},
	}
	for _, tst := range tests {
		t.Run(fmt.Sprintf("%X", tst.input), func(t *testing.T) {
			buf := make([]byte, binary.MaxVarintLen32)
			n := binary.PutUvarint(buf, uint64(tst.input))
			got := buf[:n]
			if !bytes.Equal(got, tst.want) {
				t.Logf("got %v want %v; input=%v", got, tst.want, tst.input)
			}
		})
	}
}
