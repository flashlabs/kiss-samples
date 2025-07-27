# How to Run YOLOv8 Inference Directly in Golang (with ONNX)

For details see: https://blog.skopow.ski/how-to-run-yolov8-inference-directly-in-golang-with-onnx

# Sources

- ONNX Runtime: https://github.com/yalue/onnxruntime_go
- Image detection: https://github.com/yalue/onnxruntime_go_examples/tree/master/image_object_detect
- Sample images: https://www.kaggle.com/datasets/kkhandekar/object-detection-sample-images
- Virtual Environment: https://docs.python.org/3/library/venv.html

# Usage

## Run the Inference

```shell
go run main.go
```

Expected output should be similar to this:

```shell
YoloV8 with ONNX by KISS-SAMPLES (blog.skopow.ski):
Box 0: Object laptop (confidence 0.524439): (213.599579, 243.196198), (419.911469, 350.581512)
Box 1: Object cup (confidence 0.563491): (433.477356, 257.403839), (571.929077, 355.463074)
Box 2: Object parking meter (confidence 0.578624): (406.172058, 50.842918), (565.424744, 231.428116)
Creating an ouput image with bounding boxes: ./output.jpg
```

You can find the `output.jpg` image file created with the bounding boxes around detected objects.

