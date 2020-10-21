FROM centos:7

RUN yum -y install redhat-lsb-core 

RUN curl -O -L https://golang.org/dl/go1.15.3.linux-amd64.tar.gz

RUN tar -C /usr/local -xzf go1.15.3.linux-amd64.tar.gz

RUN echo 'export PATH=$PATH:/usr/local/go/bin' > ~/.bashrc

# ENTRYPOINT ["/bin/bash"]