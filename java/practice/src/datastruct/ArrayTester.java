package datastruct;

import java.util.Arrays;

public class ArrayTester {

    public static void main(String[] args) {
        ArrayTester t = new ArrayTester();
//        t.arrayDemo();
        t.sparseArray();
    }

    /**
     * 排序和查找
     */
    void arrayDemo() {

        int[] a = {11, 3, 6, 4, 2};
        System.out.println(Arrays.toString(a));

        // 排序
        Arrays.sort(a);
        System.out.println(Arrays.toString(a));
        // 二分查找已排序好的数组
        System.out.println(Arrays.binarySearch(a, 3));
    }

    /**
     * 稀疏数组
     * 使用稀疏数组来保存一个具有稀疏特性的数组中的有效数据
     */
    void sparseArray() {
        // 创建二位数组 0：没有旗子，1：黑棋，2：白旗
        int[][] array1 = new int[11][11];
        array1[1][2] = 1;
        array1[2][3] = 1;
        System.out.println("原始数组：");
        for (int[] ints : array1) {
            for (int anInt : ints) {
                System.out.print(anInt + "\t");
            }
            System.out.println();
        }

        // 转换为稀疏数组保存
        // 获取有效值的个数
        int sum = 0;
        for (int i = 0; i < 11; i++) {
            for (int j = 0; j < 11; j++) {
                if (array1[i][j] != 0) {
                    sum++;
                }
            }
        }
        System.out.println("有效值的个数：" + sum);

        // 创建一个稀疏数组的数组
        // 3是用来存位置（row/line)和值（value)
        int[][] arraySparse = new int[sum + 1][3];
        // 第一个元素是用来存数组的总行列数和实际的有效数值数
        arraySparse[0][0] = 11;
        arraySparse[0][1] = 11;
        arraySparse[0][2] = sum;

        // 遍历二维数组，将非零的值，存放到稀疏数组中
        int count = 0;
        for (int i = 0; i < array1.length; i++) {
            for (int j = 0; j < array1[i].length; j++) {
                if (array1[i][j] != 0) {
                    count++;
                    arraySparse[count][0] = i;
                    arraySparse[count][1] = j;
                    arraySparse[count][2] = array1[i][j];
                }
            }
        }

        // 输出稀疏数组
        System.out.println("稀疏数组：");
        for (int[] ints : arraySparse) {
            System.out.println(ints[0] + "\t"
                    + ints[1] + "\t"
                    + ints[2] + "\t");
        }

        // 读取稀疏数组
        System.out.println("还原：");
        int row = arraySparse[0][0];
        int line = arraySparse[0][1];
        int[][] arrayRecover = new int[row][line];
        for (int i = 1; i < arraySparse.length; i++) {

            row = arraySparse[i][0];
            line = arraySparse[i][1];
            int value = arraySparse[i][2];
            arrayRecover[row][line] = value;
        }

        // 输出还原的数组
        System.out.println("输出还原的数组:");
        for (int[] ints : arrayRecover) {
            for (int anInt : ints) {
                System.out.print(anInt+"\t");
            }
            System.out.println();
        }
    }
}
