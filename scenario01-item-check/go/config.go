package main

//Input Parameters
const (
	IS_DETERMINISTIC = true
	MAX_ROBOTS       = 3
	MAX_OPERATORS    = 5
	MAX_MAINTAINERS  = 3
	MAX_QS           = 1

	// How long to wait for the n * 100th item
	ARRIVAL_INTERVAL = 10

	// Time Operator needs to go to the item checkup
	TRANSPORT_TIME       = 10
	TRANSPORT_TIME_SIGMA = .1

	CHECK_TIME        = 100
	CHECK_TIME_SIGMA  = .1
	REPAIR_TIME       = 3000
	REPAIR_TIME_SIGMA = .2
	SHUTDOWN_TIME     = 5 * 24 * 60

	ITEM_INTACT       = 0
	ITEM_INTACT_SIGMA = .1
	ITEM_CHECK        = .2
)
