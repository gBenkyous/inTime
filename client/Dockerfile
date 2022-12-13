FROM centos:centos7

RUN yum -y update && yum install -y git vim \
  && rm -rf /var/cache/yum/* \
  && yum clean all \
;

ENV PS1 '$ '