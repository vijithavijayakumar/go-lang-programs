package main
import "fmt"
func main() {
	var r,c,p,q,int
	var b,d,e,cp string
	s := [5][5] string {{"a","b" , "c", "d", "e"},
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
        fmt.Println("Enter the blocked position 1")
	fmt.Scanf("%s", &b)
        fmt.Println("enter the blocked position 2")
        fmt.Scanf("%s", &e)
        fmt.Println("enter the blocked position 3")
        fmt.Scanf("%s",&d)
	p,q=getindex(b)
        p,q=getindex(e)
        p,q=getindex(d)
    moves(r,c,p,q)
	}
	
	func moves(x,y,j,k int) {
		m:=0
		var list1,list2 []int
		if (x== 4){
			 list1 = []int{x-1,x}
		} else if ( x == 0){
			 list1 = []int{x,x+1}
		}else{
			 list1 = []int{x-1,x,x+1}
		}
		if (y== 4){
			 list2 = []int{y-1,y}
		} else if ( y == 0){
			 list2 = []int{y,y+1}
		}else{
			 list2 = []int{y-1,y,y+1}
		}

		for _, a := range list1{
			for _, b := range list2{
			
				if((a == j) && (b == k)){
					continue
				}else if((a == x) && (b == y)){
					continue
				}else{
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


