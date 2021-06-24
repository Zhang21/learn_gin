# learn_gin

learn  go gin web framework


<br>
<br>


## 参考

- <https://github.com/gin-gonic/gin>
- <https://github.com/eddycjy/go-gin-example>


<br>
<br>


## 项目结构

```
├── conf        // 配置文件
├── middleware  // 应用中间件
├── models      // 应用数据库模型
├── pkg         // 第三方包
├── routers     // 路由逻辑处理
├── runtime     // 运行时数据
└── vendor      // 依赖
```


<br>
<br>


## Docker

```bash
# 构建
docker build -t zhang21/learn_gin .

# 运行
docker run -dit --name=learn_gin -p 8000:8000 zhang21/learn_gin
```
