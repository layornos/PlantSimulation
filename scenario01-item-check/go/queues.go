package main

import "github.com/layornos/godes"

/*
 * Setting up of the three queues in the system
 * itemArrival is dedicated to collect the items that must be transported to the quality assurance
 * checking is dedicated to collect the items that musst be checked by the quality assurance
 * machinesToRepair are dedicated to collect machines that must be repaired before they can continue
 */
var itemArrivalQueue = godes.NewFIFOQueue("Item Arrival Queue")
var checkingQueue = godes.NewFIFOQueue("Checking Queue")
var machinesToRepairQueue = godes.NewFIFOQueue("Machine Repair Queue")
