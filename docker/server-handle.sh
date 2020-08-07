#! /bin/bash

rm -f ./core/server.go
# 生成server.go文件, 添加Router.Static("/admin", "./resource/dist")这个代码
touch ./core/server.go
filename="./core/server.go"
cat>"${filename}"<<EOF
package core

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")
	Router.Static("/admin", "./resource/dist")

	//InstallPlugs(Router)
	// end 插件描述

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Debug("server run success on ", address)

	fmt.Printf("欢迎使用 Gin-Vue-Admin默认自动化文档地址:http://127.0.0.1%s/swagger/index.html\n 默认前端文件运行地址:http://127.0.0.1:8888/admin\n", address)
	global.GVA_LOG.Error(s.ListenAndServe())
}
EOF

rm -f ./config.yaml
# 生成config.yaml文件, 用于docker-compose的使用
touch ./config.yaml
filename="./config.yaml"
cat>"${filename}"<<EOF
# Gin-Vue-Admin Global Configuration

# casbin configuration
casbin:
    model-path: './resource/rbac_model.conf'

# jwt configuration
jwt:
    signing-key: 'qmPlus'

# mysql connect configuration
mysql:
    username: root
    password: 'Aa@6447985'
    path: mysql
    db-name: 'qmPlus'
    config: 'charset=utf8&parseTime=True&loc=Local'
    max-idle-conns: 10
    max-open-conns: 10
    log-mode: true

#sqlite 配置
sqlite:
    path: db.db
    log-mode: true
    config: 'loc=Asia/Shanghai'

# oss configuration

# 请自行七牛申请对应的 公钥 私钥 bucket 和 域名地址
qiniu:
    access-key: '25j8dYBZ2wuiy0yhwShytjZDTX662b8xiFguwxzZ'
    secret-key: 'pgdbqEsf7ooZh7W3xokP833h3dZ_VecFXPDeG5JY'
    bucket: 'qm-plus-img'
    img-path: 'http://qmplusimg.henrongyi.top'

# redis configuration
redis:
    addr: redis:6379
    password: ''
    db: 0

# system configuration
system:
    use-multipoint: true
    env: 'public'  # Change to "develop" to skip authentication for development mode
    addr: 8888
    db-type: "mysql"  # support mysql/sqlite

# captcha configuration
captcha:
    key-long: 6
    img-width: 240
    img-height: 80

# logger configuration
log:
    prefix: '[GIN-VUE-ADMIN]'
    log-file: true
    stdout: 'DEBUG'
    file: 'DEBUG'
EOF

