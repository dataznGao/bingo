# bingo

bingo是一套实现go语言静态故障注入的工具集，具有简单易用，易于理解的特点。截至目前，该工具集一共实现了如下种故障注入模型：

```go
ValueFault   值故障
NullFault    空指针故障
ExceptionShortcircuitFault  异常短路故障    
ExceptionUncaughtFault    异常未捕获故障
AttributeShadowedFault    参数屏蔽故障
SwitchMissDefaultFault    switch-case default缺失故障
ConditionBorderFault    条件边界故障
ConditionInversedFault    条件反转故障
SyncFault                同步故障
AttributeReversoFault    缩放故障
其核心原理是利用词法/语法分析进行代码静态分析，找出潜在的故障注入点，进而注入特定的错误。
```

快速上手：

1. 创建env
```go
env := env.CreateFaultEnv("${inputPath}", "${outputPath}")
```
2. 注入故障点
```go
env.ConditionInversedFault("${locatePattern}")
```
    引入faultPoint的概念，faultPoint即故障点，采用一种文本进行表达
```go
"util(1/5).myStruct(1/3).myFunc(1/2).myVariable | *(3/4).*.*.*"
```
a. 必须是四段式表达，对应为包.结构体.函数.变量
b. "|" 表示或者，不存在与（&）关系，因为一个函数一般不会既在a结构体又在b结构体，
    如果是同名函数，只要再配置一个故障点即可
c.  "*"表示上层下的全部，比如包下的全部结构体，结构体下的全部方法，如果不填写则默认全部注入
d. 括号内的分数表示该故障点生效的概率
3. 创建factory并运行
```go
f := code_drill.FaultPerformerFactory{}
err := f.SetEnv(env).Run()
```
