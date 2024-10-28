# refer

https://www.topgoer.com/%E9%A1%B9%E7%9B%AE/%E5%AE%9A%E6%97%B6%E4%BB%BB%E5%8A%A1/gocron.html
https://github.com/ouqiang/gocron

https://github.com/ouqiang/gocron/wiki

[//]: # (TODO)

[gocron](https://github.com/Donaldhan/gocron)   

1. 下载gocron安装文件，启动gocron web和node(任务执行器) 服务；

调度器启动

```
./gocron web

```

任务节点启动, 默认监听0.0.0.0:5921

```
./gocron-node

```


浏览器访问 http://localhost:5920

2. 创建数据库gocron；
3. 访问http://127.0.0.1:5920, 进行数据库连接初始化，及管理信息配置；xorm没有对应的数据表，会自动创建；
4. 创建任务节点， 使用shell任务需要，可以选择任务节点；如果使用http方式，不需要选择任务节点，填写对应的url即可；
5. 创建任务

> 创建http任务
创建http任务, 需要填写对应的url， 具体如下：

![http task](/doc/image/gocron/gocron_create_http_task.jpg)

这种方式，针对分布式的web服务，需要使用锁来控制任务的并发执行问题；

>创建shell任务
创建shell任务需要基于执行器任务，必须在给定的机器上，启动gron-node服务；创建shell任务时，选择对应的执行器；

![shell task](/doc/image/gocron/gocron-create-shell-task.png)

注意crontab的调度表达式与http的有点不同， crontab从分钟开始（
0/5 * * * *），http的支持cron规则表达式（0/5 * * * * ?），可以从秒开始；


shell 任务，在node执行器上的，执行结果如下：

```
 
INFO[3059] execute cmd start: [id: 248 cmd: echo `hello gron~'] 
INFO[3059] execute cmd end: [id: 248 cmd: echo `hello gron~' err: exit status 2] 
INFO[3064] execute cmd start: [id: 251 cmd: echo `hello gron~'] 
INFO[3064] execute cmd end: [id: 251 cmd: echo `hello gron~' err: exit status 2] 
INFO[3069] execute cmd start: [id: 252 cmd: echo `hello gron~'] 
INFO[3069] execute cmd end: [id: 252 cmd: echo `hello gron~' err: exit status 2] 
INFO[3074] execute cmd start: [id: 255 cmd: echo `hello gron~'] 
INFO[3074] execute cmd end: [id: 255 cmd: echo `hello gron~' err: exit status 2] 

```



6. 任务管理
![task manager](/doc/image/gocron/gocron-task-manager.png)

在任务管理，我们可以手动执行任务

8. 任务节点管理
![task node](/doc/image/gocron/gocron-task-node.png)

9. 系统管理
针对任务失败的情况，我们可以通过开启任务通知，可以通过邮件进行失败告知；