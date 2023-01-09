package main

import (
	"os"
	"fmt"
)

func openFile(filePath string) () {

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Cannot access file %s.\n", filePath)
		//return nil, err // error code?
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		//return nil, err // error code?
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		//return nil, err //error code?
	}

	fmt.Println("bytes read: ", bytesread)
	fmt.Println("bytestream to string: ", string(buffer))
	//return bytesread, nil // no error
}



// func openFile(filePath *string) {


// 	binaryFile, err := os.Open(*filePath)
// 	if err != nil {
// 		fmt.Printf("Cannot access telemetry file %s.\n", *filePath)
// 		os.Exit(1)
// 	}
// }