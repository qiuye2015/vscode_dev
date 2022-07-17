# 配置默认模板
- 修改文件`codeonly.tpl`; 可能在wsl上，或者本机
- `~/.vscode-server/extensions/leetcode.vscode-leetcode-0.18.1/node_modules/vsc-leetcode-cli/templates`
- `C:\Users\xxxx\.vscode\extensions\leetcode.vscode-leetcode-0.18.1\node_modules\vsc-leetcode-cli\templates\codeonly.tpl`
```go
package main
```

```c++
${comment.singleLine} @lc code=start
#include <vector>
#include <list>
#include <iostream>
#include <map>
#include <stack>
#include <bitset>
#include <string>
#include <algorithm>
${code}
${comment.singleLine} @lc code=end

int main(){
    Solution s;
    //std::cout<<s<<std::endl;
    return 0;
}
```
# TODO
1. 92.反转链表2 DONE
2. 200.岛屿数量 DONE
3. 3.无重复字符的最长字串 DONE
4. 剑指offer40. 最小的k个数 DONE
5. 56.合并区间 DONE
6. 78.子集 DONE
7. 129.求根到叶子节点数字之和 DONE
8. 141.环形链表 DONE
9. 15.三数之和 DONE
10. 54螺旋矩阵 DONE
11. 103.二叉树的锯齿形层次遍历 DONE