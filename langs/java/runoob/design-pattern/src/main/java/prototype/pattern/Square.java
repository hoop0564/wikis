package prototype.pattern;

public class Square extends Shape{
    public Square() {
        type = "Square";
    }

    @Override
    public void draw() {
        System.out.println("inside Square::draw() method.");
    }
}
