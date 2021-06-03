# WSL2

## 开启X图形界面

```bash
sudo apt update && apt upgrade

sudo apt install xrdp

sudo apt install xrdce4 xrdce4-goodies

# 修改bpp像素显示
sudo vim /etc/xrdp/xrdp.ini
# max_bpp=128
# xserverbpp=128

# 生成配置文件
echo xfce4-session > ~/.xsession

# 注释测试界面
sudo vim /etc/xrdp/startwm.sh
# 注释最后两行： test.. exec ..
# 添加一行：startxfce4

# 开启远程桌面
sudo /etc/init.d/xrdp start

# 查看ubuntu的ip，准备mstsc！
ip a

# ubuntu 中安装浏览器
sudo apt install firefox
```

