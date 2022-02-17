`https://www.bilibili.com/video/BV1Ph411n7Ep?from=search&seid=14867632069241504863&spm_id_from=333.337.0.0`

**Zookeeper介绍** 

**什么是Zookeeper** 

ZooKeeper 是⼀种分布式协调服务，⽤于管理⼤型主机。在分布式环境中协调和管理服务是⼀个复杂的过程。ZooKeeper 通过其简单的架构和 API 解决了这个问题。ZooKeeper 允许开发⼈员专注于核⼼应⽤程序逻辑，⽽不必担⼼应⽤程序的分布式特性。 

**Zookeeper的应⽤场景** 

> **分布式协调组件** 
>
> 在分布式系统中，需要有zookeeper作为分布式协调组件，协调分布式系统中的状态。 保持服务1和2之间的数据一致性。![1](F:\markdown笔记\zookeeper\1.png)
>
> **分布式锁** 
>
> zk在实现分布式锁上，可以做到强⼀致性 顺序一致性，关于分布式锁相关的知识，在之后的ZAB协议中介绍。 
>
> **⽆状态化的实现** 
>
> 保存信息放在协调组件中![2](F:\markdown笔记\zookeeper\2.jpg)





**搭建Zookeeper服务器** 

> **1.zoo_sample.cfg** **配置⽂件说明**  将其改为zoo.cfg
>
> ```shell
> # zookeeper时间配置中的基本单位 (毫秒) 2s
> tickTime=2000 
> 
> # 允许follower初始化连接到leader最⼤时⻓，它表示tickTime时间倍数，即:initLimit*tickTime 20s
> initLimit=10 
> 
> # 允许follower与leader数据同步最⼤时⻓,它表示tickTime时间倍数 10s
> syncLimit=5 
> 
> #zookeper 数据存储⽬录及⽇志保存⽬录（如果没有指明dataLogDir，则⽇志也保存在这个⽂件中） 持久化机制 事务持久化和快照持久化
> dataDir=/tmp/zookeeper 
> 
> #对客户端提供的端⼝号 
> clientPort=2181 
> 
> #单个客户端与zookeeper最⼤并发连接数 
> maxClientCnxns=60 
> 
> # 保存的数据快照数量，之外的将会被清除 
> autopurge.snapRetainCount=3 
> 
> #⾃动触发清除任务时间间隔，⼩时为单位。默认为0，表示不⾃动清除。 
> autopurge.purgeInterval=1 
> ```
>
> **2.Zookeeper服务器的操作命令** 
>
> 重命名 conf中的⽂件`zoo_sample.cfg`->`zoo.cfg `
>
> 启动zk服务器：`./bin/zkServer.sh start ./conf/zoo.cfg `
>
> 查看zk服务器状态： `./bin/zkServer.sh status ./conf/zoo.cfg `
>
> 停⽌zk服务器： `./bin/zkServer.sh stop ./conf/zoo.cfg `





**Zookeepe内部的数据模型** 

> **1.zk是如何保存数据的** 
>
> zk中的数据是保存在节点上的，节点就是znode，多个znode之间构成⼀颗树的⽬录结构。 Zookeeper 的数据模型是什么样⼦呢？它很像数据结构当中的树，也很像⽂件系统的⽬录。 ![22](F:\markdown笔记\zookeeper\22.jpg)
>
> 树是由节点所组成，Zookeeper 的数据存储也同样是基于节点，这种节点叫做 **Znode** 。但是，不同于树的节点，Znode 的引⽤⽅式是路径引⽤，类似于⽂件路径： 
>
> ```shell
> /动物/猫  # 节点名   可以存数据
> /汽⻋/宝⻢ 
> ```
>
> 这样的层级结构，让每⼀个 Znode 节点拥有唯⼀的路径，就像命名空间⼀样对不同信息作出清晰的隔离。 
>
> 
>
> **2.zk中的znode是什么样的结构** 
>
> zk中的znode，包含了四个部分： 
>
> > data：保存数据 
> >
> > acl：权限，定义了什么样的⽤户能够操作这个节点，且能够进⾏怎样的操作。 
> >
> > ​		c: create 创建权限，允许在该节点下创建⼦节点 
> >
> > ​		w：write 更新权限，允许更新该节点的数据 
> >
> > ​		r：read 读取权限，允许读取该节点的内容以及⼦节点的列表信息 
> >
> > ​		d：delete 删除权限，允许删除该节点的⼦节点 
> >
> > ​		a：admin 管理者权限，允许对该节点进⾏acl权限设置 
> >
> > stat：描述当前znode的元数据  `get -s`查看
> >
> > child：当前节点的⼦节点 
>
> 
>
> **3.zk中节点znode的类型** 
>
> > **持久节点 不带任何选项**: 创建出的节点，在会话结束后依然存在。不可重复创建数据 
> >
> > **持久序号节点 -s**: 创建出的节点，根据先后顺序，会在节点之后带上⼀个数值，越后执⾏数值越⼤，适⽤于分布式锁的应⽤场景- 单调递增 
> >
> > **临时节点 -e**: 临时节点是在会话结束后，⾃动被删除的，通过这个特性，zk可以**实现服务注册与发现**的效果。那么临时节点是如何维持⼼跳呢？ 服务提供者作为客户端注册服务，消费者可在服务端查询可用服务
> >
> > ![1641797674(1)](F:\markdown笔记\zookeeper\1641797674(1).jpg)
> >
> > **临时序号节点 -e -s**：跟持久序号节点相同，适⽤于临时的分布式锁。 
> >
> > **Container节点 -c**（3.5.3版本新增）：Container容器节点，当容器中没有任何⼦节点，该容器节点会被zk定期删除（60s）。 
> >
> > **TTL节点**：可以指定节点的到期时间，到期后被zk定时删除。只能通过系统配置`zookeeper.extendedTypesEnabled=true` 开启 ,目前不稳定
>
> 
>
> **4.zk数据持久化** 
>
> zk的数据是运⾏在内存中，zk提供了两种持久化机制： 
>
> **事务⽇志** 
>
> zk把执⾏的命令以⽇志形式保存在dataLogDir指定的路径中的⽂件中（如果没有指定dataLogDir，则按dataDir指定的路径）。 
>
> **数据快照** 产生数据文件
>
> zk会在⼀定的时间间隔内做⼀次内存数据的快照，把该时刻的内存数据保存在快照⽂件中。 
>
> zk通过两种形式的持久化，在恢复时**先恢复快照⽂件中的数据到内存中，再⽤⽇志⽂件中的数据做增量恢复**，这样的恢复速度更快。 
>



**Zookeeper客户端(zkCli)的使⽤** 

> **1.多节点类型创建** 
>
> > 创建持久节点 `create /test1 data`
> >
> > 创建持久序号节点  `create -s /test1/test2 data`
> >
> > 创建临时节点  `create -e /test3 data`
> >
> > 创建临时序号节点 `create -e -s /test4 data`
> >
> > 创建容器节点 `create -c /test5 data`
>
> **2.存储数据**
>
> `create /test6`
>
> `set /test6 date6`
>
> **3.查询节点** 
>
> 普通查询 `ls` `ls -R test1` 递归查询所有子节点
>
> 查询节点详细信息 `get -s test1`
>
> > cZxid: 创建节点的事务ID  
> >
> > mZxid：修改节点的事务ID 
> >
> > pZxid：添加和删除⼦节点的事务ID 
> >
> > ctime：节点创建的时间 
> >
> > mtime: 节点最近修改的时间 
> >
> > dataVersion: 节点内数据的版本，每更新⼀次数据，版本会+1 
> >
> > aclVersion: 此节点的权限版本 
> >
> > ephemeralOwner: 如果当前节点是临时节点，该值是当前节点所有者的session id。如果节点不是临时节点，则该值为零。 
> >
> > dataLength: 节点内数据的⻓度 
> >
> > numChildren: 该节点的⼦节点个数 
> >
>
> **4.删除节点** 
>
> 普通删除 `deleteall /test1` `delete /test2` 
>
> 乐观锁删除,删除指定版本号的节点`delete -v 1 /test6`
>
> **5.权限设置** 
>
> 注册当前会话的账号和密码`addauth digest xiaowang:123456` 
>
> 创建节点并设置权限 `create /test-node abcd auth:xiaowang:123456:cdwra`
>
> 在另⼀个会话中必须先使⽤账号密码`addauth digest xiaowang:123456` ，才能拥有操作该节点的权限 





**go 客户端**





**zk实现分布式锁** 

> **1.zk中锁的种类：** 
>
> 读锁 共享锁：⼤家都可以读，要想上读锁的前提：之前的锁没有写锁 
>
> 写锁：只有得到写锁的才能写。要想上写锁的前提是，之前没有任何锁。 
>
>  **2.zk如何上读锁** 
>
> > 创建⼀个**临时序号**节点，节点的数据是read，表示是读锁 
> >
> > 获取当前zk中序号⽐⾃⼰⼩的所有节点 
> >
> > 判断最⼩节点是否是读锁： 最小节点是读锁，剩余的全为读，最小节点为写，无其余节点。
> >
> > ​		如果不是读锁的话，则上锁失败，为最⼩节点设置监听。阻塞等待，zk的watch机制。会当最⼩节点发⽣变化时通知当前节点，于是再执⾏第⼆步的流程 。
> >
> > ​		如果是读锁的话，则上锁成功 
>
> ![33](F:\markdown笔记\zookeeper\33.jpg)
>
> **3.zk如何上写锁** 
>
> > 创建⼀个**临时序号**节点，节点的数据是write，表示是写锁 
> >
> > 获取zk中所有的⼦节点，判断⾃⼰是否是最⼩的节点： 写锁前面不能有任何锁
> >
> > ​		如果是，则上写锁成功 
> >
> > ​		如果不是，说明前⾯还有锁，则上锁失败，监听最⼩的节点，如果最⼩节点有变化，则回到第⼆步。 
> >
>
> **4.⽺群效应** 
>
> 如果⽤上述的上锁⽅式，只要有节点发⽣变化，就会触发其他节点的监听事件，这样的话对zk的压⼒⾮常⼤，⽺群效应。可以调整成链式监听。解决这个问题。并且可以顺序执行。 ![1641802338(1)](F:\markdown笔记\zookeeper\1641802338(1).jpg)
>



**zk的watch机制** 

>  **1.Watch机制介绍** 
>
> 我们可以把 **Watch** 理解成是注册在特定 Znode 上的触发器。当这个 Znode 发⽣改变，也就是调⽤了 create ， delete ， setData ⽅法的时候，将会触发 Znode 上注册的对应事件， 请求 Watch 的客户端会接收到异步通知。 
>
> **具体交互过程如下：** 
>
> 客户端调⽤ getData ⽅法， watch 参数是 true 。服务端接到请求，返回节点数据，并且在对应的哈希表⾥插⼊被 Watch 的 Znode 路径，以及 Watcher 列表。 
>
> 当被 Watch 的 Znode 已删除，服务端会查找哈希表，找到该 Znode 对应的所有Watcher，异步通知客户端，并且删除哈希表中对应的 Key-Value。 
>
> ![1641802573(1)](F:\markdown笔记\zookeeper\1641802573(1).jpg)
>
> 客户端使⽤了NIO通信模式监听服务端的调⽤。 
>
> 
>
> **2.zkCli客户端使⽤watch** 
>
> `create /test `
>
> `get -w /test` ⼀次性监听节点 内容的变化，或节点的删除。 并拿取数据。不能监听当前节点的子节点
>
> `ls -w /test` 监听节点⽬录,创建和删除⼦节点会收到通知。⼦节点中新增节点不会收到通知 
>
> `ls -R -w /test` 对于⼦节点下⼦节点的变化 增加删除，但内容的变化不会收到通知 





**Zookeeper集群实战** 

> **1.Zookeeper集群⻆⾊** 
>
> > zookeeper集群中的服务器节点有三种⻆⾊ 
> >
> > Leader：处理集群的所有事务请求读+写，集群中只有⼀个Leader。 
> >
> > Follower：只能处理读请求，参与Leader选举。 
> >
> > Observer：只能处理读请求，提升集群读的性能，但不能参与Leader选举。 
> >
>
> 
>
> **2.集群搭建** 
>
> > 搭建4个节点，其中⼀个节点为Observer 
> >
> > **1）创建4个节点的myid，并设值** 
> >
> > 在/usr/local/zookeeper中创建以下四个⽂件，存入id
> >
> > ```shell
> > /usr/local/zookeeper/zkdata/zk1# echo 1 > myid 
> > /usr/local/zookeeper/zkdata/zk2# echo 2 > myid 
> > /usr/local/zookeeper/zkdata/zk3# echo 3 > myid 
> > /usr/local/zookeeper/zkdata/zk4# echo 4 > myid 
> > ```
> >
> > **2）编写4个zoo.cfg** 
> >
> > ```shell
> > # The number of milliseconds of each tick
> > tickTime=2000
> > # The number of ticks that the initial
> > # synchronization phase can take
> > initLimit=10
> > # The number of ticks that can pass between
> > # sending a request and getting an acknowledgement
> > syncLimit=5
> > # 修改对应的zk1 zk2 zk3 zk4
> > dataDir=/usr/local/zookeeper/zkdata/zk1
> > # 修改对应的端⼝ 2181 2182 2183 2184
> > clientPort=2181 # 2181是客户端访问端口
> > # 200*为集群之间通信端⼝，300*为集群leader选举端⼝，observer表示不参与集群选举
> > server.1=172.16.253.54:2001:3001
> > server.2=172.16.253.54:2002:3002
> > server.3=172.16.253.54:2003:3003
> > server.4=172.16.253.54:2004:3004:observer
> > ```
> >
> > **3）启动4台Zookeeper** 
> >
> > ```shell
> > ./bin/zkServer.sh start ./conf/zoo1.cfg 
> > ./bin/zkServer.sh start ./conf/zoo2.cfg 
> > ./bin/zkServer.sh start ./conf/zoo3.cfg 
> > ./bin/zkServer.sh start ./conf/zoo4.cfg 
> > 
> > # 查看状态 leader
> > ./bin/zkServer.sh status ./conf/zoo1.cfg 
> > ./bin/zkServer.sh status ./conf/zoo2.cfg 
> > ./bin/zkServer.sh status ./conf/zoo3.cfg 
> > ./bin/zkServer.sh status ./conf/zoo4.cfg 
> > ```
>
> **3.客户端连接Zookeeper集群** 
>
> ```shell
> ./bin/zkCli.sh -server 172.16.253.54:2181,172.16.253.54:2182,172.16.253.54:218
> ```





**ZAB协议** 

> **1.什么是ZAB协议** 
>
> zookeeper作为⾮常重要的分布式协调组件，需要进⾏集群部署，集群中会以⼀主多从的形式进⾏部署。zookeeper为了保证数据的⼀致性，使用ZAB（Zookeeper Atomic Broadcast）协议，这个协议解决了Zookeeper的崩溃恢复和主从数据同步的问题。 
>
> ![1641805510(1)](F:\markdown笔记\zookeeper\1641805510(1).jpg)
>
> **2.ZAB协议定义的四种节点状态** 
>
> > Looking ：选举状态。 
> >
> > Following ：Follower 节点（从节点）所处的状态。 
> >
> > Leading ：Leader 节点（主节点）所处状态。 
> >
> > Observing：观察者节点所处的状态 
> >
>
> **3.集群上线时的Leader选举过程**  上面的四个节点为例 一个不参与 则剩余三台，前两个进行选举，过半才可，第三个节点进来发现已经有leader，节点为基数个较好。
>
> Zookeeper集群中的节点在上线时，将会进⼊到Looking状态，也就是选举Leader的状态，这个状态具体会发⽣什么？ zXid表示事务id 增删改加1，更大的选票：先比事务id(因为最新)，再myid
>
> ![1641805630(1)](F:\markdown笔记\zookeeper\1641805630(1).jpg)
>
> **4.崩溃恢复时的Leader选举**  redis借用了此思想
>
> > Leader建⽴完后，Leader周期性地不断向Follower发送⼼跳（ping命令，没有内容的 socket）。当Leader崩溃后，Follower发现socket通道已关闭，于是Follower开始进⼊到Looking状态，重新回到上⼀节中的Leader选举过程，**此时集群不能对外提供服务**。 
> >
> > ![44](F:\markdown笔记\zookeeper\44.png)
>
> **5.主从服务器之间的数据同步** 
>
> 3：广播过程
>
> 6：整个集群的半数，leader通知follower时自己也写入内存
>
> 完成7以后，客户端才能查询，两阶段提交
>
> ![5](F:\markdown笔记\zookeeper\5.png)
>
> **6.Zookeeper中的NIO与BIO的应⽤** 
>
> > **NIO  Non-Blocking  IO**
> >
> > ⽤于被客户端连接的2181端⼝，使⽤的是NIO模式与客户端建⽴连接 
> >
> > 客户端开启Watch时，也使⽤NIO，等待Zookeeper服务器的回调 
> >
> > **BIO  Blocking IO **
> >
> > 集群在选举时，多个节点之间的投票通信端⼝，使⽤BIO进⾏通信。
> >
> > 节点发送socket心跳 



**CAP理论** 

> **1.CAP定理** 
>
> > 2000 年 7 ⽉，加州⼤学伯克利分校的 Eric Brewer 教授在 ACM PODC 会议上提出 CAP 猜想。2年后，麻省理⼯学院的 Seth Gilbert 和 Nancy Lynch 从理论上证明了 CAP。之后，CAP 理论正式成为分布式计算领域的公认定理。 
> >
> > CAP 理论为：<mark>⼀个分布式系统最多只能同时满⾜⼀致性（Consistency）、可⽤性 （Availability）和分区容错性（Partition tolerance）这三项中的两项。</mark>
> >
> > **⼀致性（Consistency）** 
> >
> > ⼀致性指 `all nodes see the same data at the same time`，即更新操作成功并返回客户端完成后，所有节点在同⼀时间的数据完全⼀致。 
> >
> > **可⽤性（Availability）** 
> >
> > 可⽤性指`Reads and writes always succeed`，即服务⼀直可⽤，⽽且是正常响应时间。 
> >
> > **分区容错性（Partition tolerance）** 
> >
> > 分区容错性指`the system continues to operate despite arbitrary message loss or failure of part of the system`，即分布式系统在遇到某节点或⽹络分区故障的时候，仍然能够对外提供满⾜⼀致性或可⽤性的服务。——避免单点故障，就要进⾏冗余部署，冗余部署相当于是服务的分区，这样的分区就具备了容错性。 
> >
>
> **2.CAP权衡** 
>
> > 通过 CAP 理论，我们知道⽆法同时满⾜⼀致性、可⽤性和分区容错性这三个特性，那要舍弃哪个呢？ 
> >
> > 对于多数⼤型互联⽹应⽤的场景，主机众多、部署分散，⽽且现在的集群规模越来越⼤，所以节点故障、⽹络故障是常态，⽽且要保证服务可⽤性达到 N 个 9，即`保证 P 和 A，舍弃C`（退⽽求其次保证最终⼀致性）。虽然某些地⽅会影响客户体验，但没达到造成⽤户流程的严重程度。 
> >
> > 对于涉及到钱财这样不能有⼀丝让步的场景，C 必须保证。⽹络发⽣故障宁可停⽌服务，这是`保证 CP，舍弃 A`。
> >
> > 孰优孰略，没有定论，只能根据场景定夺，适合的才是最好的。 
> >
> > ![66](F:\markdown笔记\zookeeper\66.png)
> >
>
> **3.BASE理论** 
>
> > eBay 的架构师 Dan Pritchett 源于对⼤规模分布式系统的实践总结，在 ACM 上发表⽂章提出BASE 理论，BASE 理论是对 CAP 理论的延伸，核⼼思想是即使⽆法做到强⼀致性（Strong Consistency，CAP 的⼀致性就是强⼀致性），但应⽤可以采⽤适合的⽅式达到最终⼀致性（Eventual Consitency）。 
> >
> > **基本可⽤（Basically Available）** 
> >
> > 基本可⽤是指分布式系统在出现故障的时候，**允许损失部分可⽤性，即保证核⼼可⽤**。 电商⼤促时，为了应对访问量激增，部分⽤户可能会被引导到降级⻚⾯，服务层也可能只提供降级服务（注册 评论 退款 不可用）。这就是损失部分可⽤性的体现。 
> >
> > **软状态（Soft State）** 
> >
> > 软状态是指允许系统存在中间状态，⽽该中间状态不会影响系统整体可⽤性。分布式存储中⼀般⼀份数据⾄少会有三个副本，允许不同节点间副本同步的延时就是软状态的体现。mysql replication 的异步复制也是⼀种体现。 
> >
> > **最终⼀致性（Eventual Consistency）** 
> >
> > 最终⼀致性是指系统中的所有数据副本经过⼀定时间后，最终能够达到⼀致的状态。
> >
>
> **4.Zookeeper追求的⼀致性** 
>
> > Zookeeper在数据同步CP时，`追求的并不是强⼀致性，⽽是最终⼀致性，顺序⼀致性(事务id的单调递增)`。
> >