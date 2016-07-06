FROM golang:1.6-alpine
EXPOSE 8888

WORKDIR /app
COPY assets /app/assets

COPY .shipped/out /app/aie-burnit
RUN chmod a+x /app/aie-burnit

CMD /app/aie-burnit
