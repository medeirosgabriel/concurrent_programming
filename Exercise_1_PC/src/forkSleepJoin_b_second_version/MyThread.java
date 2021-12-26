package forkSleepJoin_b_second_version;

import java.util.concurrent.Semaphore;

public class MyThread extends Thread {
	
	private int serviceTime;
	private int id;
	private Semaphore mutex;
	
	public MyThread(int id, int serviceTime, Semaphore mutex) {
		this.id = id;
		this.serviceTime = serviceTime;
		this.mutex = mutex;
	}
	
	@Override
	public void run() {
		try {
			System.out.println(String.format("Running Thread %d", this.id));
            Thread.sleep(this.serviceTime);
            System.out.println(String.format("Finished Thread %d", this.id));
            this.mutex.release();
        } catch (InterruptedException e) {
            System.out.println(e.getMessage());
        }
	}
}
