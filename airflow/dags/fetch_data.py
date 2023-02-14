from airflow import DAG
from airflow.operators.bash_operator import BashOperator
from datetime import datetime, timedelta

default_args = {
    'owner': 'devin',
    'depends_on_past': False,
    'start_date': datetime(2023, 2, 14),
    'email_on_failure': False,
    'email_on_retry': False,
    'retries': 1,
    'retry_delay': timedelta(minutes=5),
}

dag = DAG(
    'fetch_data',
    default_args=default_args,
    description='fetch data from https://smartcity.taipei/projmap/0?lang=zh-Hant',
    schedule_interval=timedelta(hours=24),
)


task = BashOperator(
    task_id='2',
    bash_command='cd /usr/local/airflow/dags && ./cmd fetch',
    dag=dag,
)