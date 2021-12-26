package forkSleepJoin_a;

import java.util.ArrayList;
import java.util.List;

public class Main {
	
	public static void main(String args[]) throws InterruptedException {
		
		int n = 5;
		int leftLimit = 5000;
		int rightLimit = 10000;
		List<Thread> threads = new ArrayList<>();
		
		for (int i = 0; i < n; i++) {
			int serviceTime = leftLimit + (int) (Math.random() * (rightLimit - leftLimit));
			MyThread myThread  = new MyThread(i, serviceTime);
			System.out.println(String.format("Starting Thread %d - %d miliseconds", i, serviceTime));
			myThread.start();
			threads.add(myThread);
		}
		
		System.out.println("Waiting Threads");
		
		for (Thread t : threads) {
			t.join();
		}
		
		System.out.println("====== FINISHED ALL ======");
		
	}

}
