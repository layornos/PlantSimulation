package edu.kit.ipd.sdq.simulation.abstractsimengine.plant_sim_example;

public class Config {
	public boolean IS_DETERMINISTIC    = true;
	public int MAX_ROBOTS              = 3;
	public int MAX_OPERATORS           = 5;
	public int MAX_MAINTAINERS         = 3;
	public int MAX_QS                  = 1;

	// How long to wait for the n * 100th item
	public int ARRIVAL_INTERVAL        = 10;

	// Time Operator needs to go to the item check up
	public int TRANSPORT_TIME          = 10;
	public double TRANSPORT_TIME_SIGMA = .1;

	public int CHECK_TIME              = 100;
	public double CHECK_TIME_SIGMA     = .1;
	public int REPAIR_TIME             = 3000;
	public double REPAIR_TIME_SIGMA    = 0.2;
	public int SHUTDOWN_TIME           = 5 * 24 * 60;

	public int ITEM_INTACT             = 0;
	public double ITEM_INTACT_SIGMA    = .1;
	public double ITEM_CHECK           = .2;
}
