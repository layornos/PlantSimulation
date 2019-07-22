package main

import "github.com/layornos/godes"

type Operators struct {
	max int
}

func (operators *Operators) Catch(item *Item) {
	for {
		operatorAvailableSwt.Wait(true)
		if itemArrivalQueue.GetHead().(*Item).id == item.id {
			break
		} else {
			godes.Yield()
		}
	}
	busyOperators++
	if busyOperators == operators.max {
		operatorAvailableSwt.Set(false)
	}
}

func (operators *Operators) Release(item *Item) {
	busyOperators--
	checkingQueue.Place(item)
	operatorAvailableSwt.Set(true)
}
