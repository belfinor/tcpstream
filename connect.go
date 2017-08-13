package main


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-08-13


import (
  "github.com/belfinor/Helium/log"
  "net"
  "sync"
  "time"
)


type Conn struct {
  addr      string
  last_time int64
  last_try  int64
  conn      net.Conn
  buffer    []byte
  sync.Mutex
}


func NewConn( addr string ) *Conn {
  n := &Conn{
    addr:      addr,
    last_time: 0,
    last_try:  0,
    buffer:    make( []byte, 40960 ),
  }

  return n
}


func (n *Conn) keepAlive() bool {

  if time.Now().Unix() - n.last_time > 30 && n.conn != nil {
    log.Info( "close connect to " + n.addr +  " by timeout" ) 
    n.conn.Close()
  }

  var err error

  if n.conn == nil && time.Now().Unix() - n.last_try > 5 {
    log.Info( "try connect " + n.addr )
    n.last_try = time.Now().Unix()
    n.last_time = n.last_try
    n.conn, err = net.Dial( "tcp", n.addr )
    if err != nil {
      log.Info( "connect to " + n.addr + " failed" )
      n.conn = nil
    }
  }

  return n.conn != nil
}


func (n *Conn) Send( msg []byte ) bool {

  for i := 0 ; i < 2 ; i++ {
    if !n.keepAlive() {
      continue
    }

    wt, err := n.conn.Write( msg )
    if err == nil && wt == len(msg) {
      n.last_time = time.Now().Unix()
      return true
    }

    if wt < len(msg) {
      log.Debug( "output buffer full" )
    }

    n.conn.Close()
    n.conn = nil
  }
  
  return false
}


func (n *Conn) Read() []byte {

  if n.conn == nil {
    return nil
  }

  num, err := n.conn.Read( n.buffer )
  if err != nil {
    n.conn.Close()
    n.conn = nil
    return nil
  }

  return n.buffer[:num]
}

