const express = require('express')
const { buildSchema } = require('graphql')
const graphqlHttp = require('express-graphql')

// 定义schema，查询和类型
const schema = buildSchema(`
     input AccountInput {
         name: String
         age: Int
         sex: String
         department: String
     }
     type Account {
         name: String
         age: Int
         sex: String
         department: String
     }
     type Mutation {
         createAccount(input: AccountInput): Account
         updateAccount(id: ID!, input: AccountInput): Account
     }
     # Query至少有一个，在graphQL中定义的
     type Query {
         accounts: [Account]
     }
`)

let fakeDb = {};

// 定义查询对应的处理器resolver
const root = {
    accounts() {
        return Object.values(fakeDb);
    },
    createAccount({ input }) {
        // 相当于数据库的保存
        fakeDb[input.name] = input;
        return fakeDb[input.name];
    },
    // id就认为是name
    updateAccount({ id, input }) {
        // assign相当于把第二个和第三个的属性copy到{}上来
        const updatedAccount = Object.assign({}, fakeDb[id], input);
        fakeDb[id] = updatedAccount;
        //返回保存结果
        return updatedAccount;
    }
}

const app = express();

app.use('/graphql', graphqlHttp.graphqlHTTP({
    schema: schema,
    rootValue: root,
    graphiql: true
}))

app.listen(3000)
console.log(`please go to: http://localhost:3000/graphql`)

/**
 * 输入：
 * mutation {
  createAccount(input: {
    name: "hello"
    sex: "male"
    age: 23
    department: "演员"
  }) {
    name
    age
    sex
    department
  }
}
* 输出
{
  "data": {
    "createAccount": {
      "name": "hello",
      "age": 23,
      "sex": "male",
      "department": "演员"
    }
  }
}

* 输入
query {
  accounts {
    name
  }
}
* 输出
{
  "data": {
    "accounts": [
      {
        "name": "world"
      },
      {
        "name": "world2"
      }
    ]
  }
}

* 输入
mutation {
  updateAccount(id: "world", input: {
    age: 101
  }) {
    age
  }
}
* 输出
{
  "data": {
    "updateAccount": {
      "age": 101
    }
  }
}
 */