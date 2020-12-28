package lower

import (
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"math/rand"
	"sync"
	"time"
)

type Product struct {
	ID    int
	Title string
	Price int
}

func GetProduct() (Product, error) {
	r := rand.Intn(10)
	if r < 6 {
		time.Sleep(time.Second * 5)
	}

	return Product{ID: 12, Title: "apple", Price: 32}, nil
}

func GetSpare() (Product, error) {
	return Product{ID: 0, Title: "banana", Price: 22}, nil
}

func TestLowerSync() {
	rand.Seed(time.Now().UnixNano())

	configA := hystrix.CommandConfig{
		Timeout:                1000,
		RequestVolumeThreshold: 3,
		ErrorPercentThreshold:  20,
		SleepWindow:            int(time.Second * 10),
	}
	hystrix.ConfigureCommand("my_command", configA)

	for {
		err := hystrix.Do("my_command", func() error {
			p, _ := GetProduct()
			fmt.Println(p)
			return nil
		}, func(e error) error {
			fmt.Println(GetSpare())
			return errors.New("my time out")
		})

		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second * 1)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func TestLowerAsync() {
	rand.Seed(time.Now().UnixNano())

	configA := hystrix.CommandConfig{
		Timeout:               1000,
		MaxConcurrentRequests: 5,
	}
	hystrix.ConfigureCommand("my_command", configA)

	resultChan := make(chan Product, 1)
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go (func() {
			wg.Add(1)
			defer wg.Done()
			errs := hystrix.Go("my_command", func() error {
				p, _ := GetProduct()
				resultChan <- p
				return nil
			}, func(e error) error {
				rcp, err := GetSpare()
				resultChan <- rcp
				fmt.Println(e)
				return err
			})

			select {
			case getProd := <-resultChan:
				fmt.Println(getProd)
			case err := <-errs:
				fmt.Println(err)
			}
			time.Sleep(time.Second * 1)
		})()
	}
	wg.Wait()
}
