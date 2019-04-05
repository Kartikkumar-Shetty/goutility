package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	ReadFromFileUsingReadLn()

}


func ReadFromStdIN(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}


func ReadFromFile(){
	file, err := os.Open("./file.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}


func ReadUsingReadLine(r *bufio.Reader) (string, error) {
	var (isPrefix bool = true
		 err error = nil
		 line, ln []byte
		)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln),err
  }


  func ReadFromFileUsingReadLn(){
	fi:="./file.txt"
	f, err := os.Open(fi)
	if err != nil {
		fmt.Printf("error opening file: %v\n",err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, e := ReadUsingReadLine(r)
	for e == nil {
		fmt.Println(s)
		s,e = ReadUsingReadLine(r)
	}
  }