package twoPhaseSleep;

import java.util.List;
import java.util.concurrent.Semaphore;

public class MyThread extends Thread {
	
	private int serviceTime;
	private int id;
	private List<MyThread> threads;
	private int leftLimit;
	private int rightLimit;
	private Semaphore mutex;
	private Counter counter;
	
	public MyThread(int id, int serviceTime, List<MyThread> threads, Counter counter, int leftLimit, int rightLimit) {
		this.id = id;
		this.serviceTime = serviceTime;
		this.threads = threads;
		this.leftLimit = leftLimit;
		this.rightLimit = rightLimit;
		this.mutex = new Semaphore(0);
		this.counter = counter;
	}
	
	@Override
	public void run() {
		try {
			System.out.println(String.format("Running Thread %d - First Time - %d miliseconds", 
					this.id, this.serviceTime));
            Thread.sleep(this.serviceTime);
            System.out.println(String.format("Finished Thread %d", 
            		this.id));
            
            int neightboor = (this.id + 1) % this.threads.size();
            int newServiceTime = this.leftLimit + (int) (Math.random() * (this.rightLimit - this.leftLimit));
            threads.get(neightboor).release(newServiceTime);
            this.mutex.acquire();
            System.out.println(String.format("Running Thread %d - Second Time - %d miliseconds - %d miliseconds", 
            		this.id, this.serviceTime, newServiceTime));
            Thread.sleep(this.serviceTime);
            this.mutex.release();
            System.out.println(String.format("Finished Thread %d", this.id));
            this.counter.decrement();
        } catch (InterruptedException e) {
            System.out.println(e.getMessage());
        }
	}
	
	public void release(int serviceTime) {
		this.serviceTime = serviceTime;
		this.mutex.release();
	}
}
