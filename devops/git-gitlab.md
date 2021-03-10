# git gitlab

- 国内的GitHub级别的仓库： [码云 Gitee](https://gitee.com/)
- git的动画教程：[learn-git-branching](https://oschina.gitee.io/learn-git-branching/)



## git推送到GitHub失败

推送到GitHub失败，但pull是OK的。报错信息是：

```
LibreSSL SSL_connect: SSL_ERROR_SYSCALL in connection to github.com:443
```

### 解决方案1：

```bash
git config --global --unset http.proxy
git config --global --unset https.proxy
```

此方案解决了gzc-private.git项目，对wikis.git无效

### 解决方案2:

```bash
git push origin master
```

然并卵。



可能还是github被墙调的原因



### 参考资料：

[SSL_connect: SSL_ERROR_SYSCALL in connection to github.com:443](https://stackoverflow.com/questions/48987512/ssl-connect-ssl-error-syscall-in-connection-to-github-com443)

