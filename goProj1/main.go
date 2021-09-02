package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	//"regexp"
	//"strconv"
)

func main() {
	/*fName := "/Users/amerikarno/DataTransferFromRD/log.txt"
	r := *rfile(fName)
	for x, line := range r{
	if x > 4{
		break
	}
	//fmt.Println(line)
	/*
	reg := regexp.MustCompile(`(\d+)=(\w+)`)
	//re := reg.FindAllString(line[:len(line)-1],-1)
	//fmt.Printf("%#v\n",reg.FindAllString(line,-1))
	for _,str := range reg.FindAllString(line,-1){
		//s := strings.Split(str,"=")
		//s := strings.Replace(str,"35","Msgid",-1)
		//fmt.Printf("%#v\n",s)
	
	//fmt.Printf("%#v\n",reg.SubexpNames())
	}
	*/
	t := *rtag42("fUsedTag.txt")
	fmt.Printf("%q",t["35"])
	//var tag map[string]string
	
}
//}	


func rfile(fName string) *[]string{

	fData, err := os.Open(fName)
	if err != nil {
		log.Fatal(err)
	}
	defer fData.Close()

	scanner := bufio.NewScanner(fData)
	scanner.Split(bufio.ScanLines)

	var data []string
	for scanner.Scan(){
		data = append(data, scanner.Text())
	}

	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	
	return &data
}

func rtag42(tName string) *map[string]string{
	tData, err := os.Open(tName)
	if err != nil {
		log.Fatal(err)
	}
	defer tData.Close()

	scan := bufio.NewScanner(tData)
	scan.Split(bufio.ScanLines)
	elementMap := make(map[string]string)
	for scan.Scan(){
		s := strings.Split(scan.Text(),",")
		s1, s2 := s[0],s[1]
		elementMap[s1] = s2		
	}
	return &elementMap
}
/*
func replacer(s string) *string{
	r := rtag42("fUsedTag.txt")
	
	if m {
		return &i
	} else {
		return &m
	}
}*/