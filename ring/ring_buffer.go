package queue

import (
	"fmt"
	// "log"
)

type RingBuffer struct {
	buffer []byte
	size   uint32
	mask   uint32
	in     uint32
	out    uint32
	// mu     sync.Mutex
}

func NewRingBuffer(b []byte) *RingBuffer {
	r := &RingBuffer{
		buffer: b,
		size:   uint32(len(b)),
		in:     0,
		out:    0,
	}
	r.mask = r.size - 1
	r.mask = r.size - 1

	return r
}

func Min(a ...uint32) uint32 {
	if len(a) == 0 {
		return 0

	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v

		}

	}
	return min

}

func (r *RingBuffer) Write(b []byte) error {
	left := r.size - (r.in - r.out)
	if uint32(len(b)) > left {
		return fmt.Errorf("no enough space")
	}

	l := Min(uint32(len(b)), r.size-r.in&r.mask)
	copy(r.buffer[r.in&r.mask:], b[:l])
	copy(r.buffer, b[l:])
	r.in += uint32(len(b))
	return nil
}

func (r *RingBuffer) Read(b []byte) error {
	// wantRead := uint32(len(b))
	// canRead := r.in - r.out
	// if wantRead > canRead {
	// 	return fmt.Errorf("no enough data")
	// }

	// outOff := r.out & r.mask
	// readToEnd := Min(wantRead, r.size-outOff)
	// // log.Printf("size %d, out %d, readToEnd %d", r.size, r.out, readToEnd)
	// copy(b, r.buffer[outOff:outOff+readToEnd])
	// copy(b[readToEnd:], r.buffer[:wantRead-readToEnd])
	// r.out += uint32(len(b))
	// return nil
    e := r.Peek(b)
    if e == nil {
        r.out += uint32(len(b))
    }

    return e
}

func (r *RingBuffer) Size() uint32 {
    return r.in - r.out
}

func (r *RingBuffer) Peek(b []byte) error {
	wantRead := uint32(len(b))
	canRead := r.in - r.out
	if wantRead > canRead {
		return fmt.Errorf("no enough data")
	}

	outOff := r.out & r.mask
	readToEnd := Min(wantRead, r.size-outOff)
	// log.Printf("size %d, out %d, readToEnd %d", r.size, r.out, readToEnd)
	copy(b, r.buffer[outOff:outOff+readToEnd])
	copy(b[readToEnd:], r.buffer[:wantRead-readToEnd])
	return nil

}
