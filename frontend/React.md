# React

构建用户界面的JavaScript库，主要用于构建UI界面。

使用虚拟DOM，运用Diffing算法。



>逍遥游
>
>庄子
>
>北冥有鱼，其名为鲲。鲲之大，不知其几千里也。
>
>化而为鸟，其名为鹏。鹏之背，不知其几千里也。
>
>怒而飞，其翼若垂天之云。
>
>水击三千里，抟扶摇而上者九万里。



## 特点

- **1.声明式设计** −React采用声明范式，可以轻松描述应用。
- **2.高效** −React通过对DOM的模拟，最大限度地减少与DOM的交互。
- **3.灵活** −React可以与已知的库或框架很好地配合。
- **4.JSX** − JSX 是 JavaScript 语法的扩展。React 开发不一定使用 JSX ，但我们建议使用它。
- **5.组件** − 通过 React 构建组件，使得代码更加容易得到复用，能够很好的应用在大项目的开发中。
- **6.单向响应的数据流** − React 实现了单向响应的数据流，从而减少了重复代码，这也是它为什么比传统数据绑定更简单。



## React库

| 文件名                   | 作用                                                         | 说明 |
| ------------------------ | ------------------------------------------------------------ | ---- |
| babel.min.js             | Babel 可以将 ES6 代码转为 ES5 代码，这样就能在目前不支持 ES6 浏览器上执行 React 代码。Babel 内嵌了对 JSX 的支持。通过将 Babel 和 babel-sublime 包（package）一同使用可以让源码的语法渲染上升到一个全新的水平。 |      |
| react.development.js     | React 的核心库-开发版                                        |      |
| react-dom.development.js | 提供与 DOM 相关的功能                                        |      |
| prop-type.js             |                                                              |      |
| react.min.js             | React 的核心库                                               |      |

**注意:** 在浏览器中使用 Babel 来编译 JSX 效率是非常低的。



**React 只会更新必要的部分**

值得注意的是 React DOM 首先会比较元素内容先后的不同，而在渲染过程中只会更新改变了的部分。



# React JSX

React 使用 JSX 来替代常规的 JavaScript。

JSX 是一个看起来很像 XML 的 JavaScript 语法扩展。

- JSX 执行更快，因为它在编译为 JavaScript 代码后进行了优化。

- 它是类型安全的，在编译过程中就能发现错误。

- 使用 JSX 编写模板更加简单快速。

  

## React组件

使用函数定义了一个组件：

```jsx
function HelloMessage(props) {
    return <h1>Hello World!</h1>;
}
```

使用 ES6 class 来定义一个组件:

```jsx
class Welcome extends React.Component {
  render() {
    return <h1>Hello World!</h1>;
  }
}
```

**const element = <HelloMessage />** 为用户自定义的组件。



复合组件

创建多个组件来合成一个组件，即把组件的不同功能点进行分离。

```jsx
unction Name(props) {
    return <h1>网站名称：{props.name}</h1>;
}
function Url(props) {
    return <h1>网站地址：{props.url}</h1>;
}
function Nickname(props) {
    return <h1>网站小名：{props.nickname}</h1>;
}
function App() {
    return (
    <div>
        <Name name="菜鸟教程" />
        <Url url="http://www.runoob.com" />
        <Nickname nickname="Runoob" />
    </div>
    );
}
 
ReactDOM.render(
     <App />,
    document.getElementById('example')
);
```



## React State(状态)

React 把组件看成是一个状态机（State Machines）。通过与用户的交互，实现不同状态，然后渲染 UI，让用户界面和数据保持一致。

React 里，只需更新组件的 state，然后根据新的 state 重新渲染用户界面（不要操作 DOM）。



## 注意事项

- *原生 HTML 元素名以小写字母开头，而自定义的 React 类名以大写字母开头，比如 HelloMessage 不能写成 helloMessage。除此之外还需要注意组件类只能包含一个顶层标签，否则也会报错。*
- *在添加属性时， class 属性需要写成 className ，for 属性需要写成 htmlFor ，这是因为 class 和 for 是 JavaScript 的保留字。*

- 添加自定义属性需要使用 **data-** 前缀。
- 在 JSX 中不能使用 **if else** 语句，但可以使用 **conditional (三元运算)** 表达式来替代。
- JSX 允许在模板中插入数组，数组会自动展开所有成员
-  React DOM 首先会比较元素内容先后的不同，而在渲染过程中只会更新改变了的部分。