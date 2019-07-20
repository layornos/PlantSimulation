package main

import "github.com/layornos/godes"

type QualityAssurance struct {
	max int
}

func (qs *QualityAssurance) Catch(item *Item) {
	for {
		qsAvailableSwt.Wait(true)
		if checkingQueue.GetHead().(*Item).id == item.id {
			break
		} else {
			godes.Yield()
		}
	}
	busyQS++
	if busyQS == qs.max {
		qsAvailableSwt.Set(false)
	}

	if item.defect {
		machinesToRepairQueue.Place(item)
	}
}

func (qs *QualityAssurance) Release(item *Item) {
	busyOperators--
	checkedItems++
	qsAvailableSwt.Set(true)
}
