//
// Source code recreated from a .class file by IntelliJ IDEA
// (powered by FernFlower decompiler)
//

package netdemo;

import java.io.DataInputStream;
import java.io.DataOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.Socket;

public class GreetingClient {
    public GreetingClient() {
    }

    public static void main(String[] var0) {
        String var1 = var0[0];
        int var2 = Integer.parseInt(var0[1]);

        try {
            System.out.println("连接到主机：" + var1 + " ，端口号：" + var2);
            Socket var3 = new Socket(var1, var2);
            System.out.println("远程主机地址：" + var3.getRemoteSocketAddress());
            OutputStream var4 = var3.getOutputStream();
            DataOutputStream var5 = new DataOutputStream(var4);
            var5.writeUTF("Hello from " + var3.getLocalSocketAddress());
            InputStream var6 = var3.getInputStream();
            DataInputStream var7 = new DataInputStream(var6);
            System.out.println("服务器响应： " + var7.readUTF());
            var3.close();
        } catch (IOException var8) {
            var8.printStackTrace();
        }

    }
}
