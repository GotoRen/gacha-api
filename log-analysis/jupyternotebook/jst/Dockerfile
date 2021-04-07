FROM python:3.9.0

RUN pip3 install jupyter \
        --upgrade pip \
        numpy \
        pandas \
        matplotlib \
        pytz \
        && mkdir ~/.jupyter \
        && echo 'c.NotebookApp.token = ""' > ~/.jupyter/jupyter_notebook_config.py \
        && echo 'c.NotebookApp.password = ""' >> ~/.jupyter/jupyter_notebook_config.py

WORKDIR /home/jovyan

EXPOSE 8888

CMD ["/usr/local/bin/jupyter", "notebook", "--port", "8888", "--ip=0.0.0.0", "--allow-root"]