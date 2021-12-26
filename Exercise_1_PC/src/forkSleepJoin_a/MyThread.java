package forkSleepJoin_a;

public class MyThread extends Thread {
	
	private int serviceTime;
	private int id;
	
	public MyThread(int id, int serviceTime) {
		this.id = id;
		this.serviceTime = serviceTime;
	}
	
	@Override
	public void run() {
		try {
			System.out.println(String.format("Running Thread %d", this.id));
            Thread.sleep(this.serviceTime);
            System.out.println(String.format("Finished Thread %d", this.id));
        } catch (InterruptedException e) {
            System.out.println(e.getMessage());
        }
	}
}
