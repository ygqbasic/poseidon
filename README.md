## gardens

# 简介
基于Beego开发的易用、易扩展、界面友好的轻量级功能权限管理系统。前端框架基于AdminLTE2进行资源整合，包含了多款优秀的插件，是笔者对多年后台管理系统开发经验精华的萃取。
本系统非常适合进行后台管理系统开发，统一的代码和交互给二次开发带来极大的方便，在没有前端工程师参与的情况下就可以进行快速的模块式开发，并保证用户使用的友好性和易用性。系统里整合了众多优秀的资源，在此感谢各位大神的无私奉献。
# 框架说明
  框架来源于 https://github.com/lhtzbj12/sdrms

本文博客
# 特点
1. 分页列表页面的搜索条件、搜索面板、PageSize、当前页数、显示/隐藏列在变化时自动保存，页面刷新后、重新进入时，这些状态依然保持；
2. TreeTabe列表节点展开/收缩状态、滚动条位置时自动保存，页面刷新后、重新进入时，这些状态依然保持；
3. 编辑分页列表、TreeTabe列表中数据后，当前数据行背景闪烁，如果当前数据行由于顺序变化跳出可视区域，则滚动条自动滚动，将当前数据行移动至可视区域；
4. 精确至Action的轻量级功能权限控制，后台用户与角色、角色与资源（菜单、按钮）都是多对多关系，可以灵活配置用户可访问的资源。
# 后端框架
1. 基于Beego 1.9.1，使用官方的orm、cache、session、logs等模块，感谢原作者提供了如此简单易用的框架（<a href="https://beego.me/">更多信息</a>）;
2. 代码风格源自笔者本人其他语言的风格，但参考了多个开源系统的代码风格如 PPGo_ApiAdmin（<a href="https://github.com/george518/PPGo_ApiAdmin">更多信息</a>）、ERP系统（<a href="https://github.com/hexiaoyun128/ERP">更多信息</a>）等。
# 前端框架
1. 基于AdminLTE2（<a href="https://adminlte.io/themes/AdminLTE/index2.html">更多信息</a>）；
2. 弹出层插件使用了Layer（<a href="http://layer.layui.com/">更多信息</a>）；
3. 分页列表使用Bootstrap-table（<a href="http://bootstrap-table.wenzhixin.net.cn/zh-cn/getting-started/">更多信息</a>），并集成cookie、x-editable等辅助插件（<a href="http://bootstrap-table.wenzhixin.net.cn/zh-cn/extensions/">更多信息</a>）实现状态保持、快速编辑等功能。对Bootstrap-table进行了扩展使分页导航可以通过下拉迅速定位。修复Bootstrap-table-cookie的若干Bug；
4. TreeTable列表使用jQuery-treetable插件，按照Boostrap风格进行了样式调整（<a href="http://ludo.cubicphuse.nl/jquery-treetable/">更多信息</a>）；
5. 下拉框使用Bootstrap-select（<a href="http://silviomoreto.github.io/bootstrap-select/">更多信息</a>）；
6. 高亮显示使用的是笔者自已开发的插件；
7. 分页列表里搜索条件、搜索面板状态自动保存使用的是笔者自已开发的插件。

# 增加部分
1. 增加了font-awesome图标选择的功能；
2. 增加了font-awesome图标显示页面；
3. 增加了d3.js展示客户树信息和可全屏展示（<a href="http://www.robschmuecker.com/d3-js-drag-and-drop-zoomable-tree/">更多信息</a>）；
4. 增加了用户登陆信息日志显示；
5. 增加了eCharts和百度地图的展示和全屏展示；
6. 增加了BoxWidget.prototype.maximize功能；
7. 增加了websocket测试实例；
8. 增加了nprogress（<a href="https://github.com/rstacruz/nprogress/">更多信息</a>）；
9. 增加了bootstrapTable控件表头合并列显示方式（<a href="https://github.com/wenzhixin/bootstrap-table/tree/master/src/extensions/sticky-header/">更多信息</a>）；
10. 修改dashborad中增加connectedSortable可移动box；
11. 增加了bootstrap-table列宽手工调整功能（<a href="https://github.com/wenzhixin/bootstrap-table/tree/master/src/extensions/resizable/">更多信息</a>）；
12. 对百度地图做了动态线功能
13. 对eCharts拆线图增加了扩大缩小比例尺
14. 增加了对https的支持

# 安装方法
本系统基于beego开发，默认使用mysql数据库，缓存redis
1. 安装golang环境（ 略）
2. 运行MySQL和Redis
3. 安装本系统
<br>`go get github.com/ygqbasic/poseidon`
4. 将doc目录下的gdkxdl.sql和kxtimingdata.sql导入mysql(业务需求，分库治理)
5. 修改配置文件 `conf/app.conf`
6. 运行在gdkxdl目录使用beego官方提供的命令运行
<br>`bee run`
7. 访问方式
<br>http: 在浏览器里打开 `http://localhost:8080` 进行访问
<br>https: 在浏览器里打开 `https://localhost:10443` 进行访问

# 系统截图

![登陆页面](https://github.com/yunnet/gardens/blob/master/doc/image/login.png)
![仪表板](https://github.com/yunnet/gardens/blob/master/doc/image/dashboard.png)
