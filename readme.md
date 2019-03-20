#### 说明
半小时更新一次,如果ip变化调用阿里云sdk更新解析地址


#### 配置文件
1.创建config.yaml
ex.
```

DomainName : 'wssmyy.xyz' # 主机域名
RR : '@' #主机记录，如果要解析@.exmaple.com，主机记录要填写”@”，而不是空  参考附一
Type : 'A' # A 解析一个ipv4地址
TTL : 600 # 解析生效时间 默认600 10分钟
AccessKeyId  : LT*********h6 # 阿里云获得key
AccessSecret : 4Z59*************jLnSM # 阿里云获得秘钥

```

---
#### 附一
主机记录就是域名前缀，常见用法有：
 - www：解析后的域名为www.aliyun.com。
 - @：直接解析主域名 aliyun.com。
 - *：泛解析，匹配其他所有域名 *.aliyun.com。
 - mail：将域名解析为mail.aliyun.com，通常用于解析邮箱服务器。
 - 二级域名：如：abc.aliyun.com，填写abc。
 - 手机网站：如：m.aliyun.com，填写m。
 - 显性URL：不支持泛解析（泛解析：将所有子域名解析到同一地址）