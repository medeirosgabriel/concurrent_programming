package twoPhaseSleep_barrier;

import java.util.List;
import java.util.concurrent.Semaphore;

public class MyThread extends Thread {
	
	private int serviceTime;
	private int newServiceTime;
	private int id;
	private int n;
	private List<MyThread> threads;
	private Counter counter;
	private Semaphore mutex;
	private Semaphore tick_gate_1;
	private Semaphore tick_gate_2;
	private Semaphore finish_mutex;
	
	public MyThread(int n, int id, int serviceTime, List<MyThread> threads, 
			Semaphore mutex, Counter counter, Semaphore tick_gate_1, Semaphore tick_gate_2, Semaphore finish_mutex) {
		this.id = id;
		this.n = n;
		this.serviceTime = serviceTime;
		this.newServiceTime = -1;
		this.threads = threads;
		this.counter = counter;
		this.mutex = mutex;
		this.tick_gate_1 = tick_gate_1;
		this.tick_gate_2 = tick_gate_2;
		this.finish_mutex = finish_mutex;
	}
	
	private void wait_all() throws InterruptedException {
		this.mutex.acquire();
		this.counter.increment();
		if (this.counter.getCounter() == this.n) {
			this.tick_gate_1.release();
			this.tick_gate_2.acquire();
		}
		this.mutex.release();
		this.tick_gate_1.acquire();
		this.tick_gate_1.release();
		
		this.mutex.acquire();
		this.counter.decrement();
		if (this.counter.getCounter() == 0) {
			this.tick_gate_2.release();
			this.tick_gate_1.acquire();
		}
		this.mutex.release();
		this.tick_gate_2.acquire();
		this.tick_gate_2.release();
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
            int newServiceTime = (int) (Math.random() * 5000);
            threads.get(neightboor).release(newServiceTime);
            this.wait_all();
            System.out.println(String.format("Running Thread %d - Second Time: %d miliseconds - Choose: %d miliseconds", 
            		this.id, this.newServiceTime, newServiceTime));
            Thread.sleep(this.newServiceTime);
            System.out.println(String.format("Finished Thread %d", this.id));
            this.finish_mutex.release();
        } catch (InterruptedException e) {
            System.out.println(e.getMessage());
        }
	}
	
	public void release(int serviceTime) {
		this.newServiceTime = serviceTime;
	}
}
