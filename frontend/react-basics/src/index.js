import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import './App.css'

const time = new Date().toLocaleTimeString()
const str = `当前时间是：`
const element = (
  <div>
    <h1>hello header1</h1>
    <h2>{str + time}</h2>
  </div>
)

const element3 = (
  <div>
    <p><span>横着</span></p>
    <p><span>竖着</span></p>
  </div>
)
const man = '正常'
const element2 = (
  <div>
    <h1>today it isolated?</h1>
    <h2>{man === '发热' ? <button>隔离</button> : element3}</h2>
  </div>
)

const color = 'bgRed'
const logo = 'https://www.baidu.com/s?wd=%E4%BB%8A%E6%97%A5%E6%96%B0%E9%B2%9C%E4%BA%8B&tn=SE_Pclogo_6ysd4c7a&sa=ire_dl_gh_logo&rsv_dl=igh_logo_pc'
const element5 = (
  <div className={color}>
    <img src={logo}></img>
    红色的背景颜色
  </div>
)

ReactDOM.render(
  element5,
  document.getElementById('root'),
)
// ReactDOM.render(
//   <React.StrictMode>
//     <App />
//   </React.StrictMode>,
//   document.getElementById('root')
// );

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
