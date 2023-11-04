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

FROM srb2kart-base as srb2kart-custom
WORKDIR /app
COPY run.sh run.sh
COPY addons.sh addons.sh
RUN sudo chmod 777 run.sh
RUN sudo chmod 777 addons.sh
RUN ./addons.sh
CMD ["./run.sh"]
