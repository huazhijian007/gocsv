//author:lantin_fang@163.com 2014-12-30
package gocsv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	ErrRowNotExist = errors.New("rows not exist")
	ErrReadCsv     = errors.New("read csv error")
)

type Row struct {
	MapCellsList map[string]string
}

func (this *Row) GetInteger(field string) int {

	value, isok := this.MapCellsList[field]
	if !isok {
		fmt.Printf("the field :%s is not exist\n", field)
		return int(-1)
	}

	r, _ := strconv.Atoi(value)
	return r
}

func (this *Row) GetString(field string) string {

	value, isok := this.MapCellsList[field]
	if !isok {
		fmt.Printf("the field :%s is not exist\n", field)
		return "nil"
	}

	return value
}

type Csv struct {
	MapRowsList map[int]*Row
	Feilds      []string
	reader      *csv.Reader
}

func NewCsv() *Csv {
	return &Csv{
		MapRowsList: make(map[int]*Row),
		reader:      nil,
	}
}

func (this *Csv) LoadFromFile(filename string) error {

	f, err := os.OpenFile(filename, os.O_RDONLY, 0444)
	if err != nil {
		fmt.Printf("open csv file :%s error,:%s\n", filename, err.Error())
		return ErrReadCsv
	}

	this.reader = csv.NewReader(f)

	// read fields
	fields, err := this.reader.Read()
	for _, v := range fields {
		this.Feilds = append(this.Feilds, v)
	}

	// read rows
	for {
		line, err := this.reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		// it is comment line
		if line[0] == "comment" { // if you do not want this issues,please remove this code
			continue
		}

		row := new(Row)
		row.MapCellsList = make(map[string]string)

		for i, v := range line {
			row.MapCellsList[fields[i]] = v
		}

		key, _ := strconv.Atoi(line[0])
		this.MapRowsList[key] = row
	}

	fmt.Printf("load csv file success:%s\n", filename)
	return nil
}

func (this *Csv) FindRows(idx int) (*Row, error) {
	row, isok := this.MapRowsList[idx]
	if !isok {
		return nil, ErrRowNotExist
	}

	return row, nil
}

/*
--csv file content example

idx,name,age,sex
comment,it is name,it is age,it is sex
1,lantin,28,1
2,momo,27,0
3,james,2,1

---------------------
Attention:
1.the first column must be index ,it's integer
2.the second can be comment,you can comment somthing for this fileld


example for use gocsv :

func main() {
	c := NewCsv()
	e := c.LoadFromFile("./test.csv")
	if e != nil {
		fmt.Printf("load error\n")
		return
	}

	row, _ := c.FindRows(1)
	name := row.GetString("name")
	fmt.Printf("name:%s\n", name)
	return
}
*/
