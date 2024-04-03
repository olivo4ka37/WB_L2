Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```

error - это интерфейс находящийся в Go\src\builtin\builtin.go
```go
type error interface {
	Error() string
}
```
Структура интерфейса будет равна `nil`, только если и тип, и значение равны `nil`. В нашем случае функция `test` возвращает интерфейс,  в котором данные будут `nil`, однако тип будет определен (`customError`) и не будет равняться `nil`.
