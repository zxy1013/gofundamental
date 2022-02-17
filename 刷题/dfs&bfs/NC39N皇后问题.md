 N 皇后问题是指在 n * n 的棋盘上要摆 n 个皇后，
 要求：任何两个皇后不同行，不同列也不在同一条斜线上，
 求给一个整数 n ，返回 n 皇后的摆法数。 

>         """
>         * 解题思路：
>         * 由于需要在n*n的棋盘格中放入n个皇后，就必须每一行放一个
>         * 否则就会出现一行有两个皇后的情况，会发生冲突。"""

```python
class Solution:
    def Nqueen(self , n ):
        res = 0
        def check(row,col,temp):
            # 同行
            for j in range(col):
                if temp[row][j] == "Q" :
                    return False
            # 同列
            for i in range(row):
                if temp[i][col] == "Q" :
                    return False
            # 左上 因为是从上往下填，所以只需要查看上面的是否冲突即可
            r,c = row,col
            while r > 0 and c > 0:
                if temp[r-1][c-1] == "Q" :
                    return False
                r = r-1
                c = c-1
            # 右上
            r,c = row,col
            while r > 0 and c < n - 1:
                if temp[r-1][c+1] == "Q" :
                    return False
                r = r-1
                c = c+1
            return True
        
        def dfs(row,temp):
            nonlocal res
            if row == n:
                res += 1
            else:
                for i in range(n):
                    if check(row,i,temp):
                        temp[row] = temp[row][:i]+"Q"+temp[row][i+1:]
                        dfs(row+1,temp)
                        temp[row] = temp[row][:i]+"."+temp[row][i+1:]
        chess = ["." * n for _ in range(n)]
        dfs(0,chess)
        return res
```

 N皇后问题是把N个皇后放在一个N×N棋盘上，使皇后之间不会互相攻击。 

```go
# @param n int整型 
# @return string字符串二维数组
import copy
class Solution:
    def solveNQueens(self , n ):
        res = []
        def check(row,col,temp):
            # 同行
            for j in range(col):
                if temp[row][j] == "Q" :
                    return False
            # 同列
            for i in range(row):
                if temp[i][col] == "Q" :
                    return False
            # 左上
            r,c = row,col
            while r > 0 and c > 0:
                if temp[r-1][c-1] == "Q" :
                    return False
                r = r-1
                c = c-1
            # 右上
            r,c = row,col
            while r > 0 and c < n - 1:
                if temp[r-1][c+1] == "Q" :
                    return False
                r = r-1
                c = c+1
            return True
        
        def dfs(row,temp):
            nonlocal res
            if row == n:
                res.append(copy.deepcopy(temp))
            else:
                for i in range(n):
                    if check(row,i,temp):
                        temp[row] = temp[row][:i]+"Q"+temp[row][i+1:]
                        dfs(row+1,temp)
                        temp[row] = temp[row][:i]+"."+temp[row][i+1:]
        chess = ["." * n for _ in range(n)]
        print(chess)
        dfs(0,chess)
        return res
```

