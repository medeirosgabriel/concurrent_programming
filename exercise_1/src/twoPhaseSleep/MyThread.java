package twoPhaseSleep;

import java.util.List;
import java.util.concurrent.Semaphore;

public class MyThread extends Thread {
	
	private int serviceTime;
	private int newServiceTime;
	private int id;
	private List<MyThread> threads;
	private Semaphore mutex1;
	private Semaphore mutex2;
	
	public MyThread(int id, int serviceTime, List<MyThread> threads, Semaphore mutex1, Semaphore mutex2) {
		this.id = id;
		this.serviceTime = serviceTime;
		this.newServiceTime = -1;
		this.threads = threads;
		this.mutex1 = mutex1;
		this.mutex2 = mutex2;
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
            this.mutex1.release();
            threads.get(neightboor).release(newServiceTime);
            this.mutex2.acquire();
            System.out.println(String.format("Running Thread %d - Second Time: %d miliseconds - Choose: %d miliseconds", 
            		this.id, this.newServiceTime, newServiceTime));
            Thread.sleep(this.newServiceTime);
            System.out.println(String.format("Finished Thread %d", this.id));
            this.mutex1.release();
        } catch (InterruptedException e) {
            System.out.println(e.getMessage());
        }
	}
	
	public void release(int serviceTime) {
		this.newServiceTime = serviceTime;
	}
	
	public void setMutex1(Semaphore mutex1) {
		this.mutex1 = mutex1;
	}
}
