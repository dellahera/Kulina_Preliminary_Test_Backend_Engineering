package main

import("fmt")

var (
	n, i, a,b,c int
	)
func main(){
	fmt.Print("Masukkan input : ")
	fmt.Scan(&n);
	fmt.Println(n," bilangan fibonaci pertama adalah:");
	a=1;
	b=1;
	c=0;
	fmt.Print(a," ");
	if(n>1) {
		for i := 2; i <=n; i++ {
			fmt.Print(b," ");
			c=a+b;
			a=b;
			b=c;
		}
	}
}