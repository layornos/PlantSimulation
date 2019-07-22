package edu.kit.ipd.sdq.simulation.abstractsimengine.plant_sim_example;

import org.apache.commons.math3.distribution.ExponentialDistribution;
import org.apache.commons.math3.distribution.NormalDistribution;

public class Main {
	
	
	ExponentialDistribution arrivalOfItems = new ExponentialDistribution(0);
	NormalDistribution transportItemToCheck = new NormalDistribution();
	NormalDistribution checkOfItem = new NormalDistribution();
	ExponentialDistribution repairOfRobot = new ExponentialDistribution(0);
	NormalDistribution repairTimeOfOneRobot = new NormalDistribution();
	NormalDistribution itemIntact = new NormalDistribution();

	public static void main(String[] args) {
		// TODO Auto-generated method stub

	}
	
}
