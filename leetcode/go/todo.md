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
# 访问加锁问题
```bash
 vim /home/fjp/.lc/leetcode.cn/user.json
# ,"paid":true
```
# TODO
