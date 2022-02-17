给你一个文本串 T ，一个非空模板串 S ，问 S 在 T 中出现了多少次 

![cbafa37593da8ed1666bed79f788ee7](F:\markdown笔记\刷题\字符串\cbafa37593da8ed1666bed79f788ee7.jpg)

![4e3152371bc1565426450059c4a8de3](F:\markdown笔记\刷题\字符串\4e3152371bc1565426450059c4a8de3.jpg)

```python
class Solution:
    def kmp(self , S , T ):
        # 求 prefix table
        prefix = [0] * len(S)
        i = 1
        while i <len(S):
            if S[i] == S[prefix[i-1]]:
                prefix[i] = prefix[i-1] + 1
                i += 1
            else:
                count = prefix[i-1]
                while count > 0 and S[i] != S[count]:
                    count = prefix[count-1]
                if S[i] != S[count]:
                    prefix[i] = 0
                else:
                    prefix[i] = count + 1
                i += 1
        prefix.insert(0, -1)
        prefix = prefix[:-1]
        
        i = 0
        j = 0
        count = 0 
        while i < len(T):
            if T[i] == S[j]:
                if j == len(S)-1:
                    count += 1
                    j = prefix[j]
                    continue
                i += 1
                j += 1
            else:
                j = prefix[j]
                if j == -1:
                    i+=1
                    j+=1
        return count
```

