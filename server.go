package main


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-08-13


import (
    "fmt"
    "github.com/belfinor/Helium/log"
    "net"
    "time"
)


type CALLBACK func([]byte) ([]byte, error)


type Server struct {
    Addr         string
}


var nextId int64 = 0;


func (s *Server) Start() {

    ln, err := net.Listen( "tcp", s.Addr )
	
    if err != nil {
        panic( "bind port error" )
    }

    log.Info( "server start " + s.Addr )

    for {
        
        conn, err := ln.Accept() 

        if err != nil {
            continue
        }
			
        nextId++

        go s.connet_handler(conn, nextId )
    }
}


func (s *Server) connet_handler(conn net.Conn, id int64 ) {
    defer conn.Close()

    log.Info( fmt.Sprintf( "new conection #%d", id ) )

    buffer := make( []byte, 40960 )
    
    for {
        conn.SetReadDeadline( time.Now().Add( time.Minute ) )
        n, err := conn.Read( buffer )
        if err != nil || n < 1 {
          log.Info( fmt.Sprintf( "connection #%d broken", id ) )
          break;
        }

        //list := buffer[:n]
        
    }

    log.Info( fmt.Sprintf( "connection #%d closed", id ) )
}

