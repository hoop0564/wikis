package single.pattern.n5staticinternal;

/**
 * 登记式/静态内部类
 * classloader 机制来保证初始化 instance 时只有一个线程
 * 这种方式只适用于静态域的情况，双检锁方式可在实例域需要延迟初始化时使用。
 */
public class Singleton {
    public static class SingletonHolder {
        private static final Singleton INSTANCE = new Singleton();
    }
    private Singleton(){}
    public static final Singleton getInstance(){
        return SingletonHolder.INSTANCE;
    }
}
