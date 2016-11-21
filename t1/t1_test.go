package main

import (
	"errors"
	"fmt"
	"net"
	"testing"
	"time"
)

//Struct for Conn type
type myConn struct {
	buf  []byte
	r, w int
}

func (b *myConn) Read(p []byte) (n int, err error) {
	//check if r=w means buffer is empty
	if b.r == b.w {
		errReadEmpty := errors.New("errReadEmpty")
		return 0, errReadEmpty
	}
	//copy the buffer value to p
	n = copy(p, b.buf[b.r:b.w])
	b.r += n
	if b.r == b.w {
		b.r = 0
		b.w = 0
	}
	return n, nil
}

//Function for close connection
func (c *myConn) Close() error {

	return nil
}

func (b *myConn) Write(p []byte) (n int, err error) {
	// Slide existing data to beginning.
	if b.r > 0 && len(p) > len(b.buf)-b.w {
		copy(b.buf, b.buf[b.r:b.w])
		b.w -= b.r
		b.r = 0
	}

	// Write new data.
	n = copy(b.buf[b.w:], p)
	b.w += n
	if n < len(p) {
		err = errors.New("errWriteFull")
	}
	return n, err
}

func (b *myConn) LocalAddr() net.Addr {
	//some code
	return nil
}
func (b *myConn) RemoteAddr() net.Addr {
	// some code
	return nil
}
func (b *myConn) SetDeadline(t time.Time) error {
	// some code
	return nil
}
func (b *myConn) SetReadDeadline(t time.Time) error {
	// some code
	return nil
}
func (b *myConn) SetWriteDeadline(t time.Time) error {
	// some code
	return nil
}

//Testing for Read() Function
func TestRead(t *testing.T) {

	b := make([]byte, 4)
	b[0] = 0x00
	b[1] = 0x01
	b[2] = 10
	c := &myConn{buf: b, r: 0, w: 4}

	//expectederr := errors.New("errReadEmpty")
	rCount, err1 := c.Read(b)
	if err1 != nil && rCount != 4 {
		t.Fatal("failed to read ")
	}
	//fmt.Println("read count is : \n", rCount)
	fmt.Println(b)
}

//Testing of Write() Function
func TestWrite(t *testing.T) {

	b := make([]byte, 4)
	b[0] = 0x00
	b[1] = 0x01
	b[2] = 10
	c := &myConn{buf: b, r: 0, w: 4}

	rCount, err1 := c.Read(b)
	if err1 != nil && rCount != 4 {
		t.Fatal("failed to read ")
	}
	fmt.Println(b)

	wCount, err := c.Write(b)
	if err != nil && wCount != 4 {
		t.Fatal("failed to write ")
	}
	fmt.Println(c.buf)
}

//Testing for Check function (gives reponse to a message)
func TestCheck(t *testing.T) {

	expected := "Hello"
	actual := "Hi"
	msg := Check(expected)
	fmt.Println(msg)
	if msg != actual {
		t.Fatal("failed to Interpret message ")
	}

	/*Expected := ""
	Actual := "Hello"
	msg = Check(Expected)
	fmt.Println(msg)
	if msg != Actual {
		t.Fatal("failed to Interpret message ")
	}*/
}

//testing Request function which handle the connection
func TestRequest(t *testing.T) {

	b := make([]byte, 20)

	c := &myConn{buf: b, r: 0, w: 20}

	handle := Request(c, b)
	if handle != nil {
		t.Fatal("failed to Handle Operation ")
	}

	/*handle := handleRequest(c)
	if handle != nil {
		t.Fatal("failed to Handle Operation ")
	}*/
}

func TestRequestFailure(t *testing.T) {

	b := make([]byte, 20)

	c := &myConn{buf: b, r: 0, w: 0}

	handle := Request(c, b)
	if handle == nil {
		t.Fatal("Failure")
	}

}
