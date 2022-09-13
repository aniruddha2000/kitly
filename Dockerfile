FROM ubuntu

COPY ./bin/kitly ./kitly

ENTRYPOINT ["./kitly"]
