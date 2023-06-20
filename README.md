
# RobotChat
主要是整合现在比较火的GPT和SD的AIGC模块，该项目现阶段是我们内部需要使用，所以属于快速迭代的阶段。

# ArtBot
基于StableDiffusion的接口对接应用

# ChatBot
现在基于GPT的接口对接应用

## 模块架构设计

![img.png](resource/static/images/structure.png)



## 使用Gin的网络接口访问
可以作为一个独立的微服务，供其他项目调用，比如php，java等，通过http协议访问robot，请求作业。

## 直接通过Module引入Golang项目
可以直接引入到项目中使用，比如我们自己在做的PowerPrompt项目

## 支持队列消息调度

为了应付多用户访问并发需求，支持队列消息请求