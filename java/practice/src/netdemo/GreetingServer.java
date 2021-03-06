package netdemo;

import java.io.DataInputStream;
import java.io.DataOutputStream;
import java.io.IOException;
import java.net.ServerSocket;
import java.net.Socket;
import java.net.SocketTimeoutException;

public class GreetingServer extends Thread {
    private ServerSocket serverSocket;

    public GreetingServer(int port) throws IOException {
        serverSocket = new ServerSocket(port);
        serverSocket.setSoTimeout(10000);
    }

    public void run() {
        while (true) {
            try {
                System.out.println("waiting connections, port is: " + serverSocket.getLocalPort() + "...");
                Socket server = serverSocket.accept();
                System.out.println("remote host address: " + server.getRemoteSocketAddress());
                DataInputStream in = new DataInputStream(server.getInputStream());
                System.out.println(in.readUTF());

                DataOutputStream out = new DataOutputStream(server.getOutputStream());
                out.writeUTF("thanks connections: "+ server.getLocalAddress()+"\n再见");
                server.close();
            } catch (SocketTimeoutException s) {
                System.out.println("socket time out!");
                break;
            }catch (IOException e) {
                e.printStackTrace();
            }
        }
    }

    public static void main(String[] args) {
        int port = Integer.parseInt(args[0]);
        try
        {
            Thread t = new GreetingServer(port);
            t.start();
        } catch (IOException e)
        {
            e.printStackTrace();
        }
    }

}
