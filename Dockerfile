FROM ubuntu

COPY ./bin/kitly ./kitly

EXPOSE 8989

ENTRYPOINT ["./kitly"]
