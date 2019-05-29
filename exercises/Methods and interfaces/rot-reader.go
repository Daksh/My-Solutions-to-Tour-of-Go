package main

import (
	"io"
	"os"
	"strings"
	//"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (source rot13Reader) Read(b []byte) (int, error) {
	in := make([]byte, 100)
	num, err := source.r.Read(in)
	if err != nil {
		//fmt.Println(err)
		return 0, err
	}
	//fmt.Printf("len: %d, cap: %d", len(in),cap(in))//I get len 100 and cap 10076, how can cap increase?!
	for i := 0; i < num; i++ {
		//fmt.Printf("%d, %s\n", in[i],string(in[i]))
		if 97 <= in[i] && in[i] <= 122 {
			b[i] = (in[i]-97+13)%26 + 97
		} else if 65 <= in[i] && in[i] <= 90 {
			b[i] = (in[i]-65+13)%26 + 65
		} else {
			b[i] = in[i]
		}
	}
	return num, nil
}

func main() {
	//fmt.Println("%d %d %d %d", 'a','z','A','Z')
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
