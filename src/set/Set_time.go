package set

import (
	"fmt"
	"time"
)

var set3 = ItemSet{}
var set4 = ItemSet{}

func initSet(){
	set3.Clear()
	set4.Clear()

	set3.Add(1)
	set3.Add(2)
	set3.Add(3)
	set3.Add(4)

	set4.Add(3)
	set4.Add(4)
	set4.Add(5)
	set4.Add(6)
}

func clearSet(){
	set3.Clear()
	set4.Clear()
}

func AddTime() {
	start := time.Now()
	set3.Add(7)
	elapsed := time.Since(start)

	fmt.Printf("Set add took %d nano sec\n", elapsed.Nanoseconds())
}

func DeleteTime(){
	start := time.Now()
	set3.Delete(7)
	elapsed := time.Since(start)

	fmt.Printf("Set delete took %d nano sec\n", elapsed.Nanoseconds())
}

func IntersectionTime(){
	initSet()

	start := time.Now()
	set3.Intersection(set4)
	elapsed := time.Since(start)

	fmt.Printf("Set intersection took %d nano sec\n", elapsed.Nanoseconds())
}

func DifferenceTime(){
	initSet()

	start := time.Now()
	set3.Difference(set4)
	elapsed := time.Since(start)

	fmt.Printf("Set difference took %d nano sec\n", elapsed.Nanoseconds())
}

func UnionTime(){
	initSet()

	start := time.Now()
	set3.Union(set4)
	elapsed := time.Since(start)

	fmt.Printf("Set union took %d nano sec\n", elapsed.Nanoseconds())
}

func GetAllTimes(){
	AddTime()
	clearSet()
	DeleteTime()
	clearSet()
	IntersectionTime()
	clearSet()
	DifferenceTime()
	clearSet()
	UnionTime()
	clearSet()
}