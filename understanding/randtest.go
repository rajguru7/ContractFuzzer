package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	test()
}

type Fuzzer_Rand struct{
   seed int
   last_chose int
   r  *rand.Rand
}
func NewFuzzer_Rand() *Fuzzer_Rand{
	return &Fuzzer_Rand{seed:time.Now().Nanosecond(),last_chose:-1,r:rand.New(rand.NewSource((int64)(time.Now().Nanosecond())))}
}


func test() (error) {
	g_robin := NewFuzzer_Rand()
	fmt.Print((*g_robin).seed, (*g_robin).r.Intn(5))
	return nil
}

