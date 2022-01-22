package forkSleepJoin_b_second_version;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.Semaphore;

public class Main {
	
	public static void main(String args[]) throws InterruptedException {
		
		int n = 5;
		Semaphore mutex = new Semaphore(-(n-1));
		int leftLimit = 0;
		int rightLimit = 5000;
		List<Thread> threads = new ArrayList<>();
		
		for (int i = 0; i < n; i++) {
			int serviceTime = leftLimit + (int) (Math.random() * (rightLimit - leftLimit));
			MyThread myThread  = new MyThread(i, serviceTime, mutex);
			System.out.println(String.format("Starting Thread %d - %d miliseconds", i, serviceTime));
			myThread.start();
			threads.add(myThread);
		}
		
		System.out.println("Waiting Threads");
		
		mutex.acquire();
		
		System.out.println("\n====== FINISHED ALL ======");
		
		mutex.release();
		
	}

}
