FROM puckel/docker-airflow
COPY ./airflow/dags /usr/local/airflow/dags
COPY ./config/config.yaml /usr/local/airflow/
CMD ["webserver"]