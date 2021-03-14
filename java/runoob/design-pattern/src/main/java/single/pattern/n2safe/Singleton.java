package single.pattern.n2safe;

/**
 * 懒汉式，线程安全
 * getInstance() 的性能对应用程序不是很关键（该方法使用不太频繁）。
 */
public class Singleton {
    public static Singleton instance;
    private Singleton(){}
    public static synchronized Singleton getInstance(){
        if (instance == null) {
            instance = new Singleton();
        }
        return instance;
    }
}
