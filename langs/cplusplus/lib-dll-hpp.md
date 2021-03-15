# lib/dll

- 静态库的*.lib
  - 如果文件名是lib*.lib，则是静态链接的静态库，编译时会都编译进去，运行时无需dll
  - 如果文件名是*.lib，则是动态链接的静态库，编译时不会都编译进去，运行时需要dll
- 项目导入库时，用pragma的先后顺序


# hpp头文件

实质就是.cpp的实现代码混入.h头文件中，定义和实现都包含在同一文件。

- header plus plus的简写
- xx.cpp可以include此hpp，hpp会一起直接编译到xx.obj中，不再生成单独的obj，也不用加入到project这样中进行编译
- 如果被多个cpp引用，hpp中就不可定义全局变量和全局静态变量，否则会有符号重定义错误，可封装为类的静态方法
- 类之间不可循环调用，可用前向引用声明解决
- 采用hpp将大幅减少调用project的cpp文件数和编译次数，适合做公共的开源库
- [参考资料](https://blog.csdn.net/follow_blast/article/details/81706698)
