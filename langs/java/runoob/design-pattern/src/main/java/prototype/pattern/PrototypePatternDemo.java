package prototype.pattern;

/**
 * 原型模式（Prototype Pattern）是用于创建重复的对象，同时又能保证性能。
 * 一个对象需要在一个高代价的数据库操作之后被创建。我们可以缓存该对象，在下一个请求时返回它的克隆，在需要的时候更新数据库
 * 应用实例： 1、细胞分裂。 2、JAVA 中的 Object clone() 方法。
 */
public class PrototypePatternDemo {
    public static void main(String[] args) {
        ShapeCache.loadCache();

        Shape clonedShape = (Shape) ShapeCache.getShape("1");
        System.out.println("Shape:" + clonedShape.getType());

        Shape clonedShape2 = (Shape) ShapeCache.getShape("2");
        System.out.println("Shape:" + clonedShape2.getType());

        Shape clonedShape3 = (Shape) ShapeCache.getShape("3");
        System.out.println("Shape:" + clonedShape3.getType());
    }
}
/*
Shape:Circle
Shape:Square
Shape:Rectangle
 */