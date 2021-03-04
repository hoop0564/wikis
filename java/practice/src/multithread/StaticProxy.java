package multithread;

public class StaticProxy {
    public static void main(String[] args) {
        // 线程就是静态代理模式
        new Thread(() -> System.out.println("good job")).start();

        new WeddingCompany(new You()).HappyMarry();
    }
}

interface Marry {
    void HappyMarry();
}

class You implements Marry {

    @Override
    public void HappyMarry() {
        System.out.println("you're married, happy!?");
    }
}

class WeddingCompany implements Marry {

    private Marry target;

    public WeddingCompany(Marry marry) {
        this.target = marry;
    }

    @Override
    public void HappyMarry() {
        pre();
        this.target.HappyMarry();
        post();
    }

    private void pre() {
        System.out.println("before marry: prepare dinner");
    }

    private void post() {
        System.out.println("after marry: got hong bao");
    }
}