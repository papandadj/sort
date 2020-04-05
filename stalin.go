package sort

import (
	"math/rand"
	"time"
)

//Stalin .
// Gather up the comrades and show them the list.
// Ask the comrades to raise their hands if they believe the list is sorted.
// Select any comrade who did not raise his hand, and executes him as a traitor.
// Repeat steps 2 and 3 until everyone agrees that the list is sorted.
func Stalin(arr ...int) (isOrders bool) {
	//定义一百个同志, 并且给这些士兵随机分配信仰，区间是0-1
	rand.Seed(time.Now().UnixNano())
	var comrades []*Comrade
	for i := 0; i < 100; i++ {
		comrades = append(comrades, &Comrade{
			Index: i, Belief: rand.Float64(), IsOrders: false, IsKilled: false})
	}

	//开始信仰表决
	for i := 0; i < 1000; i++ {
		// fmt.Printf("开始第%d轮表决\n", i)
		//isPass：全部同志认为队列是有序的
		var isPass = true
		for _, comrade := range comrades {
			if comrade.IsKilled == true {
				continue
			}

			if comrade.Belief > 0.5 {
				comrade.IsOrders = true
				//再次制造信仰， 让信仰更加充实
				// fmt.Printf("信仰为%f第%d个士兵认为数组的顺序:%t\n", comrade.Belief, comrade.Index, comrade.IsOrders)
				comrade.Belief = comrade.Belief * (0.5 + rand.Float64())
			} else {
				comrade.IsOrders = false
				// fmt.Printf("---信仰为%f第%d个士兵认为数组的顺序:%t\n", comrade.Belief, comrade.Index, comrade.IsOrders)
				isPass = false
			}

			if comrade.IsOrders == false {
				// fmt.Println("枪决该同志")
				comrade.IsKilled = true
			}
		}
		if isPass {
			return true
		}
	}
	return
}

//Comrade .
type Comrade struct {
	Index    int     //第几个同志
	Belief   float64 //该同志的信仰
	IsOrders bool    //该同志认为队列是否有序， 信仰大于0.5则有序
	IsKilled bool    //该同志是否死了
}
