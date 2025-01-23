package main

import (
	// "bufio"
	"fmt"
	"os"
)

func main() {
	// Create a file
	// fp,e := os.Open("file.txt")
	// if (e != nil) {
	// 	panic(e)
	// } 
	// fileInfo, err := fp.Stat()
	// if (err != nil) {
	// 	panic(err)
	// }
	// 	fmt.Println("File size:", fileInfo.Size())
	// 	fmt.Println("File name:", fileInfo.Name())
	// 	fmt.Println("File mode:", fileInfo.Mode())
	// 	fmt.Println("File mod time:", fileInfo.ModTime())
	// 	fmt.Println("File is directory:", fileInfo.IsDir())
	// defer fp.Close()


	// read file 1st way
	// fp,err := os.Open("file.txt")
	// if (err != nil) {
	// 	panic(err)
	// }
	// defer fp.Close()
	// buf := make([]byte, 16)
	// num, err := fp.Read(buf)
	// fmt.Println("Number of bytes read:", num)
	// for i := 0; i < len(buf); i++ {
	// 	fmt.Print(string(buf[i]))
	// }

	// read file 2nd way
	// buf, err := os.ReadFile("file.txt")
	// if (err != nil) {
	// 	panic(err)
	// }
	// fmt.Println(string(buf))

	// opening and reading a dir
	// dir, err := os.Open("..")
	// if (err != nil) {
	// 	panic(err)
	// }
	// info , err := dir.Stat()
	// if (err != nil) {
	// 	panic(err)
	// }
	// fmt.Println("Is directory:", info.IsDir())
	// fmt.Println("Size:", info.Size())
	// defer dir.Close()

	// entries, err := dir.ReadDir(-1)
	// if (err != nil) {
	// 	panic(err)
	// }
	// for _, entry := range entries {
	// 	fmt.Println("Name:", entry.Name())
	// 	fmt.Println("Is directory:", entry.IsDir())
	// }

	// creating a file
	// fp, err := os.Create("file2.txt")
	// if (err != nil) {
	// 	panic(err)
	// }
	// defer fp.Close()

	// writing to a file 1st way
	// fp.WriteString("Hello, World!\n")
	// fp.WriteString("This is a file created using Go!\n")

	// writing to a file 2nd way
	// buf := []byte("Hello, World!\n")
	// fp.Write(buf)

	// info, err := fp.Stat()
	// if (err != nil) {
	// 	panic(err)
	// }
	// fmt.Println("File name:", info.Name())
	// fmt.Println("File size:", info.Size())

	
	// read and write to a file (streaming fashion)
	// srcFile, err := os.Open("file.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer srcFile.Close()
	// info1, err := srcFile.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Source file name:", info1.Name())
	// fmt.Println("Source file size:", info1.Size())

	// destFile, err := os.Create("file2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer destFile.Close()
	// info2, err := destFile.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Source file name:", info2.Name())
	// fmt.Println("Source file size:", info2.Size())

	// reader := bufio.NewReader(srcFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	buf, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	e := writer.WriteByte(buf)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }

	// writer.Flush()

	// fmt.Println("File copied successfully!")

	// delete a file
	e := os.Remove("file2.txt")
	if e != nil {
		panic(e)
	}
	fmt.Println("File deleted successfully!")
}