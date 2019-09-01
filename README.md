# QMPlus gin+vue开源快速项目模板

## 技术选型
    1.后端采用golang框架gin，快速搭建基础restful风格API
    2.前端项目采用VUE框架，构建基础页面
    3.数据库采用Mysql，可能会引用redis作为缓存数据库使用（待定）
    4.使用swagger构建自动化文档
    5.使用fsnotify和viper实现json格式配置文件
    6.使用logrus实现日志记录
    7.使用gorm实现对数据库的基本操作
## TODO
    1.基本用户注册登录
    2.用户等基础数据CURD
    3.调用des实现数据加密
    4.实现基于jwt的权限管理
    5.实现基于七牛云的文件上传操作（需提前注册七牛云账号）
    6...看项目进度想到什么做什么,主要目的是方便各位快速接私活，完成项目基础功能
## 使用说明
    1.golang api server 基于go.mod 如果golang版本低于1.11 请自行升级golang版本
    2.支持go.mod的golang版本在运行go list 和 编译之前都会自动下载所需要的依赖包
    3.go server建议使用goland运行 减少出错可能性
    4.前端项目node建议高于V8.6.0
    5.到前端项目目录下运行 npm i 安装所需依赖
    6.依赖安装完成直接运行 npm run dev即可启动项目
## 个人博客
    http://www.henrongyi.top，内有前端框架教学视频，GOLANG基础入门视频正在筹备中。
    如果觉得项目对您有所帮助可以添加我的个人微信:shouzi_1994,欢迎您提出宝贵的需求。
    
 
