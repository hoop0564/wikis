const express = require('express')
const { buildSchema } = require('graphql')
const graphqlHttp = require('express-graphql')

// 定义schema，查询和类型
const schema = buildSchema(`
    type Account {
        name: String
        age: Int
        sex: String
        department: String
        salary(city: String): Int
    }
    #这里是注释：Query的Q要大写
    type Query {
        getClassMates(classNo: Int!): [String]
        account(username: String): Account
    }
`)

// 定义查询对应的处理器resolver（模拟简单的数据库访问）
const root = {
    getClassMates({ classNo }) {
        const obj = {
            31: ['张三', '李四', '王五'],
            61: ['张大三', '李大四', '王大五']
        }
        return obj[classNo];
    },
    account({ username }) {
        const name = username
        const sex = "male"
        const age = 18
        const department = "IT"
        const salary = ({ city }) => {
            if (['北京', '上海', '广州', '深圳'].includes(city)) {
                return 10000
            }
            return 6000
        }
        return {
            name,
            sex,
            age,
            department,
            salary,
        }
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
 query {
  getClassMates(classNo: 31)
}
* 输出
{
  "data": {
    "getClassMates": [
      "张三",
      "李四",
      "王五"
    ]
  }
}
 */