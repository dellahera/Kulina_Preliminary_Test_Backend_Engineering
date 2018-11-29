package main

import("fmt")

var (
	n, a, jum, nil int
	bil [100]int
	)
func main(){
	fmt.Print("Masukkan input : ")
	fmt.Scan(&n);
	jum=0;
	nil=1;
	for n>0 {
		jum++;
		if(jum>1) { nil*=10;}
		bil[jum]= (n%10) * nil;
		n/=10;
	}
	for i:=jum; i>=1; i--{
		fmt.Println(bil[i]);
	}
}