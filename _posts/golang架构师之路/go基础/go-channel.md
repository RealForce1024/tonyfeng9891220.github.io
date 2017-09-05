# go channel的正确使用姿势

```go
package main

import (
	"fmt"
	"time"
)

type Addr struct {
	City, District string
}

type Person struct {
	Name string
	Age  int
	Addr
}

type PersonHandler interface {
	Batch(<-chan Person) <-chan Person
	Handle(person *Person)
}

type PersonHandlerImpl struct{}

func (handler PersonHandlerImpl) Batch(origs <-chan Person) <-chan Person {
	//dest := make(<-chan Person, 100)
	dest := make(chan Person, 100)
	go func() {
		for p := range origs {
			handler.Handle(&p)
			dest <- p
		}
		fmt.Println("all people has been handled")
		close(dest)
	}()
	return dest
}

func (handler PersonHandlerImpl) Handle(person *Person) {
	if person.District == "haidian" {
		person.District = "changping"
	}
}

var persons []Person
var personTotal int = 200

func init() {
	for i := 0; i < personTotal; i++ {
		persons = append(persons, Person{Name: fmt.Sprintf("P%v", i), Age: 28, Addr: Addr{City: "bj", District: "haidian"}})
	}
}

func getHandler() PersonHandler {
	return PersonHandlerImpl{}
}
func main() {
	defer trace()()
	time.Sleep(time.Second)
	handler := getHandler()
	origs := make(chan Person, 100)
	fetchPersons(origs)
	dest := handler.Batch(origs)
	sign := save(dest)
	<-sign
}
func trace() func() {
	now:=time.Now()
	return func() {
		duration := time.Since(now)
		fmt.Println("运行时间:", duration)
	}
}

// fectch 将向通道中写入
func fetchPersons(origs chan<- Person) {
	go func() {
		for _, p := range persons {
			origs <- p
		}
		fmt.Println("all person has been fetched.")
		close(origs)
	}()
}

// save 从通道中读取
func save(dest <-chan Person) <-chan byte {
	sign := make(chan byte, 1)

	go func() {
		/*for p,ok := range dest {
			if !ok {
				break
			}
			savePerson(p)
		}*/

		for {
			p, ok := <-dest
			if !ok {
				fmt.Println("all people saved")
				sign <- 0
				break
			}
			savePerson(p)
		}
	}()

	//sign <- 1 //这里会同步执行
	return sign
}

func savePerson(p Person) bool {
	fmt.Println(p)
	return true
}

```

