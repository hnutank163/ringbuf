package queue

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestQueue(t *testing.T) {
    Convey("test queue", t, func() {
        Convey("test queue put", func() {
            q := NewQueue()

            e := q.Put([]byte{1,2,3})
            So(e, ShouldBeNil)
            e = q.Put([]byte{4,5,6})
            So(e, ShouldBeNil)

            e = q.Put(make([]byte, 1024))
            So(e, ShouldNotBeNil)
        })

        Convey("test queue pop", func() {
            q := NewQueue()
            e := q.Put([]byte{1,2,3})
            e = q.Put([]byte{4,5,6})

            b, e := q.Pop()
            So(e, ShouldBeNil)
            So(b, ShouldResemble, []byte{1,2,3})
            b, e = q.Pop()
            So(b, ShouldResemble, []byte{4,5,6})
        })
    })
}
