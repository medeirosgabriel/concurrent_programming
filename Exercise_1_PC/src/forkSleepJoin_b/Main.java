package forkSleepJoin_b;

import java.util.ArrayList;
import java.util.List;

public class Main {
	
	public static void main(String args[]) throws InterruptedException {
		
		int n = 5;
		int leftLimit = 0;
		int rightLimit = 5000;
		Counter counter = new Counter(n);
		List<Thread> threads = new ArrayList<>();
		
		for (int i = 0; i < n; i++) {
			int serviceTime = leftLimit + (int) (Math.random() * (rightLimit - leftLimit));
			MyThread myThread  = new MyThread(i, serviceTime, counter);
			System.out.println(String.format("Starting Thread %d - %d miliseconds", i, serviceTime));
			myThread.start();
			threads.add(myThread);
		}
		
		System.out.println("Waiting Threads");
		
		while (counter.getCounter() != 0) {}
		
		System.out.println("\n====== FINISHED ALL ======");
		
	}

}
