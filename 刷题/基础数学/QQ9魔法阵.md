小Q搜寻了整个魔法世界找到了四块魔法石所在地，当4块魔法石正好能构成一个正方形的时候将启动魔法阵，小Q就可以借此实现一个愿望。现在给出四块魔法石所在的坐标，小Q想知道他是否能启动魔法阵

输入描述：

输入的第一行包括一个整数（1≤T≤5）表示一共有T组数据

每组数据的第一行包括四个整数`x[i](0≤x[i]≤10000)`，即每块魔法石所在的横坐标

每组数据的第二行包括四个整数`y[i](0≤y[i]≤10000)`,即每块魔法石所在的纵坐标

> -  对角线相等的菱形是正方形 
> -  四条边都相等的是菱形。 

```go
package main
import (
    "fmt"
    "math"
)
func main(){
    var num int 
    fmt.Scan(&num)
    for i:=0;i<num;i++{
        // 读取坐标
        x := make([]int,4)
        tempx := 0
        y := make([]int,4)
        tempy := 0
        fmt.Scanf("%d",&tempx)
        fmt.Scanf("%d",&tempy)
        for i:=3;i>-1;i--{
            x[i] = tempx % 10
            y[i] = tempy % 10
            tempx /= 10
            tempy /= 10
        }
        // 计算四条边 判断是否相等 不用开方了  
        /*
        x0 x1
        x2 x3
        */
        len1 := int(math.Pow(float64(x[0]-x[1]),2) +  math.Pow(float64(y[0]-y[1]),2)) 
        len2 := int(math.Pow(float64(x[0]-x[2]),2) +  math.Pow(float64(y[0]-y[2]),2)) 
        len3 := int(math.Pow(float64(x[1]-x[3]),2) +  math.Pow(float64(y[1]-y[3]),2)) 
        len4 := int(math.Pow(float64(x[3]-x[2]),2) +  math.Pow(float64(y[3]-y[2]),2)) 
        if len1 != len2 || len2 != len3 || len3 != len4 || len4 != len1{
            fmt.Println("No")
            continue
        }
        // 计算对角线
        len5 := int(math.Pow(float64(x[0]-x[3]),2) +  math.Pow(float64(y[0]-y[3]),2)) 
        len6 := int(math.Pow(float64(x[1]-x[2]),2) +  math.Pow(float64(y[1]-y[2]),2)) 
        if len5 != len6 {
            fmt.Println("No")
            continue
        }
        fmt.Println("Yes")
    }
}
```

