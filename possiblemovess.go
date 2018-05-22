package main
import "fmt"
func main() {
	
	var r,c,n int
	var p [5] int
	var  q [5] int
	var cp string
	var b [5] string
	s:= [5][5] string {{"a","b" , "c", "d", "e"},
	{"f", "g", "h", "i","j"},
	{"k", "l", "m", "n","o"},
	{"p", "q", "r", "s","t"},
	{ "u", "v", "w", "x","y"}}

	for _, e := range s {
		fmt.Println(e)
		}
	fmt.Println("Enter the current position")
	fmt.Scanf("%s", &cp)
	r,c=getindex(cp)

	fmt.Println("Enter the number of blocks1")
	fmt.Scanf("%d", &n)
        for i:=0;i <n ;i++ {
		fmt.Println("Enter the blocked position of ", i+1)
		fmt.Scanf("%s", &b[i])
		p[i],q[i]=getindex(b[i])
	}
		
	m:=0
	var list1,list2 []int
	if (r== 4){
		 list1 = []int{r-1,r}
	} else if ( r == 0){
		 list1 = []int{r,r+1}
	}else{
		 list1 = []int{r-1,r,r+1}
	}

	if (c== 4){
		 list2 = []int{c-1,c}
	} else if ( c == 0){
		 list2 = []int{c,c+1}
	}else{
		  list2 = []int{c-1,c,c+1}
	}
	
	for _, a := range list1{
		for _, b := range list2{
			flag:=0
			for v:= 0; v < n; v++{
				if((a == p[v]) && (b == q[v])){
					flag=1
				}else if((a == r) && (b == c)){
					flag=1
				}
		
			}
			if(flag==0){
				m++
			}
		}
				
		}
		
	fmt.Println(" Number of moves =",m)
				
	
	}

	

	func getindex (x string )(int, int) {
		a := [25]string {"a","b" , "c", "d", "e","f", "g", "h", "i","j","k", "l", "m", "n","o","p", "q", "r", "s","t", "u", "v", "w", "x","y"}
		var row,col int
		row=0
		col=0
	
		for _, r := range a {
				col++
				if(r==x){
					break
				}
				if (col % 5 == 0){
					row++
					col=0
				}
				
		
				
		}			
		return row,col-1
	}

