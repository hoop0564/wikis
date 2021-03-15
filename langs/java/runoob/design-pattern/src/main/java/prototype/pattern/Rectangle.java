package prototype.pattern;

public class Rectangle extends Shape{
    public Rectangle() {
        type = "Rectangle";
    }

    @Override
    public void draw() {
        System.out.println("inside Rectangle::draw() method.");
    }
}
