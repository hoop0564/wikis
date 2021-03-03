import java.io.*;
import java.io.InputStreamReader;

public class BRRead {
    public static void main(String[] args) throws IOException{
        String str;
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        System.out.println(("input character, quit with 'quit'"));
        do {
            str = br.readLine();
            System.out.println("input is:" + str);
        } while (!str.equals("quit"));
    }
}
