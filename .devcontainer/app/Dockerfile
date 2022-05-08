FROM golang:1.15-buster

COPY ./.alias /root

# タイムゾーンの設定、パッケージのインストール
RUN apt-get update && apt-get install -y tzdata && apt-get install -y vim\
    && rm -rf /var/lib/apt/lists/* \
    && go get -u github.com/gin-gonic/gin \
    && go get -u gorm.io/gorm \
    && go get -u gorm.io/driver/mysql \
    && go get -u github.com/derekparker/delve/cmd/dlv \
    && cat /root/.alias >> /root/.bashrc

# タイムゾーンを設定する環境変数
ENV TZ=Asia/Tokyo
