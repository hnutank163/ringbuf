package queue

import (
    "encoding/binary"
    "fmt"
)

type Queue struct {
    r *RingBuffer
}

type QueueItem struct {
     
}

func NewQueue() *Queue {
    q := &Queue{
    }

    q.r = NewRingBuffer(make([]byte, 1024))

    return q
}

func (q *Queue) Put(b []byte) error {
    buf := make([]byte, 2, 2+len(b))
    binary.LittleEndian.PutUint16(buf, uint16(len(b)))
    return q.r.Write(append(buf, b...))
}

func (q *Queue) Pop() ([]byte, error){
    buf := make([]byte, 2)
    err := q.r.Peek(buf)
    if err == nil {
        l := binary.LittleEndian.Uint16(buf)
        if uint32(l + 2) <= q.r.Size() {
            b := make([]byte, l+2)
            q.r.Read(b)
            return b[2:], nil
        }

        return nil, fmt.Errorf("no enough data")
    }

    return nil, err
}

func foo() {
    // b := bufio.NewReadWriter().Peek
}
