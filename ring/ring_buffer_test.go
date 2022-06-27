package queue

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func TestRing(t *testing.T) {
	Convey("Test Write", t, func() {
		Convey("Test write success", func() {
			q := NewRingBuffer(make([]byte, 1024))
			for i := 0; i < 1024; i++ {
				e := q.Write(make([]byte, 1))
				So(e, ShouldBeNil)
			}

			q = NewRingBuffer(make([]byte, 1024))
			e := q.Write(make([]byte, 1024))
			So(e, ShouldBeNil)

			q = NewRingBuffer(make([]byte, 1024))
			e = q.Write(make([]byte, 1023))
			So(e, ShouldBeNil)
		})

		Convey("Test write fail", func() {
			q := NewRingBuffer(make([]byte, 1024))
			e := q.Write(make([]byte, 1025))
			So(e, ShouldNotBeNil)

			q = NewRingBuffer(make([]byte, 1024))
			e = q.Write(make([]byte, 1000))
			So(e, ShouldBeNil)
			e = q.Write(make([]byte, 25))
			So(e, ShouldNotBeNil)
		})
	})

	Convey("Test Read", t, func() {
		Convey("Test read success", func() {
			q := NewRingBuffer(make([]byte, 1024))
			for i := 0; i < 1024; i++ {
				e := q.Write([]byte{byte(i)})
				So(e, ShouldBeNil)
			}

			for i := 0; i < 1024; i++ {
				b := make([]byte, 1)
				e := q.Read(b)
				So(e, ShouldBeNil)
				So(b, ShouldResemble, []byte{byte(i)})
			}

			// empty can't read
			e := q.Read(make([]byte, 1))
			So(e, ShouldNotBeNil)

			// write, read again, make pos > size
			for i := 0; i < 1024; i++ {
				e := q.Write([]byte{byte(i)})
				So(e, ShouldBeNil)
			}
			e = q.Write(make([]byte, 1))
			So(e, ShouldNotBeNil)

			for i := 0; i < 1024; i++ {
				b := make([]byte, 1)
				e := q.Read(b)
				So(e, ShouldBeNil)
				So(b, ShouldResemble, []byte{byte(i)})
			}

		})

		Convey("test write pos before read", func() {
			q := NewRingBuffer(make([]byte, 4))
			q.Write([]byte{1, 2, 3, 4})
			q.Read(make([]byte, 2))
			q.Write([]byte{5, 6})
			b := make([]byte, 4)
			e := q.Read(b)
			So(e, ShouldBeNil)
			So(b, ShouldResemble, []byte{3, 4, 5, 6})
		})

		Convey("test pos over max uint32", func() {
			q := NewRingBuffer(make([]byte, 1024))
			q.in = math.MaxUint32 - 10
			q.out = q.in

			for i := 0; i < 1024; i++ {
				e := q.Write([]byte{byte(i)})
				So(e, ShouldBeNil)
			}
			e := q.Write(make([]byte, 1))
			So(e, ShouldNotBeNil)

			for i := 0; i < 1024; i++ {
				b := make([]byte, 1)
				e := q.Read(b)
				So(e, ShouldBeNil)
				So(b, ShouldResemble, []byte{byte(i)})
			}

		})

		Convey("Test read fail", func() {
			q := NewRingBuffer(make([]byte, 1024))
			e := q.Read(make([]byte, 1))
			So(e, ShouldNotBeNil)

			q.Write(make([]byte, 10))
			e = q.Read(make([]byte, 11))
			So(e, ShouldNotBeNil)
		})

		Convey("Test peek", func() {
			q := NewRingBuffer(make([]byte, 1024))
			for i := 0; i < 1024; i++ {
				e := q.Write([]byte{byte(i)})
				So(e, ShouldBeNil)
			}

			for i := 0; i < 1024; i++ {
				b := make([]byte, 1)
				e := q.Peek(b)
				So(e, ShouldBeNil)
				So(b, ShouldResemble, []byte{byte(i)})

				e = q.Read(b)
				So(e, ShouldBeNil)
				So(b, ShouldResemble, []byte{byte(i)})
			}

		})
	})
}
