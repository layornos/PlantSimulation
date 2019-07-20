package main

import "github.com/layornos/godes"

type Maintainers struct {
	max int
}

func (maintainers *Maintainers) Catch(item *Item) {
	for {
		maintainerAvailableSwt.Wait(true)
		robotAvailableSwt.Wait(true)
		if machinesToRepairQueue.GetHead().(*Item).id == item.id {
			break
		} else {
			godes.Yield()
		}
	}
	busyMaintainers++
	robotsAvailable--

	if busyMaintainers == maintainers.max {
		maintainerAvailableSwt.Set(false)
	}

	if robotsAvailable == 0 {
		robotAvailableSwt.Set(false)
	}
}

func (maintainers *Maintainers) Release() {
	busyMaintainers--
	robotsAvailable++
	maintainerAvailableSwt.Set(true)
	robotAvailableSwt.Set(true)
}
