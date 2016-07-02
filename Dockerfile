FROM golang:1.6
EXPOSE 8888
WORKDIR /go/src/app
COPY . /go/src/app

RUN apt-get update && apt-get install -y hashalot
RUN curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.31.0/install.sh | bash
RUN . /root/.bashrc && nvm install v5.0 && nvm use v5.0 && npm install -g grunt-cli
RUN chmod a+x .shipped/build .shipped/run .shipped/test

RUN . /root/.bashrc && .shipped/build
RUN rm -rf node_modules client marathon handlers resources vendor names *.*
CMD .shipped/run
