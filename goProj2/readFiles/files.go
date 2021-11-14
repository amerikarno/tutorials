package readFiles

import (
	"bufio"
	"time"
	//"fmt"
	"log"
	"os"
	"strings"
)

func Logs(fname string) *[]string{
	fData, err := os.Open(fname)
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

	//fmt.Println(data)
	
	return &data
}

func Nline(file []string, n int) *[]string{
	var data []string
	for i, j := range file {
		data = append(data,j)
		if i >= n {
			break
		}
	}
	return &data
}

func GetTime(c string,n *[]string) *time.Time{
	var h []string
	for _,k := range *n {
		if strings.Contains(k,c) {
			h = strings.Split(k," ")
			
		}
//		fmt.Println(k)
	}
	i,_ := time.Parse("15:04:05.000",h[0])
	return &i
}