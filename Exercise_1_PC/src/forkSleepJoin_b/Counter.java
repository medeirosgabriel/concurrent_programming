package forkSleepJoin_b;

import java.util.concurrent.Semaphore;

public class Counter {
	
	private Semaphore mutex;
	private int counter;
	
	public Counter(int n) {
		this.mutex = new Semaphore(1);
		this.counter = n;
	}
	
	public void decrement() throws InterruptedException {
		this.mutex.acquire();
		this.counter --;
		this.mutex.release();
	}
	
	public int getCounter() throws InterruptedException {
		this.mutex.acquire();
		int c = this.counter;
		this.mutex.release();
		return c;
	}
}
 