package pizza

import (
	"os"
	"github.com/sirupsen/logrus"
	"bufio"
	"strings"
	"strconv"
)

func check(e error) {
	if e != nil {
		logrus.Error(e.Error())
		panic(e)
	}
}

type Pizza struct{
	R int8 //number of row
	C int8 //number of column
	L int8 //minimun number of each cell in slice
	H int8 //max slice cell size
	Cells []Cell
}

func New(path string) (p Pizza,err error) {
	inFile, err := os.Open(path)
	if err!=nil {
		return p, err
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	lineindex := 0
	for scanner.Scan() {
		if lineindex==0 {
			line := strings.Split(scanner.Text(), " ")
			i, err := strconv.ParseInt(line[0],10,8)
			if err != nil {
				return p, err
			}
			p.R = int8(i)
			i, err = strconv.ParseInt(line[1],10,8)
			if err != nil {
				return p, err
			}
			p.C = int8(i)
			i, err = strconv.ParseInt(line[2],10,8)
			if err != nil {
				return p, err
			}
			p.L = int8(i)
			i, err = strconv.ParseInt(line[3],10,8)
			if err != nil {
				return p, err
			}
			p.H = int8(i)
			lineindex+=1
			continue
		}
		line := scanner.Text()
		for index:=0; index< len(line); index++ {
			c:=Cell{
				x: lineindex-1,
				y: index,
				taken:false,
			}
			if string(line[index]) == "T"{
				c.val=false
			}else{
				c.val=true
			}

			p.Cells = append(p.Cells,c)
		}
		lineindex+=1
	}
	return p,nil
}

type Cell struct {
	x int
	y int
	taken bool
	val bool //T=false M=true
}


