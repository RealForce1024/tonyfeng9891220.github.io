# rabbitmq
越来越轻量级的mq，为了asynchronized

http是同步还是异步? 当然是同步

Decoupling system 解耦系统

jms api 目前这套有点过时，但系统理论可以学习下。 ActiveMQ实现jms 

RabbitMQ
message的格式 往往让我们更关注
- AMQP协议
- 提供什么api，由client端提供，较为灵活
- rabbitmq也支持jms
message delivery guarantees
- at most once
- at least once (kafka)
- exactly once (RabbitMQ)
和系统的性能，需求的设计 需要权衡

消防栓的例子 最好是指哪打哪 exactly once
而一开始并不一定能控制好水量，这是at least once 

kafka的优势在于消息可以存储，重新消费

rabbitmq的特点

简单的队列

- [点对点](https://www.rabbitmq.com/tutorials/tutorial-one-go.html) 
- [work queue](https://www.rabbitmq.com/tutorials/tutorial-two-go.html) 
- [订阅发布publish subscribe](https://www.rabbitmq.com/tutorials/tutorial-three-go.html)
- [routing](https://www.rabbitmq.com/tutorials/tutorial-four-go.html)
- [topic](https://www.rabbitmq.com/tutorials/tutorial-five-go.html)
- [rpc](https://www.rabbitmq.com/tutorials/tutorial-six-go.html)


