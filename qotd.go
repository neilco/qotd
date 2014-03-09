package main

// qotd.txt source: ftp://ftp.mrynet.com/USENIX/85.1/langston/sun/TODAY/qotd.txt

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"time"
)

type Quotes map[string]string

func (file *Quotes) Parse(qfile string) {
	bytes, err := ioutil.ReadFile(qfile)
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	content := strings.Replace(string(bytes), "}\n", "}", -1)
	fields := strings.FieldsFunc(content, brace)

	for i, field := range fields {
		if len(field) == 4 {
			(*file)[field] = fields[i+1]
		}
	}
}

func brace(r rune) bool {
	return r == '{' || r == '}'
}

// --------------------------------------------------------------------------------

func qotd(conn net.Conn) {
	today := time.Now().Format("0102")
	buf := []byte(quotes[today] + "\n")

	// RFC 865 (https://tools.ietf.org/html/rfc865) states that the quote should
	// be less than 512 characters
	if len(buf) > 512 {
		buf = buf[:512]
	}

	_, err := conn.Write(buf)
	if err != nil {
		log.Println("Error send reply:", err.Error())
	}

	defer conn.Close()
}

// --------------------------------------------------------------------------------

var quotes Quotes
var qfile = flag.String("file", "qotd.txt", "The QOTD file")

func main() {
	flag.Parse()

	quotes = make(Quotes)
	quotes.Parse(*qfile)

	listener, err := net.Listen("tcp", ":17")
	if err != nil {
		log.Fatal("Error listening:", err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accept:", err.Error())
			continue
		}
		go qotd(conn)
	}
}
