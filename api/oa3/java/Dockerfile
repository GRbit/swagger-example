FROM adoptopenjdk/openjdk11:jre-11.0.9.1_1

ARG USER_ID=1000
ARG GROUP_ID=1000

RUN groupadd -g ${GROUP_ID} user &&\
    useradd -l -u ${USER_ID} -g user user &&\
    install -d -m 0755 -o user -g user /home/user

RUN apt-get update
RUN apt-get install wget tar -y
RUN wget https://golang.org/dl/go1.18.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.18.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin
ENV GO111MODULE=on
USER 1000
RUN go install golang.org/x/tools/cmd/goimports@latest
