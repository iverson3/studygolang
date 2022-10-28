// +build linux

package main

import (
	"fmt"
	"net"
	"os"
	"reflect"
	"sync"
	"syscall"

	"golang.org/x/sys/unix"
)

var epoller *epoll

func main() {
	listener, err := net.Listen("tcp", ":9876")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	epoller, err := NewPoll()
	if err != nil {
		panic(err)
	}

	go start()

	for {
		conn, err := listener.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				fmt.Errorf("accept temp err: %v", ne)
				continue
			}

			return
		}

		err = epoller.Add(conn)
		if err != nil {
			fmt.Errorf("failed to add connection, error: %v", err)
			conn.Close()
		}
	}
}

func start() {
	buf := make([]byte, 1024)
	for {
		conns, err := epoller.Wait(3)
		if err != nil {
			fmt.Errorf("failed to epoll wait, error: %v", err)
			continue
		}

		for _, conn := range conns {
			if conn == nil {
				continue
			}

			n, err := conn.Read(buf)
			if err != nil {
				err = epoller.Remove(conn)
				if err != nil {
					fmt.Errorf("failed to remove from epoll, error: %v", err)
				}
				conn.Close()
			}

			fmt.Printf("got msg from connection, conn: %v, msg: %s", conn, string(buf[:n]))

			resp := "i got it."
			_, err = conn.Write([]byte(resp))
			if err != nil {
				err = epoller.Remove(conn)
				if err != nil {
					fmt.Errorf("failed to remove from epoll, error: %v", err)
				}
				conn.Close()
			}
		}
	}
}

type epoll struct {
	fd int
	connections map[int]net.Conn
	mu *sync.RWMutex
}

func NewPoll() (*epoll, error) {
	epfd, err := syscall.EpollCreate1(0)
	if err != nil {
		return nil, err
	}

	return &epoll{
		fd:          epfd,
		connections: make(map[int]net.Conn),
		mu:          &sync.RWMutex{},
	}, nil
}

func (e *epoll) Add(conn net.Conn) error {
	fd := socketFD(conn)
	err := syscall.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, fd, &syscall.EpollEvent{
		Events: syscall.EPOLLIN | syscall.EPOLLHUP,
		Fd:     int32(fd),
	})
	if err != nil {
		return err
	}

	e.mu.Lock()
	defer e.mu.Unlock()
	e.connections[fd] = conn
	return nil
}

func (e *epoll) Remove(conn net.Conn) error {
	fd := socketFD(conn)
	e.mu.Lock()
	defer e.mu.Unlock()
	// 判断连接是否存在
	if _, ok := e.connections[fd]; !ok {
		return nil
	}

	err := syscall.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
	if err != nil {
		return err
	}

	delete(e.connections, fd)
	return nil
}

func (e *epoll) Wait(timeout int) ([]net.Conn, error) {
	// 每次最多只处理100个就绪事件
	events := make([]syscall.EpollEvent, 100)
retry:
	n, err := syscall.EpollWait(e.fd, events, timeout)
	if err != nil {
		if err == syscall.EINTR {
			goto retry
		}
		return nil, err
	}

	conns := make([]net.Conn, 0)
	e.mu.RLock()
	defer e.mu.RUnlock()
	for i := 0; i < n; i++ {
		conns = append(conns, e.connections[int(events[n].Fd)])
	}
	return conns, nil
}

func socketFD(conn net.Conn) int {
	//tls := reflect.TypeOf(conn.UnderlyingConn()) == reflect.TypeOf(&tls.Conn{})
	// Extract the file descriptor associated with the connection
	//connVal := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn").Elem()
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	//if tls {
	//	tcpConn = reflect.Indirect(tcpConn.Elem())
	//}
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")

	return int(pfdVal.FieldByName("Sysfd").Int())
}

func test() {
	epfd, err := syscall.EpollCreate1(0)
	if err != nil {
		panic(err)
	}

	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	syscall.Bind(fd, &syscall.SockaddrInet4{
		Port: 9090,
		Addr: [4]byte{},
	})
	//accept, sa, err := syscall.Accept(fd)

	syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, fd, &syscall.EpollEvent{Events: syscall.EPOLLIN | syscall.EPOLLHUP, Fd: int32(fd)})

	events := make([]syscall.EpollEvent, 0)
	n, err := syscall.EpollWait(epfd, events, 3)
	if err != nil {
		panic(err)
	}
	for i := 0; i < n; i++ {
		switch events[i].Events {
		case syscall.EPOLLIN:
			fmt.Println("got epoll in event: ", events[i].Fd)
		case syscall.EPOLLHUP:
			fmt.Println("got epoll hup event: ", events[i].Fd)
		default:
			fmt.Println("got epoll in event: ", events[i].Fd)
		}
	}


	unix.EpollCreate1(0)
	unix.EpollWait(epfd, nil, 10)


	syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, "")
	os.MkdirAll("rootfs/oldrootfs", 0700)
	syscall.PivotRoot("rootfs", "rootfs/oldrootfs")
	os.Chdir("/")
}