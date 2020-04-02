﻿
<div align=center>
<img src="http://qmplusimg.henrongyi.top/gvalogo.jpg" width=300" height="300" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/vue-2.6.10-brightgreen"/>
<img src="https://img.shields.io/badge/element--ui-2.12.0-green"/>
<img src="https://img.shields.io/badge/golang-1.12-blue"/>
<img src="https://img.shields.io/badge/gin-1.4.0-lightBlue"/>
<img src="https://img.shields.io/badge/gorm-1.9.10-red"/>
</div>

## 开发文档
[在线文档](http://doc.henrongyi.top/) [http://doc.henrongyi.top/](http://doc.henrongyi.top/)

本模板使用前端ui框架为 element-ui https://element.eleme.cn/#/zh-CN 前端组件可查看elementUi文档使用

## 基本介绍
>GIN-VUE-ADMIN是一个基于vue和gin开发的全栈前后端分离的后台管理系统，拥有jwt鉴权，动态路由，动态菜单，casbin鉴权，表单生成器，代码生成器等功能，提供了多种示例文件，让大家把更多时间专注在业务开发上。

## 技术选型
1. 后端采用golang框架gin，快速搭建基础restful风格API
2. 前端项目采用VUE框架，构建基础页面
3. 数据库采用Mysql(5.6.44)版本不同可能会导致SQL导入失败
4. 使用redis实现记录当前活跃用户的jwt令牌并实现多点登录限制
5. 使用swagger构建自动化文档
6. 使用fsnotify和viper实现json格式配置文件
7. 使用logrus实现日志记录
8. 使用gorm实现对数据库的基本操作

## 项目目录

```
    ├─QMPlusServer  	（后端文件夹）
    │  ├─cmd     	（启动文件）
    │  ├─config    	（配置包）
    │  ├─controller  	（api和servers存放位置）
    │  ├─db       	（数据库脚本）
    │  ├─docs  	（swagger文档目录）
    │  ├─init      	（初始化路由 数据库 日志等）
    │  ├─log     	（日志存放地址）
    │  ├─middleware   	（中间件）
    │  ├─model            	（结构体层）
    │  ├─router          	（路层）
    │  ├─static	（静态文件--配置文件 casbin模型等）
    │  ├─tools  (后端工具包)
    │  └─tpl		（自动化代码模板）
    └─QMPlusVuePage	（前端文件）
        ├─public	（发布模板）
        └─src
            ├─api	（向后台发送ajax的封装层）
            ├─assets	（静态文件）
            ├─components（组件）
            ├─router	（前端路由）
            ├─store	（vuex 状态管理仓）
            ├─style	（通用样式文件）
            ├─utils	（前端工具库）
            └─view	（前端页面）

```

## 主要功能
1. 权限管理：基于jwt和casbin实现的权限管理 

2. 文件上传下载：实现基于七牛云的文件上传操作（需提前注册七牛云账号） （为了方便大家测试，我公开了自己的七牛测试号的各种重要token，恳请大家不要乱传东西）

3. 分页封装：等装了分页方法，实现分页接口并且复制粘贴就可使用分页，前端分页mixin封装 分页方法调用mixins即可 

4. 用户管理：系统管理员分配用户角色和角色权限。

5. 角色管理：创建权限控制的主要对象，可以给角色分配不同api权限和菜单权限。

6. 菜单管理：实现用户动态菜单配置，实现不同角色不同菜单。

7. api管理：不同用户可调用的api接口的权限不同。

8. 配置管理：配置文件可前台修改（测试环境不开放此功能）

9. 富文本编辑器：富文本编辑器，MarkDown编辑器功能嵌入 

10. 条件搜索：增加条件搜索示例 前端文件参考src\view\superAdmin\api\api.vue 后台文件参考 model\dnModel\api.go √

11. 多点登录限制：体验需要再 static\config中 把 system中的useMultipoint 修改为 true(需要自行配置redis和config中的redis参数)(测试阶段，有bug请及时反馈)

12. 分片长传：提供文件分片上传和大文件分片上传功能示例 

13. 表单生成器：表单生成器借助 [@form-generator](https://github.com/JakHuang/form-generator)

14. 代码生成器：后台基础逻辑以及简单curd的代码生成器 

## 计划任务
1. 导入，导出Excel

2. Echart图表支持

3. 工作流，任务交接功能开发

4. 单独前端使用模式以及数据模拟
## 使用说明
1. golang api server 基于go.mod 如果golang版本低于1.11 请自行升级golang版本

2. 支持go.mod的golang版本在运行go list 和 编译之前都会自动下载所需要的依赖包

3. go server建议使用goland运行 减少出错可能性

4. 前端项目node建议高于V8.6.0

5. 到前端项目目录下运行 npm i 安装所需依赖

6. 依赖安装完成直接运行 npm run serve即可启动项目

7. 如果要使用swagger自动化文档 首先需要安装 swagger

````
go get -u github.com/swaggo/swag/cmd/swag
````

由于国内没法安装到X包下面的东西 如果可以翻墙 上面的命令就可以让你安心使用swagger了
如果没有翻墙的办法那就先装一下 gopm

````
go get -v -u github.com/gpmgo/gopm
````

此时你就可以使用 gopm了
这时候执行

````
gopm get -g -v github.com/swaggo/swag/cmd/swag
````

等待安装完成以后
到我们GOPATH下面的/src/github.com/swaggo/swag/cmd/swag路径
执行

````
go install
````

安装完成过后在项目目录下运行

````
swag init
````

项目文件夹下面会有 doc文件夹出现

这时候登录 localhost:8888/swagger/index.html

就可以看到 swagger文档啦

## 团队博客
    https://blog.henrongyi.top，内有前端框架教学视频，GOLANG基础入门视频正在筹备中。
    如果觉得项目对您有所帮助可以添加我的个人微信:shouzi_1994,欢迎您提出宝贵的需求。
    
## docker镜像
   感谢 [@chenlinzhong](https://github.com/chenlinzhong)提供docker镜像
   
      #启动容器
      docker run -itd --net=host --name=go_container shareclz/go_node /bin/bash;
      
      #进入容器
      docker exec -it go_container /bin/bash;
      git clone https://github.com/piexlmax/gin-vue-admin.git /data1/www/htdocs/go/admin;
      
      #启动前端
      cd /data1/www/htdocs/go/admin/QMPlusVuePage;
      cnpm i ;
      npm run serve;
      
      #修改数据库配置
      vi /data1/www/htdocs/go/admin/QMPlusServer/static/dbconfig/config.json;
      
      #启动后端
      cd /data1/www/htdocs/go/admin/QMPlusServer;
      go run main.go;
      
## 一点建议
    各位在clone项目以后，把db文件导入自己创建的库后，最好前往七牛云申请自己的空间地址，
    替换掉项目中的七牛云公钥，私钥，仓名和默认url地址，以免发生测试文件数据错乱
    
## 测试环境地址

测试环境:[http://qmplus.henrongyi.top/](http://qmplus.henrongyi.top/)
 
账号/密码: admin/123456

## 环境搭建教学视频

腾讯视频：https://v.qq.com/x/page/e3008xjxqtu.html    (等待最新视频录制)
    
## 模板使用教学及展示视频

腾讯视频：https://v.qq.com/x/page/c3008y2ukba.html    (等待最新视频录制)

## 联系方式

|  奇淼   | krank666  |qq群|
|  :---:  |  :---: | :---: |
|  <img src="http://qmplusimg.henrongyi.top/jjz.jpg" width="180"/>  |  <img src="http://qmplusimg.henrongyi.top/yx.jpg" width="180"/> | <img src="http://qmplusimg.henrongyi.top/qq.jpg" width="180"/> |

<div align=center>
<h3>qq交流群:622360840</h3>
<h3>微信交流群可以添加任意一位开发者备注"加入gin-vue-admin交流群"</h3>
</div>

## 开发者(贡献者)列表

|  开发者   | 功能  | 姓名  |
|  ----  | ----  | ----  |
| [@piexlmax](https://github.com/piexlmax)  | 项目发起者 | 蒋\*兆 |
| [@krank666](https://github.com/krank666)  | 前端联合作者 | 尹\* |
| [@1319612909](https://github.com/1319612909)  | 前端css优化 |  杜\*兰 |
| [@granty1](https://github.com/granty1)  | 代码积极贡献者 | 印\*林 |

## 更新日志

|  日期   | 日志  |
|  :---:  | --- |
|2020/01/07| 角色增加数据资源功能 增加数据资源关联返回 演示环境代码已同步 开启了多点登录拦截 可能会被其他人挤掉 |
|2020/01/13| 增加了配置管理功能 此功能不发表至测试环境 待保护机制以及服务重启机制发开完成后才会发表值测试环境 请自行clone且导入sql体验 |
|2020/02/21| 修改了casbin的自定义鉴权方法，使其完全支持RESTFUL的/:params以及?query= 的接口模式 |
|2020/03/17| 增加了验证码功能 使用了 [@dchest/captcha](https://github.com/dchest/captcha)库 |
|2020/03/30| 代码生成器开发完成 表单生成器开发完成 使用了[@form-generator](https://github.com/JakHuang/form-generator) 库 |
|2020/04/01| 增加前端历史页签功能，增加（修改）条件查询示例，前端背景色调修改为白色 如不需要此功能可以在 view\/layout\/index\/   屏蔽HistoryComponent 背景色调 为本页260行 &.el-main 的background |

## golang基础教学视频录制中...
地址:https://space.bilibili.com/322210472/channel/detail?cid=108884


## 捐赠
如果你想请团队喝可乐

|  支付宝   | 微信  |
|  :---:  | :---: |
| ![markdown](http://qmplusimg.henrongyi.top/zfb.png "支付宝") |  ![markdown](http://qmplusimg.henrongyi.top/wxzf.png "微信") |


## 捐赠列表

|  捐赠者   | 金额  |
|  :---:  | :---: |
| 老**途 |  100￥ |
| y*g |  10￥ |
| *波 |  50￥ |
| *雄 |  15￥ |
