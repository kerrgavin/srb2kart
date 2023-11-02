FROM ubuntu as srb2k-base
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

FROM srb2k-base as srb2k-custom
WORKDIR /app
COPY run.sh run.sh
RUN sudo chmod 777 run.sh
