# XML

- 在 XML 中，您应该尽量避免使用属性。如果信息感觉起来很像数据，那么请使用元素。

- 元数据（**有关数据的数据**）应当存储为属性，而数据本身应当存储为元素。

  ```xml
  <messages>
  <note id="501">
  <to>Tove</to>
  <from>Jani</from>
  <heading>Reminder</heading>
  <body>Don't forget me this weekend!</body>
  </note>
  <note id="502">
  <to>Jani</to>
  <from>Tove</from>
  <heading>Re: Reminder</heading>
  <body>I will not</body>
  </note>
  </messages>
  ```

- XML中请用**实体引用**来代替 "<" 字符，

  - `&lt; `表示 <
  - `&gt;`表示 >
  - `&amp;`表示 &
  - `&apos;` 表示 '
  - `&quot;`表示 "

- XML 以 LF 存储换行
- 使元素名称具有描述性，推荐使用下划线的名称：`<book_title>`

