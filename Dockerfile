FROM dakr/golangnode
EXPOSE 8888

WORKDIR /app
COPY . /app
RUN npm install
RUN grunt build
RUN bash -c ". ~/.bashrc && .shipped/build"

RUN cp .shipped/out /app/aie-burnit
RUN chmod a+x /app/aie-burnit

CMD /app/aie-burnit
