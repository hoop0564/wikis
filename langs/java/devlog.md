# 集成开发日志



## JPA

- JPA注解表名称和字段名称爆红：`cannot resolve table/column 'xxx'`

  ```java
  @Table(name = "RTC_Machine") 
  public class RTC_Machine {
      @Column(name = "Name")
      private String name;
  }
  ```

  > 配置下DataSource：[idea中Entity实体中报错：cannot resolve column/table/...解决办法。](https://blog.csdn.net/flyingshadower/article/details/81974912?utm_medium=distribute.pc_relevant.none-task-blog-baidujs_title-1&spm=1001.2101.3001.4242)

```
Hibernate: alter table rtc_machine add column is_actived boolean not null
```



```properties
spring.jpa.hibernate.naming.implicit-strategy=org.hibernate.boot.model.naming.ImplicitNamingStrategyLegacyJpaImpl
spring.jpa.hibernate.naming.physical-strategy=org.hibernate.boot.model.naming.PhysicalNamingStrategyStandardImpl
```



```sql
create table RTC_Machine (Id int8 not null, IsActived boolean not null, Name varchar(255), primary key (Id))
```



- 表名和字段名大小写敏感

  1. 代码中做转义

  ```java
  @Table(name = "\"RTC_Machine\"") 
  ```

  2. 