package format

import (
	"bytes"
	"fmt"
)

func ExampleChunkDATA_SetData() {
	data := []byte("hello")
	s := new(ChunkDATA)
	s.SetData(data)
	fmt.Printf("%s\n", s.Data)
	fmt.Printf("%d\n", s.Length)
	fmt.Printf("%d\n", len(s.Padding))
	//Output:
	//hello
	//13
	//3
}

func ExampleChunkDATA_WriteTo() {
	data := []byte("hello")
	s := new(ChunkDATA)
	s.SetData(data)
	buf := new(bytes.Buffer)
	s.ToChunkField().WriteTo(buf)
	fmt.Printf("%v\n", buf.Bytes())
	//Output:
	//[0 0 0 13 0 0 0 0 0 0 0 0 104 101 108 108 111 0 0 0]
}
