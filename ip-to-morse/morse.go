package main
import (
  "net"
  "bytes"
)

var mM = map[string]string{
  "0": "_ _ _ _ _",
  "1": ". _ _ _ _",
  "2": ". . _ _ _",
  "3": ". . . _ _",
  "4": ". . . . _",
  "5": ". . . . .",
  "6": "_ . . . .",
  "7": "_ _ . . .",
  "8": "_ _ _ . .",
  "9": "_ _ _ _ .",
  ".": "  ",
}


func main() {
  ln, _ := net.Listen("tcp", ":5555")
  for {
    conn, err := ln.Accept()
    if err != nil {
      panic(err)
    }
    go func() {
      clientIP, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
      var mIP = getMorse(clientIP)
      conn.Write([]byte(mIP + "\n"))
      conn.Close()
    }()
  }
}

func getMorse(ip string) string {
  var buffer bytes.Buffer
  for _, c := range ip {
    key := string(c)
    _, exists := mM[key]
    if exists {
      buffer.WriteString(mM[key])
      buffer.WriteString("\n")
    } else {
      buffer.WriteString("\n")
    }
  }
  return buffer.String()
}
