package single.pattern.n6enumeration;

/**
 * 枚举
 * 这种方式是 Effective Java 作者 Josh Bloch 提倡的方式，它不仅能避免多线程同步问题，而且还自动支持序列化机制，
 * 防止反序列化重新创建新的对象，绝对防止多次实例化。用这种方式写不免让人感觉生疏，在实际工作中，也很少用。
 * 不能通过 reflection attack 来调用私有构造方法。
 * 如果涉及到反序列化创建对象时，可以尝试使用第 6 种枚举方式。
 */
public enum Singleton {
    INSTANCE;
    public void whateverMethod(){

    }
}
