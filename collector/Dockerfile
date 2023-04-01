FROM python:3.7-slim-buster

WORKDIR /app

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY article article

COPY interest interest

COPY model model

COPY main.py ./

CMD ["python", "-u", "main.py"]

