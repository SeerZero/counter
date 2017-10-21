package main

import (
	"bufio"
	"os"
	"fmt"
	"io"
	"strings"
	"strconv"
)


func main() {
	inputFile, inputError := os.Open("millionRandom.txt")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	m:=make(map[string]int64)
	for {
		s,error:=inputReader.ReadString('\n')
		if error==io.EOF{
			break
		}
		m[s]++
	}


	outputFile, outputError := os.OpenFile("counter.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("File open failed.")
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	for i:=range m {//i 本身带了个回车Orz
		str:=i
		str=strings.Replace(str,"\n","",-1)
		num:= strconv.FormatInt(m[i], 10)
		outputWriter.WriteString(str+"   "+ num+"\n")
	}

}
