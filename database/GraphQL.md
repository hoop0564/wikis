# GraphQL

- 需要啥，客户端请求中明确定义
- 服务端需要预先实现这些schema



## 基本参数类型

String，Int，Float，Boolean，ID，可以在schema声明的时候直接使用

- 注意ID本质上是String类型，类型为ID的不能有重复值！
- [类型]代表数组，例如：[Int] 代表整形数组



## 参数传递

- 和js一样，小括号内定义形参，参数需要定义类型

- !（叹号）代表此参数必须传递。否则为可选参数。

  ```javascript
  type Query {
  	// funcName(paramName: paramType!): returnType
  	rollDice(numDice: Int!, numSides: Int): [Int]
  }
  ```




## 自定义参数类型

通常用来描述要获取的资源的属性：

```javascript
type Account {
	name: String
  age: Int
  sex: String
  department: String
  # 下面是个函数，city为输入参，返回Int
  salary(city: String): Int
}
type Query {
  account(name: String): Account
}
```

