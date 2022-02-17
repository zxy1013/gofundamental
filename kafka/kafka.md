https://www.bilibili.com/video/BV1Xy4y1G7zA?p=25

**消息队列**：一个队列专门用来存放消息。

为什么需要消息队列，以电商为例。

使用同步的通信方式解决多个服务之间的通信，但是会存在性能问题

![q](F:\markdown笔记\kafka\q.png)

所以需要异步处理，可以让上游快速成功，极大提高了系统的吞吐量。通过下游对个服务的分布式事务的保障，也能保障业务执行的最终一致性

下订单时发送具体的业务信息给消息队列，然后直接变为订单创建成功，消息队列接收消息后将其分配给队列，订阅队列的微服务进行消费。服务完成需要使用分布式事物保证![o](F:\markdown笔记\kafka\o.png)

消息队列解决的具体问题是--通信问题 同步变为异步



**什么是MQ** 

Message Queue（MQ），消息队列中间件。很多⼈都说：MQ 通过将消息的发送和接收分离来实现应⽤程序的异步和解耦，这个给⼈的直觉是——MQ 是异步的，⽤来解耦的，但是这个只是 MQ 的效果⽽不是⽬的。MQ 真正的⽬的是为了通讯，屏蔽底层复杂的通讯协议，定义了⼀套应⽤层的、更加简单的通讯协议。⼀个分布式系统中两个模块之间通讯要么是HTTP，要么是⾃⼰开发的（rpc）TCP，但是这两种协议其实都是原始的协议。HTTP 协议很难实现两端通讯——模块 A 可以调⽤ B，B 也可以主动调⽤ A，如果要做到这个两端都要背上WebServer，⽽且还不⽀持⻓连接（HTTP 2.0 的库根本找不到）。TCP 就更加原始了，粘包、⼼跳、私有的协议，想⼀想头⽪就发麻。MQ 所要做的就是在这些协议之上构建⼀个简单的“协议”——⽣产者/消费者模型。MQ 带给我的“协议”不是具体的通讯协议，⽽是更⾼层次通讯模型。它定义了两个对象——发送数据的叫⽣产者；接收数据的叫消费者， 提供⼀个SDK 让我们可以定义⾃⼰的⽣产者和消费者实现消息通讯⽽⽆视底层通讯协议。

> ⽬前消息队列的中间件选型有很多种： 
>
> rabbitMQ：内部的可玩性（功能性）是⾮常强的 
>
> rocketMQ： 阿⾥内部⼀个⼤神，根据kafka的内部执⾏原理，⼿写的⼀个消息队列中间件。性能是与Kafka相⽐肩，除此之外，在功能上封装了更多。
>
> kafka：全球消息处理性能最快的⼀款MQ 



**消息队列的流派** 

**有Broker的MQ** 

这个流派通常有⼀台服务器作为 Broker，所有的消息都通过它中转。⽣产者把消息发送给它就结束⾃⼰的任务了，Broker 则把消息主动推送给消费者（或者消费者主动轮询） 

> **重Topic** 
>
> 整个broker，依据topic来进⾏消息的中转。在重topic的消息队列⾥必然需要topic的存在
>
> kafka、JMS（ActiveMQ）rocketMQ就属于这个流派，⽣产者会发送 key 和数据到 Broker，由 Broker⽐较 key 之后决定给哪个消费者。这种模式是我们最常⻅的模式，是我们对 MQ 最多的印象。在这种模式下⼀个 topic 往往是⼀个⽐较⼤的概念，甚⾄⼀个系统中就可能只有⼀个topic，topic 某种意义上就是 queue，⽣产者发送 key 相当于说：“hi，把数据放到 key 的队列中” 
>
> ![qq](F:\markdown笔记\kafka\qq.png)
>
> 如上图所示，Broker 定义了三个队列，key1，key2，key3，⽣产者发送数据的时候会发送key1 和 data，Broker 在推送数据的时候则推送 data（也可能把 key 带上）给key1。虽然架构⼀样但是 kafka 的性能要⽐ jms 的性能不知道⾼到多少倍，所以基本这种类型的 MQ 只有 kafka ⼀种备选⽅案。如果你需要⼀条暴⼒的数据流（在乎性能⽽⾮灵活性）那么 kafka 是最好的选择 
>
> **轻Topic** 
>
> 这种的代表是 RabbitMQ（或者说是 AMQP）内部可以用topic也可以不用topic。⽣产者发送 key 和数据，消费者定义订阅的队列，Broker 收到数据之后会通过⼀定的逻辑计算出 key 对应的队列，然后把数据交给队列 
>
> ![1](F:\markdown笔记\kafka\1.png)
>
> 这种模式下解耦了 key 和 queue，在这种架构中 queue 是⾮常轻量级的（在 RabbitMQ中它的上限取决于你的内存），消费者关⼼的只是⾃⼰的 queue；⽣产者不必关⼼数据最终给谁，只要指定 key 就⾏了，中间的那层映射在 AMQP 中叫 exchange（交换机）。 
>
> > **AMQP 中有四种 exchange** 
> >
> > Direct exchange：key 就等于 queue 
> >
> > Fanout exchange：⽆视 key，给所有的 queue 都来⼀份 
> >
> > Topic exchange：key 可以⽤“宽字符”模糊匹配 queue 
> >
> > Headers exchange：⽆视 key，通过查看消息的头部元数据来决定发给哪个queue（AMQP 头部元数据⾮常丰富⽽且可以⾃定义） 
>
> 这种结构的架构给通讯带来了很⼤的灵活性，我们能想到的通讯⽅式都可以⽤这四种exchange 表达出来。如果你需要⼀个企业数据总线（在乎灵活性）那么RabbitMQ 绝对的值得⼀⽤ 
>



**⽆Broker的MQ** 

> ⽆ Broker 的 MQ 的代表是 ZeroMQ。该作者⾮常睿智，他⾮常敏锐的意识到——MQ 是更⾼级的 Socket，它是解决通讯问题的。所以 ZeroMQ 被设计成了⼀个“库”⽽不是⼀个中间件，这种实现也可以达到——没有 Broker 的⽬的 
>
> ![2](F:\markdown笔记\kafka\2.png)
>
> 节点之间通讯的消息都是发送到彼此的队列中，每个节点都既是⽣产者⼜是消费者。ZeroMQ做的事情就是封装出⼀套类似于 Socket 的 API 可以完成发送数据，读取数据 
>
> ZeroMQ 其实就是⼀个跨语⾔的、重量级的 Actor 模型邮箱库。你可以把⾃⼰的程序想象成⼀个 Actor，ZeroMQ 就是提供邮箱功能的库；ZeroMQ 可以实现同⼀台机器的 RPC 通讯也可以实现不同机器的 TCP、UDP 通讯，如果你需要⼀个强⼤的、灵活、野蛮的通讯能⼒，别犹豫 ZeroMQ。
>



**Kafka介绍** 

Kafka是最初由Linkedin公司开发，是⼀个分布式、⽀持分区的（partition）、多副本的 （replica），基于zookeeper协调的分布式消息系统，它的最⼤的特性就是可以实时的处理⼤量数据以满⾜各种需求场景：⽐如基于hadoop的批处理系统、低延迟的实时系统、 Storm/Spark流式处理引擎，web/nginx⽇志、访问⽇志，消息服务等等，⽤scala语⾔编写，Linkedin于2010年贡献给了Apache基⾦会并成为顶级开源项⽬。 

**Kafka的使⽤场景** 

> **⽇志收集**：⼀个公司可以⽤Kafka收集各种服务的log，通过kafka以统⼀接⼝服务的⽅式开放给各种consumer，例如hadoop、Hbase、Solr等。 
>
> **消息系统**：解耦和⽣产者和消费者、缓存消息等。 
>
> **⽤户活动跟踪**：Kafka经常被⽤来记录web⽤户或者app⽤户的各种活动，如浏览⽹⻚、搜索、点击等活动，这些活动信息被各个服务器发布到kafka的topic中，然后订阅者通过订阅这些topic来做实时的监控分析，或者装载到hadoop、数据仓库中做离线分析和挖掘。 
>
> **运营指标**：Kafka也经常⽤来记录运营监控数据。包括收集各种分布式应⽤的数据，⽣产各种操作的集中反馈，⽐如报警和报告。 



**Kafka基本概念** 

kafka是⼀个分布式的，分区的消息(官⽅称之为commit log)服务。它提供⼀个消息系统应该具备的功能，但是确有着独特的设计。可以这样来说，Kafka借鉴了JMS规范的思想，但是确并没有完全遵循JMS规范。

⾸先，让我们来看⼀下基础的消息(Message)相关术语： 

![3](F:\markdown笔记\kafka\3.png)

因此，从⼀个较⾼的层⾯上来看，producer通过⽹络发送消息到Kafka集群，然后consumer来进⾏消费，如下图： 

 ![4](F:\markdown笔记\kafka\4.png)

服务端(brokers)和客户端(producer、consumer)之间通信通过**TCP协议**来完成。 



**kafka基本使⽤** 

**安装前的环境准备** 

> 安装jdk 
>
> 安装zookeeper
>
> 官⽹下载kafka的压缩包:http://kafka.apache.org/downloads 
>
> 解压缩⾄如下路径`/usr/local/kafka/`
>
> 修改配置⽂件：/usr/local/kafka/kafka2.11-2.4/config/server.properties
>
> ```cmd
> # broker.id属性在kafka集群中必须要是唯⼀
> broker.id=0
> # kafka部署的机器ip和提供服务的端⼝号
> listeners=PLAINTEXT://192.168.65.60:9092 
> # kafka的消息存储⽂件
> log.dir=/usr/local/data/kafka-logs
> # kafka连接zookeeper的地址
> zookeeper.connect=192.168.65.60:2181
> ```
>
> ![5](F:\markdown笔记\kafka\5.png)
>
>  ![6](F:\markdown笔记\kafka\6.png)
>
> 进⼊到bin⽬录内，执⾏以下命令来启动kafka服务器（带着配置⽂件）
>
> `./kafka-server-start.sh -daemon ../config/server.properties`
>
> 校验kafka是否启动成功： `ps -aux | grep server.properties` 查看本机是否有进程关联到
>
> 进⼊到zk服务器内查看是否有kafka的节点： `/brokers/ids/0` broker id 为0的kafka



**创建主题topic**

执⾏以下命令创建名为“test”的topic，这个topic只有⼀个partition，并且备份因⼦也设置为1 一个副本 一个分区 

通过kafka命令向zk中创建⼀个主题

`./kafka-topics.sh --create --zookeeper 172.16.253.35:2181 --replication-factor 1 --partitions 1 --topic test`

查看当前kafka内有哪些topic ,查看当前zk中所有的主题

`./kafka-topics.sh --list --zookeeper 172.16.253.35:2181`



**发送消息** 

kafka⾃带了⼀个producer命令客户端，可以从本地⽂件中读取内容，或者我们也可以从命令⾏中直接输⼊内容，并将这些内容以消息的形式发送到kafka集群中。在默认情况下，每⼀个⾏会被当做成⼀个独⽴的消息。使⽤kafka的发送消息的客户端，需要指定发送到的kafka服务器地址和topic 

`./kafka-console-producer.sh --broker-list 172.16.253.38:9092 --topic test`



**消费消息** 

对于consumer，kafka同样也携带了⼀个命令⾏客户端，会将获取到内容在命令中进⾏输出，**默认是消费最新的消息**。使⽤kafka的消费消息的客户端，从指定kafka服务器的指定topic中消费消息 

⽅式⼀：从最后⼀条消息的偏移量+1开始消费，不消费打开服务端之前的消息，只消费新消息

`./kafka-console-consumer.sh --bootstrap-server 172.16.253.38:9092 --topic test`

⽅式⼆：从头开始消费 

`./kafka-console-consumer.sh --bootstrap-server 172.16.253.38:9092 --from-beginning --topic test`

> ⼏个注意点： 
>
> 消息会被存储在本地的⽇志⽂件中`/usr/local/kafka/data/kafka-logs/主题-分区/00000000.log`
>
> 消息是顺序存储 ,通过offset偏移量来描述消息的有序性
>
> 消息是有偏移量的 
>
> 消费时可以指明偏移量进⾏消费
>
> kafka是新加入的消息 未指定from-beginning则从kafka开始消费
>
> ![7](F:\markdown笔记\kafka\7.png)



**Kafka中的关键细节** 

**1消息的顺序存储** 

消息的发送⽅会把消息发送到broker中，broker会存储消息，消息是按照发送的顺序进⾏存储。因此消费者在消费消息时可以指明主题中消息的偏移量。默认情况下，是从最后⼀个消息的下⼀个偏移量开始消费。 

**2单播消息的实现**

在⼀个kafka的topic中，启动两个消费者，⼀个⽣产者，问：⽣产者发送消息，这条消息是否同时会被两个消费者消费？ 

如果多个消费者在同⼀个消费组，那么只有⼀个消费者可以收到订阅的topic中的消息。换⾔之，同⼀个消费组中只能有⼀个消费者收到⼀个topic中的消息。由新连接的消费者消费而不是均匀分配。

`./kafka-console-consumer.sh --bootstrap-server 172.16.253.38:9092 --consumer-property group.id=testGroup --topic test`

**3多播消息的实现** 

在⼀些业务场景中需要让⼀条消息被多个消费者消费，那么就可以使⽤多播模式。kafka实现多播，只需要让不同的消费者处于不同的消费组即可。 

```shell
./kafka-console-consumer.sh --bootstrap-server 172.16.253.38:9092 --consumer-property group.id=testGroup1 --topic test
./kafka-console-consumer.sh --bootstrap-server 172.16.253.38:9092 --consumer-property group.id=testGroup2 --topic test
```

单播和多播的区别

![8](F:\markdown笔记\kafka\8.png)

**4查看消费组及信息** 

```shell
# 查看当前主题下有哪些消费组
./kafka-consumer-groups.sh --bootstrap-server 172.16.253.38:9092 --list
# 查看消费组中的具体信息：⽐如当前消费完的偏移量、最后⼀条消息的偏移量、堆积的消息数量
./kafka-consumer-groups.sh --bootstrap-server 172.16.253.38:9092 --describe --group testGroup
```

> `Currennt-offset`: 当前消费组的已消费偏移量 
>
> `Log-end-offset`: 主题对应分区消息的结束偏移量(HW) 
>
> `Lag`: 当前消费组未消费的消息数 



**主题、分区的概念** 

**主题Topic** 

主题-topic在kafka中是⼀个逻辑的概念，kafka通过topic将消息进⾏分类。不同的topic会被订阅该topic的消费者消费。 

但是有⼀个问题，如果说这个topic中的消息⾮常⾮常多，多到需要⼏T来存，因为消息是会被保存到log⽇志⽂件中的。为了解决这个⽂件过⼤的问题，kafka提出了Partition分区的概念 

**分区Partition** 

**分区的概念** 

⼀个主题中的消息量是⾮常⼤的，因此可以通过分区的设置，来分布式存储这些消息。⽐如⼀个topic创建了3个分区。那么topic中的消息就会分别存放在这三个分区中。

![9](F:\markdown笔记\kafka\9.png)

> 通过partition将⼀个topic中的消息分区来存储。这样的好处有多个： 
>
> 分区存储，可以解决统⼀存储⽂件过⼤的问题 
>
> 提供了读写的吞吐量：读和写可以并行在多个分区中进⾏

**为⼀个主题创建多个分区** 

`./kafka-topics.sh --create --zookeeper 172.16.253.35:2181 --replication-factor 1 --partitions 2 --topic test1`

**可以通过这样的命令查看topic的分区信息** 

`./kafka-topics.sh --list --zookeeper 172.16.253.35:2181`

> 分区的作⽤： 
>
> 可以分布式存储 
>
> 可以并⾏写 
>

> 数据实际上是存在data/kafka-logs/test1-0 和 test1-1中的0000000.log⽂件中 
>
> 000000.index表示某个区间的始末索引，例如：1的位置，10的位置，90的位置。。。。。
>
> 000000.timeindex 表示某个时刻数据的索引
>
> 若要查找某时刻内的数据，可以通过两种稀疏索引很快的定位数据的内容

> **`__consumer_offsets-49`:** 
>
> kafka内部⾃⼰创建了`__consumer_offsets`主题包含了50个分区,因为可能会接收⾼并发的请求。这个⽤来存放消费者消费某个主题的偏移量。
>
> 因为每个消费者都会⾃⼰维护着消费的主题的偏移量，也就是说每个消费者会把消费的主题的偏移量⾃主上报给kafka中的默认主题: consumer_offsets。提交过去的时候，key是consumerGroupId+topic+分区号，value就是当前offset的值，kafka会定期清理topic⾥的消息
>
> 因此kafka为了提升主题的并发性，默认设置了50个分区，提交到哪个分区：通过hash函数：`hash(consumerGroupId) % __consumer_offsets `也解释了为什么同组不能多次消费，不同组可以。
>
> 某一个消费者宕机了另一个消费者就去分区中找偏移量
>
> ⽂件中保存的消息，默认保存7天。七天到后消息会被删除。



**Kafka集群及副本的概念** 

**搭建kafka集群，3个broker** 

准备3个server.properties⽂件 server.properties server1.properties server2.properties 

每个⽂件中的这些内容要调整 

```shell
broker.id=0 
listeners=PLAINTEXT://192.168.65.60:9092 
log.dir=/usr/local/data/kafka-logs 
```

```shell
broker.id=1 
listeners=PLAINTEXT://192.168.65.60:9093 
log.dir=/usr/local/data/kafka-logs-1 
```

```shell
broker.id=2 
listeners=PLAINTEXT://192.168.65.60:9094 
log.dir=/usr/local/data/kafka-logs-2 
```

使⽤如下命令来启动3台服务器 

```shell
./kafka-server-start.sh -daemon ../config/server0.properties 
./kafka-server-start.sh -daemon ../config/server1.properties 
./kafka-server-start.sh -daemon ../config/server2.properties 
```

搭建完后通过查看zk中的`/brokers/ids `看是否启动成功 ,是否有三个znode（0，1，2）



**副本的概念** 

副本是对主题中的分区的备份。在集群中，不同的副本会被部署在不同的broker上,会有⼀个副本作为leader，其他是follower。

下⾯例⼦：创建1个主题，2个分区、3个副本。 

`./kafka-topics.sh --create --zookeeper 172.16.253.35:2181 --replication-factor 3 --partitions 2 --topic my-replicated-topic`

通过kill掉leader后再查看主题情况

```shell
# kill掉leader
ps -aux | grep server.properties
kill 17631
# 查看topic情况
./kafka-topics.sh --describe --zookeeper 172.16.253.35:2181 --topic my-replicated-topic
```

通过查看主题信息，其中的关键数据： `__consumer_offsets-`只需要有一份即可，因为它是存放消费者消费某个主题的偏移量

![10](F:\markdown笔记\kafka\10.png)

> **replicas**：当前副本存在的broker节点 
>
> **leader**：副本⾥的概念，每个partition都有⼀个broker作为leader。 消息发送⽅要把消息发给哪个broker 就看副本的leader是在哪个broker上⾯。副本⾥的leader专⻔⽤来接收消息。接收到消息，其他follower通过poll的⽅式来同步数据。 当leader挂了，经过主从选举，从多个follower中选举产⽣⼀个新的leader
>
> **isr**： 可以同步的broker节点和已同步的broker节点，存放在isr集合中。新leader 从isr集合中选取。如果isr中的节点性能较差，会被踢出isr集合。
>
> **follower**：leader处理所有针对这个partition的读写请求，⽽follower被动复制leader，不提供读写（主要是为了保证多副本数据与消费的⼀致性），如果leader所在的broker挂掉，那么就会进⾏新leader的选举，⾄于怎么选，在之后的controller的概念中介绍。 

集群中有多个broker，创建主题时可以指明主题有多个分区（把消息拆分到不同的分区中存储），可以为分区创建多个副本，不同的副本存放在不同的broker⾥。

 

**kafka集群消息的发送** 

```shell
./kafka-console-producer.sh --broker-list 172.16.253.38:9092,172.16.253.38:9093,172.16.253.38:9094 --topic my-replicated-topic
```

**kafka集群消息的消费** 

```shell
./kafka-console-consumer.sh --bootstrap-server
172.16.253.38:9092,172.16.253.38:9093,172.16.253.38:9094 --from-beginning --topic my- replicated-topic

# 指定消费组来消费消息
./kafka-console-consumer.sh --bootstrap-server
172.16.253.38:9092,172.16.253.38:9093,172.16.253.38:9094 --from-beginning --consumer-property group.id=testGroup1 --topic my-replicated-topic
```

**关于分区消费组消费者的细节** 

![11](F:\markdown笔记\kafka\11.png)

图中Kafka集群有两个broker，每个broker中有多个partition。

⼀个partition只能被⼀个消费组⾥的某⼀个消费者消费，从⽽保证消费顺序。Kafka只在partition的范围内保证消息消费的局部顺序性，不能在同⼀个topic中的多个partition中保证总的消费顺序性。

⼀个消费者可以消费多个partition。建议同⼀个消费组中消费者的数量不要超过partition的数量，否则多的消费者消费不到消息

如果消费者挂了，那么会触发rebalance机制（后⾯介绍），会让其他消费者来消费该分区



**Kafka的go客户端⽣产者**

⽣产者同步发消息，在收到kafka的ack告知发送成功之前⼀直处于阻塞状态，

⽣产者发消息，发送完后不⽤等待broker给回复，直接执⾏下⾯的业务逻辑。可以提供 callback，让broker异步的调⽤callback，告知⽣产者，消息发送的结果

**在同步发消息的场景下：⽣产者发送到broker上后，ack会有3种不同的选择：** 

```
（1）acks=0： 表示producer不需要等待任何broker确认收到消息的回复，就可以继续发送下⼀条消息。性能最⾼，但是最容易丢消息。
（2）acks=1： ⾄少要等待leader已经成功将数据写⼊本地log，但是不需要等待所有follower是否成功写⼊。就可以继续发送下⼀条消息。这种情况下，如果follower没有成功备份数据，⽽此时leader⼜挂掉，则消息会丢失。
（3）acks=-1或all： 需要等待 min.insync.replicas(默认为1，推荐配置⼤于等于2) 这个参数配置的副本个数都成功写⼊⽇志，这种策略会保证只要有⼀个备份存活就不会丢失数据。这是最强的数据保证。⼀般除⾮是⾦融级别，或跟钱打交道的场景才会使⽤这种配置。
```

**Kafka的go客户端消费者**

**⾃动提交offset**

> 消费者poll到消息后默认情况下，会⾃动向broker的_consumer_offsets主题提交当前主题-分区消费的偏移量。 
>
> **⾃动提交会丢消息：**因为如果消费者还没消费完poll下来的消息就⾃动提交了偏移量，那么此时消费者挂了，于是下⼀个消费者会从已提交的offset的下⼀个位置开始消费消息。之前未被消费的消息就丢失掉了。

**⼿动提交offset**

> ⼿动同步提交offset，当前线程会阻塞直到offset提交成功，⼀般使⽤同步提交，因为提交之后⼀般也没有什么逻辑代码了
>
> ⼿动异步提交offset，当前线程提交offset不会阻塞，可以继续处理后⾯的程序逻辑





**Kafka集群Controller、Rebalance和HW** 

`https://blog.csdn.net/shenshouniu/article/details/84716671`

**1.Controller** ---leader 

> Kafka集群中的broker在zk中创建临时序号节点，序号最⼩的节点（最先创建的节点）将作为集群的controller，负责管理整个集群中的所有分区和副本的状态
>
> 当某个分区的leader副本出现故障时，由控制器负责为该分区选举新的leader副本。当检测到某个分区的ISR集合发⽣变化时，由控制器负责通知所有broker更新其元数据信息。 
>
> 当集群中有⼀个副本的leader挂掉，需要在集群中选举出⼀个新的leader，选举的规则是从isr集合中最左边获得。 
>
> 当集群中有broker新增或减少，controller会同步信息给其他broker 
>
> 当使⽤kafka-topics.sh脚本为某个topic增加分区数量时，当集群中有分区新增或减少，controller会同步信息给其他broker

KafkaController中共定义了五种selector选举器 

>   1、ReassignedPartitionLeaderSelector从可用的ISR中选取第一个作为leader，把当前的ISR作为新的ISR，将重分配的副本集合作为接收LeaderAndIsr请求的副本集合。
>   2、PreferredReplicaPartitionLeaderSelector如果从assignedReplicas取出的第一个副本就是分区leader的话，则抛出异常，否则将第一个副本设置为分区leader。
>   3、ControlledShutdownLeaderSelector将ISR中处于关闭状态的副本从集合中去除掉，返回一个新的ISR集合，然后选取第一个副本作为leader，然后令当前AR作为接收LeaderAndIsr请求的副本。
>   4、NoOpLeaderSelector原则上不做任何事情，返回当前的leader和isr。
>   5、OfflinePartitionLeaderSelector从活着的ISR中选择一个broker作为leader，如果ISR中没有活着的副本，则从assignedReplicas中选择一个副本作为leader，leader选举成功后注册到Zookeeper中，并更新所有的缓存。



**2.Rebalance机制** 

**前提**：消费组中的消费者没有指明分区来消费 

**触发的条件**：当消费组中的消费者和分区的关系发⽣变化的时候 

**分区分配的策略**：在rebalance之前，分区怎么分配会有这么三种策略 

> range：根据公式计算得到每个消费消费哪⼏个分区：前⾯的消费者是分区总数/消费者数量+1,之后的消费者是分区总数/消费者数量 。
>
> 轮询：⼤家轮着来  从第一个消费者开始分配
>
> sticky：粘合策略，如果需要rebalance，会在之前已分配的基础上调整，不会改变之前的分配情况。如果这个策略没有开，那么就要进⾏全部的重新分配。建议开启。



**3.HW和LEO**  防止数据丢失

![12](F:\markdown笔记\kafka\12.png)

HW俗称⾼⽔位，HighWatermark的缩写。取⼀个分区对应的ISR中最⼩的LEO(log-end-offset)作为HW 如图中3的位置，consumer最多只能消费到HW所在的位置，若新加入4则不能立即消费。另外每个replica都有HW,leader和follower各⾃负责更新⾃⼰的HW的状态。对于leader新写⼊的消息，consumer不能⽴刻消费，leader会等待该消息被所有ISR中的replicas同步后更新HW，此时消息才能被consumer消费。这样就保证了如果leader所在的broker失效，该消息仍然可以从新选举leader中获取。

如果不按照规定，若消费4，消费完成后宕机，其余broker不再同步成功4，则消息丢失。

消息在写⼊broker时，且每个broker完成这条消息的同步后，hw才会变化。在这之前消费者是消费不到这条消息的。在同步完成之后，HW更新之后，消费者才能消费到这条消息，这样的⽬的是防⽌消息的丢失。



**Kafka线上问题优化** 

**1如何防⽌消息丢失** 

> **发送⽅** 使⽤同步发送。ack是1 或者-1/all 可以防⽌消息丢失，如果要做到99.9999%成功率，ack设成all，把min.insync.replicas配置成分区备份数，分区数>=2 。
>
> **消费⽅** 把⾃动提交改为⼿动提交。 



**2如何防⽌消息的重复消费** 

在防⽌消息丢失的⽅案中，如果⽣产者发送完消息后，因为⽹络抖动，没有收到ack，但实际上broker已经收到了。此时⽣产者会进⾏重试，于是broker就会收到多条相同的消息，⽽造成消费者的重复消费。

如果为了消息的不重复消费，⽽把⽣产端的重试机制关闭、消费端的⼿动提交改成⾃动提交，这样反⽽会出现消息丢失

那么可以直接在防治消息丢失的⼿段上再加上消费消息时的幂等性保证，就能解决消息的重复消费问题。 

**幂等性**如何保证： 

> 所谓的幂等性：多次访问的结果是⼀样的。eg: 对于rest的请求（get（幂等）、post（⾮幂等）、put（幂等）、delete（幂等）多次处理id相同） 
>
> 解决⽅案： 
>
> 在数据库中创建联合主键，防⽌相同的主键 创建出多条记录 
>
> 使⽤分布式锁，以业务id为锁。保证只有⼀条记录能够创建成功

 

**3如何做到顺序消费RocketMQ** 

发送⽅：在发送时将ack不能设置0，关闭重试，使⽤同步发送，等到发送成功再发送下⼀条。确保消息是顺序发送的。 

接收⽅： 同个消费者组，多个消费者方式 。消息只能发送到⼀个分区中，只能有⼀个消费组的消费者来接收消息（单播）。因此，kafka的顺序消费会牺牲掉性能。 kafka的顺序消费使⽤场景不多，因为牺牲掉了性能，但是⽐如rocketmq在这⼀块有专⻔的功能已设计好。



**4解决消息积压问题** 

消息的消费者的消费速度远赶不上⽣产者的⽣产消息的速度，导致kafka中有⼤量的数据没有被消费。随着没有被消费的数据堆积越多，消费者寻址的性能会越来越差，最后导致整个kafka对外提供的服务的性能很差，从⽽造成其他服务也访问速度变慢，造成服务雪崩。

在这个消费者中，使⽤多线程，充分利⽤机器的性能进⾏消费消息。

通过业务的架构设计，提升业务代码层⾯消费的性能。 

> ⽅案⼀：在⼀个消费者中启动多个线程，让多个线程同时消费。——提升⼀个消费者的消费能⼒。 
>
> ⽅案⼆：如果⽅案⼀还不够的话，这个时候可以启动多个消费者，多个消费者部署在不同的服务器上。其实多个消费者部署在同⼀服务器上也可以提⾼消费能⼒——充分利⽤服务器的cpu资源。
>
> ⽅案三：让⼀个消费者去把收到的消息往另外⼀个topic上发，另⼀个topic设置多个分区和多个消费者 ，进⾏具体的业务消费。 即创建⼀个消费者，该消费者在kafka另建⼀个主题，配上多个分区，多个分区再配上多个消费者。该消费者将poll下来的消息，不进⾏消费，直接转发到新建的主题上。此时，新的主题的多个分区的多个消费者就开始⼀起消费了。——不常⽤



**5延迟队列** 

延迟队列的应⽤场景：在订单创建成功后，如果超过30分钟没有付款，则需要取消订单，此时可⽤延时队列来实现

![13](F:\markdown笔记\kafka\13.png)

创建多个topic，每个topic表示延时的间隔

> topic_5s: 延时5s执⾏的队列 
>
> topic_1m: 延时1分钟执⾏的队列 
>
> topic_30m: 延时30分钟执⾏的队列 

消息发送者发送消息到相应的topic，并带上消息的发送时间。消费者订阅相应的topic，消费时轮询消费整个topic中的消息。

第一轮时当前时间为3:00，所以无法消费订单5，所以记录。下次继续从5开始消费。

如果消息的发送时间和消费的当前时间超过预设的值，⽐如30分钟，去数据库中修改订单状态为已取消

如果消息的发送时间和消费的当前时间没有超过预设的值，则不消费当前的offset 及之后的offset的所有消息，等待1分钟后，再次向kafka拉取该offset及之后的消息，继续进⾏判断，以此反复。

