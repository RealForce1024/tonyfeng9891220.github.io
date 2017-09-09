# es

```sh
wget https://download.elastic.co/elasticsearch/elasticsearch/elasticsearch-1.7.2.tar.gz
```


[es 1.7 aws discovery](https://www.elastic.co/guide/en/elasticsearch/reference/1.7/modules-discovery-ec2.html)

[es aws plugin](https://github.com/elastic/elasticsearch-cloud-aws)


```sh
bin/plugin install elasticsearch/elasticsearch-cloud-aws/VERSION
```


[es on aws](http://pavelpolyakov.com/2014/08/14/elasticsearch-cluster-on-aws-part-2-configuring-the-elasticsearch/)

[aws 配置es](http://www.awshao.com/%E5%9C%A8aws%E4%B8%8A%E9%85%8D%E7%BD%AEelasticsearch/)


![](media/15049650121344.jpg)


![](media/15049710360481.jpg)

重启一遍服务即可
![](media/15049713415726.jpg)



```yml
cluster.name: awstutorialseries
cloud.aws.access_key: AK********7I3A
cloud.aws.secret_key: B+W*******hYC9Vrbt1RS3eg5D
cloud.aws.region: cn-north-1
discovery.type: ec2
discovery.ec2.tag.Name: "AWS Tutorial Series - Elasticsearch"
http.cors.enabled: true
http.cors.allow-origin: "*"
```
[aws elk](https://www.youtube.com/watch?v=ge8uHdmtb1M&list=PL5zjQdAWZiUyxxHI72D_O5i77jlJrxKZr)



elk大功告成

![](media/15049783061512.jpg)

![](media/15049783131565.jpg)

![](media/15049784184546.jpg)

![](media/15049784439800.jpg)


![](media/15049782857062.jpg)


![](media/15049778559631.jpg)

![](media/15049782655907.jpg)

