package twoPhaseSleep;

import java.util.ArrayList;
import java.util.List;

import twoPhaseSleep.Counter;

public class Main {

	public static void main(String[] args) throws InterruptedException {
		
		int n = 5;
		int leftLimit = 0;
		int rightLimit = 5000;
		List<MyThread> threads = new ArrayList<>();
		Counter counter = new Counter(n);
		
		for (int i = 0; i < n; i++) {
			int serviceTime = leftLimit + (int) (Math.random() * (rightLimit - leftLimit));
			MyThread myThread  = new MyThread(i, serviceTime, threads, counter, leftLimit, rightLimit);
			myThread.start();
			threads.add(myThread);
		}
		
		System.out.println("Waiting Threads");
		
		while (counter.getCounter() != 0) {}
		
		System.out.println("\n====== FINISHED ALL ======");

	}
}
