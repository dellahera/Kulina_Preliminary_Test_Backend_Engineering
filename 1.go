package main

import("fmt")

var (
	n, i, j int
	cek bool
	)
func main(){
	fmt.Print("Masukkan input : ")
	fmt.Scan(&n);
	fmt.Println("Bilangan prima antara 1 sampai ",n," adalah:");
	for i := 2; i <=n; i++ {
		cek=true;
		for j := 2; j*j <= i; j++{
			if(i%j==0) {
				cek=false;
			}
		}
		if(cek){
		fmt.Print(i," ");
		}
	}
	
}