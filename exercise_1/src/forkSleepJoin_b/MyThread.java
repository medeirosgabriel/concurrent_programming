package forkSleepJoin_b;

public class MyThread extends Thread {
	
	private int serviceTime;
	private int id;
	private Counter counter;
	
	public MyThread(int id, int serviceTime, Counter counter) {
		this.id = id;
		this.serviceTime = serviceTime;
		this.counter = counter;
	}
	
	@Override
	public void run() {
		try {
			System.out.println(String.format("Running Thread %d", this.id));
            Thread.sleep(this.serviceTime);
            System.out.println(String.format("Finished Thread %d", this.id));
            this.counter.decrement();
        } catch (InterruptedException e) {
            System.out.println(e.getMessage());
        }
	}
}
