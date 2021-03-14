package multithread;

public class TestLambda {
    // 静态内部类
    static class Like2 implements ILike {

        @Override
        public void lambda() {
            System.out.println("static internal class like2");
        }
    }

    public static void main(String[] args) {
        ILike like = new Like();
        like.lambda();

        like = new Like2();
        like.lambda();

        // 局部内部类
        class Like3 implements ILike {

            @Override
            public void lambda() {
                System.out.println("local class like3");
            }
        }

        like = new Like3();
        like.lambda();

        like = new ILike() {
            @Override
            public void lambda() {
                System.out.println("匿名内部类：like4");
            }
        };
        like.lambda();

        // 使用lambda实现函数式接口
        // 避免内部定义类实现过多
        like = () -> {
            System.out.println("lambda like5");
        };
        like.lambda();
    }
}

// 函数式接口：只有一个方法的接口
interface ILike {
    void lambda();
}

// 实现类
class Like implements ILike {

    @Override
    public void lambda() {
        System.out.println("outer class like");
    }
}
