# Building a Golang Microservice for Machine Learning Inference with TensorFlow

For details see: https://blog.skopow.ski/building-a-golang-microservice-for-machine-learning-inference-with-tensorflow

# Sources

- Sample images: https://www.kaggle.com/datasets/kkhandekar/object-detection-sample-images
- ImageNet labels: https://storage.googleapis.com/download.tensorflow.org/data/ImageNetLabels.txt

# Usage

## Run the Application

`go run main.go`

Expected output should be similar to this:
```shell
go run main.go
Loading TF model...
2025-05-18 13:15:43.349562: I tensorflow/cc/saved_model/reader.cc:83] Reading SavedModel from: model/saved_mobilenet_v2
2025-05-18 13:15:43.355372: I tensorflow/cc/saved_model/reader.cc:52] Reading meta graph with tags { serve }
2025-05-18 13:15:43.355394: I tensorflow/cc/saved_model/reader.cc:147] Reading SavedModel debug info (if present) from: model/saved_mobilenet_v2
WARNING: All log messages before absl::InitializeLog() is called are written to STDERR
I0000 00:00:1747566943.416340 11537485 mlir_graph_optimization_pass.cc:425] MLIR V1 optimization pass is not enabled
2025-05-18 13:15:43.422688: I tensorflow/cc/saved_model/loader.cc:236] Restoring SavedModel bundle.
2025-05-18 13:15:43.631278: I tensorflow/cc/saved_model/loader.cc:220] Running initialization op on SavedModel bundle at path: model/saved_mobilenet_v2
2025-05-18 13:15:43.687152: I tensorflow/cc/saved_model/loader.cc:471] SavedModel load for tags { serve }; Status: success: OK. Took 337592 microseconds.
Loading labels...
Setting up handlers...
listening on :8080
```

## Execute the Inference


Run `CURL` to call the REST API for a prediction:
```shell
curl -X POST -F image=@static/example.jpg http://localhost:8080/predict
```

You should see the response like this:
```shell
curl -X POST -F image=@static/example.jpg http://localhost:8080/predict
{"class_id": 469, "label": "cab", "confidence": 12.6021}
```
