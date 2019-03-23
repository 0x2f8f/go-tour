package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	//gor1() //запуск 10 горутин
	//mutex1() //использование мьютекса
	//lock1() //множественные блокировки
	//chan1() //использование буферизованных каналов для записи и чтения
	//select1() //Мы посылаем 20 сообщений/сек, а как обработчики могут принять 10, поэтому половина будет отброшена
	//timeAfter1() //Блокировка на максимально возможное время
	//timeAfter2() //Вывод кол-ва секунд таймаута
	timeAfter3() //запуск с перерывами
}

var (
	counter = 0
	lock sync.Mutex
)

//запуск 10 горутин
func gor1()  {
	for i := 0; i < 10; i++ {
		go incr()
	}
	time.Sleep(time.Millisecond * 10) //без time.Sleep горутины работать не будут!
}

func incr() {
	counter++
	fmt.Println(counter)
}

/**
	использование мьютекса для блокировок
	в этом случае горутины будут выполняться в правильном порядке
	и на экран будет выведено 1 2 3 4 5 6 7 8 9 10
 */
func mutex1()  {
	for i := 0; i < 10; i++ {
		go incr2()
	}
	time.Sleep(time.Millisecond * 10)
}

func incr2() {
	lock.Lock()
	defer lock.Unlock()
	counter++
	fmt.Println(counter)
}

/**
	множественные блокировки
	выведтся ошибка - all goroutines are asleep - deadlock!
 */
func lock1() {
	go func() { lock.Lock() }()
	time.Sleep(time.Millisecond * 10)
	lock.Lock()
}

type Worker struct {
	id int
}

func (w Worker) process(ch chan int) {
	for {
		data := <-ch
		fmt.Printf("обработчик %d получил %d\n", w.id, data)
		time.Sleep(time.Millisecond * 500)
	}
}

/**
	горутины и каналы
	канал объемом 10
	пишем в канал быстрее, чем из него читаем
 */
func chan1()  {
	ch := make(chan int, 10)
	for i := 0; i < 5; i++ {
		w := Worker{id: i}
		go w.process(ch)
	}

	j:=1
	for {
		ch <- j
		fmt.Println(len(ch))
		j+=1
		time.Sleep(time.Millisecond * 50)
	}
}

/**
  Мы посылаем 20 сообщений в секунду, а обработчики могут принять только 10, поэтому половина будет отброшена
 */
func select1()  {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		w := Worker{id: i}
		go w.process(ch)
	}

	j:=1
	for {
		select {
			case ch <- j:
				j += 1

			default:
				//тут можно ничего не писать, чтобы данные молча отбрасывались
				fmt.Println("выброшено")
		}
		time.Sleep(time.Millisecond * 50)
	}
}

/**
  Блокировка на максимально возможное время
 */
func timeAfter1()  {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		w := Worker{id: i}
		go w.process(ch)
	}

	j:=1
	for {
		select {
		case ch <- j:
			j += 1
		case <-time.After(time.Millisecond * 100):
			fmt.Println("тайм-аут", )
		}
		time.Sleep(time.Millisecond * 50)
	}
}

/**
  Вывод кол-ва секунд таймаута

- Первый доступный канал будет выбран.
- Если доступно несколько каналов, будет выбран случайный.
- Если нет доступных каналов, будет выполнен блок default.
- Если блока default нет, выполнение select блокируется.
 */
func timeAfter2()  {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		w := Worker{id: i}
		go w.process(ch)
	}

	j:=1
	for {
		select {
		case ch <- j:
			j += 1
		case t := <-time.After(time.Millisecond * 100):
			fmt.Println("тайм-аут после ", t)
		}
		time.Sleep(time.Millisecond * 50)
	}
}

func (w Worker) process2(ch chan int) {
	for {
		select {
		case data := <-ch:
			fmt.Printf("обработчик %d получил %d\n", w.id, data)
		case <-time.After(time.Millisecond * 10):
			fmt.Println("Перерыв")
			time.Sleep(time.Second)
		}
	}
}

//запуск с перерывами
func timeAfter3()  {
	ch := make(chan int, 10)
	for i := 0; i < 5; i++ {
		w := Worker{id: i}
		go w.process2(ch)
	}


	for j:=1; j<50; j++ {
		ch <- j
		time.Sleep(time.Millisecond * 50)
	}
}