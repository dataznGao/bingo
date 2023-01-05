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
