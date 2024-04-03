package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123


Требования:
Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP. После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться через timeout

*/

// shutdown закрывает соединение и завершает программу
// в случае получения сигналов SIGQUIT, SIGTERM, SIGINT
func shutdown(conn net.Conn) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	<-ch

	conn.Close()

	fmt.Println("Program shutdown")
	os.Exit(0)
}

// telnet устанавливает соединение с сервером
func telnet(address string, timeout time.Duration) error {
	// подключаемся к серверу
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return err
	}

	// закрываем соединение в случае получения сигнала SIGQUIT
	go shutdown(conn)

	// общаемся с сервером
	for {
		// считываем сообщение из стандартного ввода
		reader := bufio.NewReader(os.Stdin)
		request, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		// отправляем сообщение на сервер
		if _, err := fmt.Fprintf(conn, request+"\n"); err != nil {
			return err
		}

		// прослушиваем ответ
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return err
		}

		// печатаем ответ в стандартный вывод
		fmt.Println(response)
	}
}

func main() {
	// устанавливаем флаги
	timeout := flag.Duration("timeout", 10*time.Second, "Установить таймаут")
	flag.Parse()

	// проверяем аргументы
	if len(flag.Args()) != 2 {
		fmt.Fprintln(os.Stderr, "usege: ./telnet [-timeout t] host port")
		os.Exit(1)
	}

	// подготавливаем параметры
	host := flag.Arg(0)
	port := flag.Arg(1)
	address := fmt.Sprintf("%s:%s", host, port)

	if err := telnet(address, *timeout); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
