# Common

## Data Collection Rule
### 1. kafka naming
    monitor.standard.metrics

### 2. data format
    {timestamp}|{meta}|{data}

|field|format|examples|
|--|--|--|
|timestamp|yyyy-MM-dd HH:mm:ss|2017-01-06 13:06:29|
|meta|{"app":"", "project":"", "env":"", "ip":""}|{"app":"monitor","project":"standard-metric-thread_new","env":"production","ip":"10.10.150.102"}|
|data|{"k1":v1, "k2":v2}|{"app.alive":1}|

#### 2.1 data levels   
    app.project.env.ip.k1 = v1

#### 2.2 meta data options
|label|usage|options|necessity|
|:--:|:--:|:--:|:--:|
|app|source group|app<br/>nginx<br/>redis<br/>mysql<br>...|mandatory|
|project|source project|message-group-service<br>...|mandatory|
|env|source environment|production<br>rc|mandatory|
|ip|source machine|xxx.xxx.xxx.xxx|optional|

#### 2.3 others
    Level will be increased if there's "." inside "k1"
    Use "_" instead "." if there's no more levels

    v1 only supports numerics