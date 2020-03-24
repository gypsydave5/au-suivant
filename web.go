package suivant

import (
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"html/template"
	"net"
	"net/http"
)

type Server struct {
	s *Suivant
	http.Handler
}

func NewServer(s *Suivant) *Server {
	server := new(Server)
	server.s = s

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(server.indexHandler))
	mux.Handle("/ws", http.HandlerFunc(server.wsHandler))

	server.Handler = mux
	return server
}

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.template.go.html")
	_ = t.Execute(w, nil)
}

func (s *Server) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _, _, _ := ws.UpgradeHTTP(r, w)
	nextChan := s.s.Next()
	wsconn := NewWSConn(conn)

	for {
		select {
		case command := <-wsconn.Receive:
			switch command {
			case "start":
				s.s.Start()
			case "stop":
				s.s.Stop()
			}
		case n := <-nextChan:
			wsconn.Send <- n
		}
	}
}

type WSConn struct {
	c       net.Conn
	w       *wsutil.ControlWriter
	Send    chan<- string
	Receive <-chan string
}

func NewWSConn(conn net.Conn) (wsconn *WSConn) {
	send := make(chan string)
	receive := make(chan string)
	writer := wsutil.NewControlWriter(conn, ws.StateServerSide, ws.OpText)

	wsc := &WSConn{
		c:       conn,
		w:       writer,
		Send:    send,
		Receive: receive,
	}

	go func(c chan<- string) {
		for {
			payload, _ := wsutil.ReadClientText(conn)
			c <- string(payload)
		}
	}(receive)

	go func(c <-chan string) {
		for {
			msg := <-c
			wsc.w.Write([]byte(msg))
			wsc.w.Flush()
		}
	}(send)

	return wsc
}
