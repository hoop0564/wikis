//
// Source code recreated from a .class file by IntelliJ IDEA
// (powered by FernFlower decompiler)
//

package netdemo;

import java.io.DataInputStream;
import java.io.DataOutputStream;
import java.io.IOException;
import java.net.ServerSocket;
import java.net.Socket;
import java.net.SocketTimeoutException;

public class GreetingServer extends Thread {
    private ServerSocket serverSocket;

    public GreetingServer(int var1) throws IOException {
        this.serverSocket = new ServerSocket(var1);
        this.serverSocket.setSoTimeout(10000);
    }

    public void run() {
        while(true) {
            try {
                System.out.println("waiting connections, port is: " + this.serverSocket.getLocalPort() + "...");
                Socket var1 = this.serverSocket.accept();
                System.out.println("remote host address: " + var1.getRemoteSocketAddress());
                DataInputStream var2 = new DataInputStream(var1.getInputStream());
                System.out.println(var2.readUTF());
                DataOutputStream var3 = new DataOutputStream(var1.getOutputStream());
                var3.writeUTF("thanks connections: " + var1.getLocalAddress() + "\n再见");
                var1.close();
            } catch (SocketTimeoutException var4) {
                System.out.println("socket time out!");
                return;
            } catch (IOException var5) {
                var5.printStackTrace();
            }
        }
    }

    public static void main(String[] var0) {
        int var1 = Integer.parseInt(var0[0]);

        try {
            GreetingServer var2 = new GreetingServer(var1);
            var2.start();
        } catch (IOException var3) {
            var3.printStackTrace();
        }

    }
}
