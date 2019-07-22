package main

import "github.com/layornos/godes"

type Item struct {
	*godes.Runner
	id     string
	defect bool
}

func (item *Item) Run() {
	operators.Catch(item)
	itemArrivalQueue.Get()
	godes.Advance(transportItemToCheck.Get(TRANSPORT_TIME, TRANSPORT_TIME_SIGMA))
	operators.Release(item)

	qs.Catch(item)
	checkingQueue.Get()
	godes.Advance(checkOfItem.Get(CHECK_TIME, CHECK_TIME_SIGMA))
	if item.defect {
		defectCount++
		maintainers.Catch(item)
		machinesToRepairQueue.Get()
		godes.Advance(repairTimeOfOneRobot.Get(REPAIR_TIME, REPAIR_TIME_SIGMA))
		maintainers.Release()
		repairedRobots++
	}

	qs.Release(item)
	itemsProcessed++
}
