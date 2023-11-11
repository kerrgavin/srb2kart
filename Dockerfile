FROM ubuntu as srb2kart-base
WORKDIR /app
COPY ./setup ./setup
RUN chmod 777 setup/*
RUN ./setup/init.sh
USER docker
CMD /bin/bash
RUN ./setup/dependencies.sh
RUN ./setup/build_kart.sh
RUN ./setup/assets.sh
RUN ./setup/clean_up.sh

FROM golang:1.21 as go_build
WORKDIR /app
COPY ./go/* .
RUN go build addons.go

FROM srb2kart-base as srb2kart-custom
WORKDIR /app
COPY run.sh run.sh
COPY kartserv.cfg .srb2kart/kartserv.cfg
COPY --from=go_build /app/addons addonHandler
RUN sudo chmod 777 run.sh
RUN sudo chmod 777 addonHandler
CMD ["./run.sh"]
