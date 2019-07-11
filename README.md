## Grade Service

<!-- ![](https://travis-ci.org/muxih4ck/Go-Web-Application-Template.svg?branch=master) -->

匣子成绩查询服务

### 开发

```
mkdir $GOPATH/src/github.com/asynccnu && cd $GOPATH/src/github.com/asynccnu
git clone https://github.com/asynccnu/grade_service_v2.git
cd grade_service_v2
make && ./main
```

### 测试

```
make test
```

### 环境变量

```
CCNUBOX_GRADE_DATA_SERVICE_URL // data 服务 URL
CCNUBOX_GRADE_RUNMODE
```