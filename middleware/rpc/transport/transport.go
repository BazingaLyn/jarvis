package transport

import (
	"io"
	"net"
	"time"
)

type Transport interface {
	Dial(network, addr string, option DialOption) error
	io.ReadWriteCloser
	RemoteAddr() net.Addr
	LocalAddr() net.Addr
}

type Socket struct {
	conn net.Conn
}

func (s *Socket) Dial(network, addr string, option DialOption) error {
	var dialer net.Dialer
	if option.Timeout > time.Duration(0) {
		dialer.Timeout = option.Timeout
	}
	conn, err := dialer.Dial(network, addr)

	s.conn = conn
	return err
}

func (s *Socket) Read(p []byte) (n int, err error) {
	return s.conn.Read(p)
}

func (s *Socket) Write(p []byte) (n int, err error) {
	return s.conn.Write(p)
}

func (s *Socket) Close() error {
	return s.conn.Close()
}

func (s *Socket) RemoteAddr() net.Addr {
	return s.conn.RemoteAddr()
}

func (s *Socket) LocalAddr() net.Addr {
	return s.conn.LocalAddr()
}

type ServerTransport interface {
	Listen(network, addr string) error
	Accept() (Transport, error)
	io.Closer
}

type ServerSocket struct {
	ln net.Listener
}

func (s *ServerSocket) Listen(network, addr string) error {
	ln, err := net.Listen(network, addr)
	s.ln = ln
	return err
}

func (s *ServerSocket) Accept() (Transport, error) {
	conn, e := s.ln.Accept()
	return &Socket{conn: conn}, e
}

func (s *ServerSocket) Close() error {
	return s.ln.Close()
}

type DialOption struct {
	Timeout time.Duration
}
