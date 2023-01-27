FROM centos:centos7

RUN yum -y update && yum install -y git vim make wget \
  && rm -rf /var/cache/yum/* \
  && yum clean all \
;
# goインストール
RUN wget https://go.dev/dl/go1.19.2.linux-amd64.tar.gz \
  && tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz \
  && rm -rf /go1.19.2.linux-amd64.tar.gz \
;
# go path add
ENV PATH="$PATH:/usr/local/go/bin"
# go swagインストール
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.6 \
;
# swag path add
ENV PATH="$PATH:/root/go/bin"

RUN go install -v golang.org/x/tools/gopls@latest

ENV PS1 '$ '