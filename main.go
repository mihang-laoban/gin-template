package main

import "gin-template/lower"

func start() {
	//LimitReal()
	//Run()
	//lower.TestLowerAsync()
	lower.TestLowerSync()
}

//todo consul register services

func main() {
	start()
}
