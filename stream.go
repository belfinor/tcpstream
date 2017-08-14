package main


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-08-14


type Stream struct {
  Connects []*Conn
}


func MakeStream( addrs []string ) {
  str := &Stream{ 
    Connects: make( []*Conn, len(addrs) ),
  }

  for i, addr := range addrs {
    str.Connects[i] = NewConn( addr )
  }
}


func (s *Stream) Send( msg []byte ) {
  for _, c := range s.Connects {
    c.Lock()
    c.Send( msg )
    c.Unlock()
  }
}

