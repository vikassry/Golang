package main
import "fmt"

func mult(a, b int) int {
    result := 0;
    for b != 0 {              
        if b&1 == 1 {              
            result = result + a;  
        }
        a<<=1;                    
        b>>=1;                    
    }
    return result
}

func main() {
	fmt.Println("result: %d,", mult(3,4));
}