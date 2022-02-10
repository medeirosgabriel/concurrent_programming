package twoPhaseSleep;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.Semaphore;

public class Main {

	public static void main(String[] args) throws InterruptedException {
		
		int n = 5;
		List<MyThread> threads = new ArrayList<>();
		Semaphore mutex1 = new Semaphore(-(n - 1));
		Semaphore mutex2 = new Semaphore(0);
		
		for (int i = 0; i < n; i++) {
			int serviceTime = (int) (Math.random() * (5000));
			MyThread myThread  = new MyThread(i, serviceTime, threads, mutex1, mutex2);
			myThread.start();
			threads.add(myThread);
		}
				
		mutex1.acquire();
		mutex1 = new Semaphore(-(n - 1));
		
		for (MyThread t : threads) {
			t.setMutex1(mutex1);
		}
		
		mutex2.release(n);
		mutex1.acquire();
		
		System.out.println("\n====== FINISHED ALL ======");

	}
}
