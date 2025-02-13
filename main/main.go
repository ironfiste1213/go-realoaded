package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"mimon/helperfunction"
)


func main() {

	if len(os.Args) != 3 {
		fmt.Println("use thi form run input output: file1.txt file2.txt")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}
	scanner := bufio.NewScanner(file)
	if filepath.Ext(os.Args[1]) != ".txt"  || filepath.Ext(os.Args[2]) != ".txt"  {
		fmt.Println("use .txt and file1 != file2" )
		return
	}
	file, err = os.Create(os.Args[2])
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	for scanner.Scan() {
		
		fmt.Fprintln(file, format.All(scanner.Text()))   
		
	}
}
