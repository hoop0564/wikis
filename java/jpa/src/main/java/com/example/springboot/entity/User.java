package com.example.springboot.entity;

import javax.persistence.*;

// 配置映射关系，使用jpa注解
@Entity // 这是一个bean实体类，和数据表映射的类，
@Table(name = "tbl_user") // 指定和哪个表对应；如果省略，表名默认为user（User 小写）
public class User {
    @Id //这是一个主键
    @GeneratedValue(strategy = GenerationType.IDENTITY) // 自增主键
    private Integer id;

    @Column(name="last_name", length = 50) // 这是和表对应的一个列
    private String lastName;
    @Column // 省略默认列名就是属性名
    private String email;

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public String getLastName() {
        return lastName;
    }

    public void setLastName(String lastName) {
        this.lastName = lastName;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }
}
