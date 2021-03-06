# net demo

## server
```bash
cd src/netdemo
javac GreetingServer.java

cd ..
java netdemo.GreetingServer 6066
```
## client
```bash
cd src/netdemo
javac GreetingClient.java

cd ..
java netdemo.GreetingClient localhost 6066
```