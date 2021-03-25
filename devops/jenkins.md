# Jenkins



```bash
docker run -d -u root \
-p 8080:8080 \
-v /var/run/docker.sock:/var/run/docker.sock \
-v ${which docker}:/bin/docker \
-v /var/jenkins_home:/var/jenkins_home \
jenkins
```

