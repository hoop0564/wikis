package multithread;

public class TestThread {
    public static void main(String[] args) {
        RunnableDemo r1 = new RunnableDemo("thread-1");
//        r1.start();

        new Thread(r1, "thread-new").start();

        RunnableDemo r2 = new RunnableDemo("thread-2");
//        r2.start();
    }
}
