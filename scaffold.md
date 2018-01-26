# scaffold
自动化代码生成，加快开发效率

BaseDao(mapper)  
BaseBiz(service)
BaseController
BaseEntiy 基础数据封装(前后端契约mobile/pc/h5/，接口可复用性，输出，安全校验的契约，异常返回码等)

```
{
 status: "200",
 message: "ok",
 Data: {
 
 }
}

分页
{
 status: "200",
 message: "ok",
 Data: {
    
    rows:
    count:
    total: 
 }
}

异常
{
    status: "40101", 40303无权限等
    message: "token无效",
}

前端树 数组
{
    id:
    parentId:
}

前端树: 子节点树
{
    id:
    parentId:
    children:
}
```
Exception Interceptor
Context (服务拆分需要存储当前相关信息 threadlocal)
junit


脚手架目的

1. 模板代码,开发效率提升
2. 屏蔽技术细节，更关注业务
3. 统一相关规范，沟通效率高
4. 通用工具搭建，较少重复轮子


通用异常类封装

mvc拦截异常统一处理，最佳实践是建立自己的统一业务异常。业务异常通常定义在common scaffold中.



## todo
tracingId 日志异常跟踪 消息返回 便于排查
根据返回的消息id，快速排查当时发生的异常场景

