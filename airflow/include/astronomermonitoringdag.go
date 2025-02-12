package include

import "strings"

var MonitoringDag = strings.TrimSpace(`
"""A Monitoring DAG used by Astronomer to alert users via Control Plane when tasks aren't executing"""

import os
from datetime import timedelta

from airflow import DAG
from airflow.operators.bash import BashOperator
from airflow.utils.dates import days_ago

with DAG(
    dag_id="astronomer_monitoring_dag",
    schedule_interval=os.environ.get("AIRFLOW_MONITORING_DAG_SCHEDULE_INTERVAL", "*/5 * * * *"),
    start_date=days_ago(2),
    catchup=False,
    is_paused_upon_creation=os.environ.get("PAUSE_ASTRONOMER_MONITORING_DAG", "False").lower()=="true",
    dagrun_timeout=timedelta(minutes=30),
    description=__doc__,
    doc_md=__doc__,
    tags=["monitoring"],
) as dag:
    hello = BashOperator(
        task_id="hello",
        bash_command="echo Hello from Astronomer!",
        depends_on_past=False,
        priority_weight=2 ** 31 - 1,
        do_xcom_push=False,
    )
`)
