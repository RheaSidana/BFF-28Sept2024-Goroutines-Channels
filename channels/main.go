//code 1: simple channel

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	waterTank := make(chan int)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(tapID int) {
			defer wg.Done()
			waterInLitre := 3
			fmt.Printf("Tap %d: Pouring water %dlitre into the tank\n", tapID, waterInLitre)
			waterTank <- waterInLitre
			time.Sleep(1 * time.Second)

		}(i)
	}

	go func() {
		for water := range waterTank {
			fmt.Printf("Outlet: Releasing water %dlitre from the tank\n", water)
			time.Sleep(2 * time.Second)
		}
	}()

	wg.Wait()

	close(waterTank)

	fmt.Println("\nAll taps are closed, and the water tank is emptying.")
}

// code 2: Unbuffered channel
/*
package main

import (
	"fmt"
	"sync"
)

// Function to simulate the water taps (send-only channel)
func tap(tapID int, waterTank chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	waterInLitre := 3
	fmt.Printf("Tap %d: Pouring %d litres into the tank\n", tapID, waterInLitre)
	waterTank <- waterInLitre 
}

// Function to simulate the outlet (receive-only channel)
func outlet(waterTank <-chan int) {
	for water := range waterTank {
		fmt.Printf("Outlet: Releasing %d litres from the tank\n", water)
	}
}

func main() {
	waterTank := make(chan int)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go tap(i, waterTank, &wg) 
	}
	
	go outlet(waterTank) 
	wg.Wait()


	close(waterTank)

	fmt.Println("\nAll taps are closed, and the water tank is emptying.")
}
*/

// code 3: buffered channels

/*
package main

import (
	"fmt"
	"sync"
)

// Function to simulate the water taps (send-only channel)
func tap(tapID int, waterTank chan<- int) {
	waterInLitre := 3
	fmt.Printf("Tap %d: Pouring %d litres into the tank\n", tapID, waterInLitre)
	waterTank <- waterInLitre 
}

// Function to simulate the outlet (receive-only channel)
func outlet(waterTank <-chan int, wg *sync.WaitGroup) {
	for water := range waterTank {
		fmt.Printf("Outlet: Releasing %d litres from the tank\n", water)
		wg.Done()
	}
}

func main() {
	waterTank := make(chan int, 2) // buffere = 2

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go tap(i, waterTank)
	}

	go outlet(waterTank, &wg)

	wg.Wait()

	close(waterTank)

	fmt.Println("\nAll taps are closed, and the water tank is emptying.")
}
*/

// code 4: select statement

/*
package main

import (
	"fmt"
	"sync"
	"time"
)

// Function to simulate the water taps (send-only channel)
func tap(tapID int, waterTank chan<- int) {
	waterInLitre := 3
	fmt.Printf("Tap %d: Pouring %d litres into the tank\n", tapID, waterInLitre)
	waterTank <- waterInLitre // Send water to the tank
	time.Sleep(1 * time.Second) // Simulate time delay for pouring
}

// Function to simulate the outlet (receiving water from the tank)
func outlet(waterTank <-chan int, wg *sync.WaitGroup) {
	for {
		select {
		case water, ok := <-waterTank: // Receive water from the tank
			if !ok { // If the channel is closed, exit the function
				return
			}
			fmt.Printf("Outlet: Releasing %d litres from the tank\n", water)
			time.Sleep(2 * time.Second) // Simulate time to release water
			wg.Done()
		}
	}
}

func main() {
	waterTank := make(chan int, 2)
	
	var wg sync.WaitGroup

	// Start the tap Goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go tap(i, waterTank)
	}

	go outlet(waterTank, &wg)

	wg.Wait()

	close(waterTank)

	fmt.Println("\nAll taps are closed, and the water tank is emptying.")
}
*/