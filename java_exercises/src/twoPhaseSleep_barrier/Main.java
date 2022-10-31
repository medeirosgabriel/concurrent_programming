package twoPhaseSleep_barrier;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.Semaphore;

public class Main {

	public static void main(String[] args) throws InterruptedException {
		
		int n = 5;
		List<MyThread> threads = new ArrayList<>();
		Semaphore mutex_1 = new Semaphore(1);
		Semaphore finish_mutex = new Semaphore(-(n - 1));
		Semaphore tick_gate_1 = new Semaphore(0);
		Semaphore tick_gate_2 = new Semaphore(1);
		Counter counter = new Counter();
		
		for (int i = 0; i < n; i++) {
			int serviceTime = (int) (Math.random() * (5000));
			MyThread myThread  = new MyThread(n, i, serviceTime, threads, 
						mutex_1, counter, tick_gate_1 , tick_gate_2, finish_mutex);
			myThread.start();
			threads.add(myThread);
		}
		
		finish_mutex.acquire();
		
		System.out.println("\n====== FINISHED ALL ======");

	}
}
