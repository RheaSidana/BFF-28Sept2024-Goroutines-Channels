// code 1: How Goroutine are created

package main

import (
	"fmt"
)

func serveTable(table int, task string) {
	fmt.Printf("Waiter is %s for Table %d\n", task, table)
	fmt.Printf("Finished %s for Table %d\n", task, table)
}

func main(){
	go serveTable(1, "taking the order")

	fmt.Println("Waiter is attending to other tasks...")
}


//code 2: synchronization using Time.Sleep

/*
package main

import (
	"fmt"
	"time"
)

func serveTable(table int, task string) {
	fmt.Printf("Waiter is %s for Table %d\n", task, table)
	fmt.Printf("Finished %s for Table %d\n", task, table)
}

func main(){
	go serveTable(1, "taking the order")

	time.Sleep(1 * time.Second)
	fmt.Println("Waiter is attending to other tasks...")
}
*/

//code 3: why not Time.Sleep is efficient?

/*
package main

import (
	"fmt"
	"time"
)

func serveTable(table int, task string) {
	fmt.Printf("Waiter is %s for Table %d\n", task, table)
	time.Sleep(1 * time.Second)
	fmt.Printf("Finished %s for Table %d\n\n", task, table)
}

func main() {
	tasks := []string{
		"taking the order",
		"serving the food",
		"cleaning the table",
	}

	for index, task := range tasks {
		go serveTable((index + 1), task)
	}

	time.Sleep(1 * time.Second)  // here
	fmt.Println("Waiter is attending to other tasks...")
}
*/

//code 4: Anonymous Functions

/*
package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := []string{
		"taking the order",
		"serving the food",
		"cleaning the table",
	}

	for index, task := range tasks {
		go func(table int, task string){

			fmt.Printf("Waiter is %s for Table %d\n", task, table)
			time.Sleep(1 * time.Second)
			fmt.Printf("Finished %s for Table %d\n\n", task, table)

		} ((index + 1), task)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Waiter is attending to other tasks...")
}
*/

// code 5: synchronization using WaitGroup
/*
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	tasks := []string{
		"taking the order",
		"serving the food",
		"cleaning the table",
	}

	for index, task := range tasks {
		wg.Add(1)

		go func(table int, task string){

			fmt.Printf("Waiter is %s for Table %d\n", task, table)
			fmt.Printf("Finished %s for Table %d\n\n", task, table)

		} ((index + 1), task)

		wg.Done() // here, why?
	}

	wg.Wait()
	fmt.Println("Waiter is attending to other tasks...")
}
*/

// code 6: synchronization using Mutex

/*
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	tasks := []string{
		"taking the order",
		"serving the food",
		"cleaning the table",
	}

	for index, task := range tasks {
		wg.Add(1)

		go func(table int, task string){
			defer wg.Done()

			mutex.Lock()
			fmt.Printf("Waiter is %s for Table %d\n", task, table)
			time.Sleep(1 * time.Second)
			fmt.Printf("Finished %s for Table %d\n\n", task, table)
			mutex.Unlock()

		} ((index + 1), task)
	}

	wg.Wait()
	fmt.Println("Waiter is attending to other tasks...")
}
*/

// code 7: Parallelism

/*
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// number of waiters: 3
	runtime.GOMAXPROCS(3)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	tables := []int{1, 2, 3, 4, 5, 6}
	tasks := []string{
		"taking the order",
		"serving the food",
		"cleaning the table",
	}

	for _, table := range tables {
		for _, task := range tasks {
			wg.Add(1)

			go func(table int, task string) {
				defer wg.Done()

				mutex.Lock()
				fmt.Printf("Waiter is %s for Table %d\n", task, table)
				time.Sleep(1 * time.Second)
				fmt.Printf("Finished %s for Table %d\n\n", task, table)
				mutex.Unlock()

			}(table, task)
		}
	}

	wg.Wait()
	fmt.Println("Waiters are attending to other tasks...")
}
*/