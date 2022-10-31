package twoPhaseSleep_barrier;

public class Counter {
	
	private int counter;
	
	public Counter() {
		this.counter = 0;
	}
	
	public void decrement() {
		this.counter --;
	}
	
	public void increment() {
		this.counter ++;
	}
	
	public int getCounter() {
		return this.counter;
	}
}
 