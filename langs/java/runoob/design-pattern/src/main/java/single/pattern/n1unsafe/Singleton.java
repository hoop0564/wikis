package single.pattern.n1unsafe;

/**
 * 懒汉式，线程不安全
 * 不支持多线程。因为没有加锁 synchronized
 */
public class Singleton {
    private static Singleton instance;
    private Singleton(){}
    public static Singleton getInstance(){
        if (instance == null){
            instance = new Singleton();
        }
        return instance;
    }
}
