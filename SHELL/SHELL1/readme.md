# 统计文件的行数

## 描述

写一个 bash脚本以输出一个文本文件 nowcoder.txt中的行数

### 示例

假设 nowcoder.txt 内容如下：

```
#include <iostream>
using namespace std;
int main()
{
    int a = 10;
    int b = 100;
    cout << "a + b:" << a + b << endl;
    return 0;
}
```

你的脚本应当输出：

9

## notes

### wc

```bash
cat nowcoder.txt | wc -l
wc -l < nowcoder.txt
wc -l nowcoder.txt | gawk '{print $1}'
```

### grep

```bash
grep -c "" ./nowcoder.txt
grep -n "" nowcoder.txt | tail -n1 | gawk -F: '{print $1}'
```

### sed

```bash
sed -n '$=' nowcoder.txt
```

### gawk

```bash
gawk '{print NR}' nowcoder.txt | tail -n1
gawk 'END{print NR}' nowcoder.txt
```

