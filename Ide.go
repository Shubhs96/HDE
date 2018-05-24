/* Code by Shubham Saini */
package main
import (
    "fmt"
    )
    //declaring variablesin int32 format
    var n int32
    var y int32
    var s int32
    var x int32
    
func main() {
    fmt.Scanf("%d",&n)//Input N value
	//check n value according to contraints
    
    if n<1 || n>100 {
        return
    }else if n>1 || n<100 {
        testcase(n) //call for testcase function
    }
}


func sumofpositive(z int32, sum int32) int32 {
    if z==0 {
        return sum
    }else 
    {
        fmt.Scanf("%d",&y)
        if y < -100 || y> 100 {
            return 0
            
        }else if y<0 {
            return sumofpositive(z-1,sum)
        }else {
            sum=sum+(y*y)
            return sumofpositive(z-1,sum)
        }
    }
}
func testcase(t int32) int32 {
    if t == 0 {
        return 0
    }else
    {
        fmt.Scanf("%d",&x)
        if x>0 && x<=100 {
            s=sumofpositive(x,0)
            fmt.Println(s)
            testcase(t-1)
            return s
        }else
        {
            return 0
        }
    }
    
}
