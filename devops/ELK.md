# ELK



Beats：轻量级的数据采集器，其中的filebeat和metricbeat最重要

```bash
./filebeat -e -c itcast.yml
```



x server：连接linux系统的windows的UI工具



## Kibana

数据分析的可视化平台

### 数据探索

1. 路径：

   >  导航栏中：Management -> Index Patterns -> Create index pattern：

2. 会自动显示出已有的ES中的索引库。或手动输入做查找：

   > metricbeat-*

3. Time Filter field name 选择 @timestamp，确认创建 【Create index pattern】

4. 导航栏中 Discover，可查看ES中的数据

