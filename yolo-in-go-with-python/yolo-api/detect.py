from fastapi import FastAPI, File, UploadFile
from fastapi.responses import JSONResponse
import torch
from PIL import Image
import io

# Initialize the FastAPI application
app = FastAPI()
# Load the pretrained YOLOv5s model from the Ultralytics repository
model = torch.hub.load("ultralytics/yolov5", "yolov5s", pretrained=True)

# Define the endpoint to handle object detection requests
@app.post("/detect")
async def detect(file: UploadFile = File(...)):
    # Read the uploaded image file as bytes
    image_bytes = await file.read()
    # Convert the byte data to a PIL Image
    image = Image.open(io.BytesIO(image_bytes))
    # Run the image through the YOLO model
    results = model(image)
    # Convert the detection results to a JSON response
    return JSONResponse(results.pandas().xyxy[0].to_dict(orient="records"))
