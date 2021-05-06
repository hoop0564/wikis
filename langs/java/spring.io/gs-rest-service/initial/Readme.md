# Readme

如果使用Maven命令运行该应用程序：
```bash
# linux
./mvnw spring-boot:run

# windows 到pom.xml同级的目录下执行命令，且mvn命令需要在PATH路径中
mvn spring-boot:run
```

使用此命令来构建生成 `jar` 文件：
```bash
# linux
./mvnw clean package

# windows 到pom.xml同级的目录下执行命令
mvn clean package
```

运行JAR文件：
```bash
cd /d E:\repos\java\spring-boot-examples\spring-boot-helloWorld
java -jar target\spring-boot-helloworld-1.0.0-SNAPSHOT.jar
```

## 测试服务
现在该服务已启动，请访问：
```http request
http://localhost:8080/hello
```
