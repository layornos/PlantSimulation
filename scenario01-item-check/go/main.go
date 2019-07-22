package main

import (
	"fmt"
	"strconv"

	"github.com/layornos/godes"
)

var arrivalOfItems = godes.NewExpDistr(IS_DETERMINISTIC)
var transportItemToCheck = godes.NewNormalDistr(IS_DETERMINISTIC)
var checkOfItem = godes.NewNormalDistr(IS_DETERMINISTIC)
var repairOfRobot = godes.NewExpDistr(IS_DETERMINISTIC)
var repairTimeOfOneRobot = godes.NewNormalDistr(IS_DETERMINISTIC)
var itemIntact = godes.NewNormalDistr(IS_DETERMINISTIC)

var robotAvailableSwt = godes.NewBooleanControl()
var operatorAvailableSwt = godes.NewBooleanControl()
var maintainerAvailableSwt = godes.NewBooleanControl()
var qsAvailableSwt = godes.NewBooleanControl()

var robotsAvailable = MAX_ROBOTS
var busyOperators = 0
var busyMaintainers = 0
var busyQS = 0
var defectCount = 0

var operators *Operators
var maintainers *Maintainers
var qs *QualityAssurance

var itemCount = 0
var itemsProcessed = 0
var checkedItems = 0
var repairedRobots = 0

func main() {
	itemArrivalQueue.Clear()
	checkingQueue.Clear()
	machinesToRepairQueue.Clear()

	operators = &Operators{MAX_OPERATORS}
	maintainers = &Maintainers{MAX_MAINTAINERS}
	qs = &QualityAssurance{MAX_QS}

	operatorAvailableSwt.Set(true)
	maintainerAvailableSwt.Set(true)
	robotAvailableSwt.Set(true)
	qsAvailableSwt.Set(true)

	godes.Run()

	for {
		robotAvailableSwt.Wait(true)
		var item *Item
		if itemIntact.Get(ITEM_INTACT, ITEM_INTACT_SIGMA) > ITEM_CHECK {
			item = &Item{&godes.Runner{}, strconv.Itoa(itemCount), true}
		} else {
			item = &Item{&godes.Runner{}, strconv.Itoa(itemCount), false}
		}
		itemArrivalQueue.Place(item)
		godes.AddRunner(item)
		godes.Advance(arrivalOfItems.Get(1. / ARRIVAL_INTERVAL))
		if godes.GetSystemTime() > SHUTDOWN_TIME {
			break
		}

	}

	godes.WaitUntilDone()
	fmt.Printf("Number of checked items: %d\n", checkedItems)
	fmt.Printf("Number of defect items: %d\n", defectCount)
	fmt.Printf("Number of repaired robots: %d\n", repairedRobots)
	fmt.Println("Done")
}
