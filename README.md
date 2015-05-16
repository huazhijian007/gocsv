# gocsv
read cvs file for golangfile,it's so easy

csv file content example:

idx,name,age,sex

comment,it is name,it is age,it is sex

1,lantin,28,1

2,momo,27,0

3,james,2,1

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
