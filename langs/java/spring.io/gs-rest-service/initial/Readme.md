# Readme

如果使用Maven命令运行该应用程序：
```bash
./mvnw spring-boot:run
```

使用此命令来构建JAR文件：
```bash
./mvnw clean package
```

运行JAR文件：
```bash
java -jar target/rest-service-0.0.1-SNAPSHOT.jar
```

## 测试服务
现在该服务已启动，请访问：
```http request
http://localhost:8080/greeting
```
