# HTML

Hyper Text Markup Language 超文本标记语言，超文本表示它包括：文本、图片、音视频、动画等

浏览器的F12里，看到的左侧 `Elements` 表示html元素，右侧`Styles`表示css样式

html5和css3是未来的互联网方向！





## W3C

World Wide Web Consortium 万维网联盟，W3C标准包括：

- **结构**化标准语言：HTML、XML
- **表现**标准语言：CSS
- **行为**标准：DOM、ECMAScript



## HTML标签

```html
<!-- DOCTYPE：告诉浏览器，我们要使用什么规范 -->
<!DOCTYPE html>
<html lang="en">
  <!-- head标签代表网页头部 -->
  <head>
    <!-- meta描述性标签，用来描述网站的一些信息，譬如关键词keywords等 -->
    <meta charset="UTF-8">
    <meta name="keywords" content="编程,html,前端">
    <meta name="description" content="欢迎来学html的前端编程开发">
    <title>网页标题！</title>
  </head>
  <!-- body标签代表网页主体 -->
  <body>
    
  </body>
</html>
```



## 行内元素和块元素

块元素：无论内容多少，该元素独占一行，例如：p h1-h6

行原色：内容撑开长度，左右都是行内元素的可以，a strong em ...



## 页面结构分析

| 元素名  | 描述                                             |
| ------- | ------------------------------------------------ |
| header  | 标记头部区域的内容，用于页面或页面中的一块区域   |
| footer  | 标记脚部区域的内容，用于整个页面或页面的一块区域 |
| section | web页面的一块独立区域                            |
| article | 独立的文章内容                                   |
| aside   | 相关内容或应用，常用于侧边栏                     |
| nav     | 导航类辅助内容                                   |

