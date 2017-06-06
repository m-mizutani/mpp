package main

import (
	"github.com/kr/pretty"
	"github.com/ugorji/go/codec"
	"io"
	"os"
)

func main() {
	mh := &codec.MsgpackHandle{RawToString: true}
	dec := codec.NewDecoder(os.Stdin, mh)

	for {
		var msg interface{}
		err := dec.Decode(&msg)
		if err == nil {
			pretty.Println(msg)
		} else if err == io.EOF {
			break
		} else {
			panic(err)
		}
	}
}
