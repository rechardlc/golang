关于go包相关
1. 一个文件夹是一个包，包内的文件的.go文件第一行都应该是package 包名(package utils)
2. 同一个包中不可以有相同名字~无路是函数名还是变量名甚至定义类型等等~都不可以同名
3. 引用包名可以取别名，如：import (utils "path/utils"")
4. 同一个包下的所有的.go文件~无论是变量还是函数等等都可以直接使用~无论做其他操作
5. 若要编译一个可执行文件，一个项目只能存在一个main包，即main函数所在的包