# css

Cascading Style Sheet 层叠级联样式表



## 什么是css

1. css选择器（重点+难点）

2. 美化网页（颜色，字体，宽度，高度，背景图片，文字、阴影、超链接、列表、渐变。。）

3. 盒子模型：

   ![image-20210113082847390](./resources/image/box-model.png)

4. 浮动：广告弹窗

5. 定位：广告弹窗

6. 网页动画（特效效果）

参见，菜鸟教程，**css教程**和**css3教程**



## 发展史

css 1.0

css 2.0 	div块+css，html和css结构分离的思想，网页变得简单，利于SEO（搜索引擎优化）

css 2.1 	浮动，定位

css 3.0	圆角，阴影，动画。。。浏览器兼容性 



**css的导入方式**

1. 行内样式

2. 内部样式

3. 外部样式（独立的.css文件）

   1. 链接式

      ```html
      	<link rel="stylesheet" href="style.css">
      ```

   2. 导入式，不建议！

      ```html
      <style>
        @import url("style.css")
      </style>
      ```

      

      

样式生效的优先级是**就近原则**：哪个样式离html元素的渲染最近，就以哪个为准！