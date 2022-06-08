## 评价中心

### 目录结构
- app
    - console   --------------------**定时任务**
      - schedule.go 
    - controllers   -----------------**控制器**
      - appraise.go
    - services  --------------------**逻辑层**  
      - appraise.go
    - models    ---------------------**数据层**
        - dao   --------------------**结构体**
          - class_center.go
          - wm_appraise.go  
        - appraise.go
- pkg
  - cache   ----------------------**缓存**
    - redis.go
  - database    -------------------**数据库**
    - mysql.go
  - errcode --------------------**错误码**
    - code.go
    - desc.go
  - support --------------------**通用方法**
    - helper.go
    - response.go
- route ---------------------------**路由**
  - api.go
- vendor
- .env.example -------------------**环境配置**
- .gitignore
- go.mod
- go.sum
- main.go ------------------------**入口文件**

### 使用和部署
```
1. git clone git@gitlab.weimiaocaishang.com:weimiao/wm-appraise-center-go.git
```
```
2. cd wm-appraise-center-go && cp .env.example .env
```
```
3. go run main.go
```