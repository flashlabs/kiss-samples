# How to Run YOLOv5 Inference From Golang with Python API

For details see: https://blog.skopow.ski/how-to-run-yolov5-inference-from-golang-with-python-api

# Sources

- Sample images: https://www.kaggle.com/datasets/kkhandekar/object-detection-sample-images
- Virtual Environment: https://docs.python.org/3/library/venv.html

# Usage

## Run the Python Inference Server (YOLOv5)

```shell
cd yolo-api
pip install -r requirements.txt
uvicorn detect:app --host 0.0.0.0 --port 8000
```

## Run the Inference

```shell
cd go-backend && go run main.go
```

Expected output should be similar to this:

```shell
go run main.go
[{"xmin":451.77557373046875,"ymin":256.8055114746094,"xmax":572.8908081054688,"ymax":355.9529724121094,"confidence":0.8660547733306885,"class":41,"name":"cup"},{"xmin":216.73318481445312,"ymin":242.79660034179688,"xmax":417.9637756347656,"ymax":352.3187561035156,"confidence":0.3558332026004791,"class":67,"name":"cell phone"},{"xmin":0.4250640869140625,"ymin":0.6914291381835938,"xmax":276.78839111328125,"ymax":174.0032958984375,"confidence":0.27563828229904175,"class":73,"name":"book"},{"xmin":211.1724090576172,"ymin":242.36141967773438,"xmax":421.87457275390625,"ymax":351.2012634277344,"confidence":0.26584678888320923,"class":63,"name":"laptop"}]
```
