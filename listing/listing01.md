Что выведет программа? Объяснить вывод программы.

```go
package main

import (
"fmt"
)

func main() {
a := [5]int{76, 77, 78, 79, 80}
var b []int = a[1:4]
fmt.Println(b)
}
```
Ответ:
```
[77 78 79]
```
При объявлении b мы объявили срез, начало этого среза ссылается на элемент массива `a` с индексом 1.
`b` имеет длину 3 и емкость 4, `b` будет ссылаться на массив `a` до тех пор пока мы не добавим количество элементов в него, 
превыщающее оставшиеся свободные места в емкости. В данном случае это, до тех пор пока мы не добавим два элемента в срез `b`.
