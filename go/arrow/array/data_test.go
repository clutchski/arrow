package array

import (
	"testing"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/memory"
	"github.com/stretchr/testify/assert"
)

func TestDataBuffers(t *testing.T) {
	assert := assert.New(t)

	// create LSB packed bits with the following pattern:
	// 01010011 11000101
	data := memory.NewBufferBytes([]byte{0xca, 0xa3})

	// create LSB packed validity (null) bitmap, where every 4th element is null:
	// 11101110 11101110
	nullBitmap := memory.NewBufferBytes([]byte{0x77, 0x77})

	d := NewData(arrow.FixedWidthTypes.Boolean, 16, []*memory.Buffer{nullBitmap, data}, nil, UnknownNullCount, 0)

	buffers := d.Buffers()
	assert.Equal(2, len(buffers))
	assert.Equal(buffers[0], nullBitmap)
	assert.Equal(buffers[1], data)
}
