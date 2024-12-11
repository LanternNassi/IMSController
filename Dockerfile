FROM golang:latest As builder

WORKDIR /app

COPY . .

ARG DBHOST
ARG DBPORT
ARG DBUSER
ARG DBPASSWORD
ARG DBNAME
ARG APPHOST

ARG test_DBHOST
ARG test_DBPORT
ARG test_DBUSER
ARG test_DBPASSWORD
ARG test_DBNAME

RUN echo "DBHOST=${DBHOST}" >> .env
RUN echo "DBPORT=${DBPORT}" >> .env
RUN echo "DBUSER=${DBUSER}" >> .env
RUN echo "DBPASSWORD=${DBPASSWORD}" >> .env
RUN echo "DBNAME=${DBNAME}" >> .env
RUN echo "APPHOST=0.0.0.0:10000" >> .env

RUN echo "test_DBHOST=db" >> .env
RUN echo "test_DBPORT=5432" >> .env
RUN echo "test_DBUSER=postgres" >> .env
RUN echo "test_DBPASSWORD=postgres" >> .env
RUN echo "test_DBNAME=testdb" >> .env



RUN go mod download



FROM builder As final

RUN go build -o IMSController .


EXPOSE 10000

CMD ["./IMSController"]




