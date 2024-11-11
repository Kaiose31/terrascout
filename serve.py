import torch
import torch.nn as nn
from flask import Flask, request, jsonify
import json
import numpy as np
import torchvision.models as models
from torchvision.transforms import transforms
from PIL import Image
import tensorflow as tf
import keras
import pickle
from sklearn.preprocessing import LabelEncoder

# Pipeline 1
model = models.mobilenet_v2(pretrained=False)
model.classifier[1] = nn.Linear(model.classifier[1].in_features, 4)
model.load_state_dict(torch.load("data/models/mobilenetft.pth"))
model.eval()

# Pipeline 2
plant = tf.keras.models.load_model("data/models/Plant_model.h5")
insect = tf.keras.models.load_model("data/models/Insect_model.h5")
bird = tf.keras.models.load_model("data/models/Bird_model.h5")
mammal = tf.keras.models.load_model("data/models/Mammal_model.h5")

app = Flask(__name__)

categories = {0: "Plantae", 1: "Insecta", 2: "Aves", 3: "Mammalia"}


@app.route("/inference", methods=["POST"])
def inference():
    data = request.files["image"].stream
    image = Image.open(data).convert("RGB")
    image2 = image
    preprocess = transforms.Compose(
        [
            transforms.Resize(256),
            transforms.CenterCrop(224),
            transforms.ToTensor(),
            transforms.Normalize(mean=[0.485, 0.456, 0.406], std=[0.229, 0.224, 0.225]),
        ]
    )
    image = preprocess(image)
    image = image.unsqueeze(0)

    def preprocess_image(img):
        img = img.resize((300, 300))  # Resize to 224x224
        img_array = np.array(img)  # Convert to a NumPy array
        img_array = img_array / 255.0  # Normalize to [0, 1]
        img_array = np.expand_dims(img_array, axis=0)  # Add batch dimension
        return img_array

    image_tf = preprocess_image(image2)

    if torch.cuda.is_available():
        image = image.to("cuda")
        model.to("cuda")

    with torch.no_grad():
        output = model(image)
    probabilities = torch.nn.functional.softmax(output[0], dim=0)

    prob, cat = torch.topk(probabilities, 1)

    output_data = (prob.tolist()[0], cat.tolist()[0])

    if output_data[1] == 0:
        # Run plant model
        with open("data/models/Plant_encoder.pkl", "rb") as f:
            label_encoder = pickle.load(f)
        model2 = plant

    if output_data[1] == 1:
        # Run insect model
        with open("data/models/Insect_encoder.pkl", "rb") as f:
            label_encoder = pickle.load(f)
        model2 = insect

    elif output_data[1] == 2:
        # Run Aves(bird) model
        with open("data/models/Bird_encoder.pkl", "rb") as f:
            label_encoder = pickle.load(f)
        model2 = bird

    elif output_data[1] == 3:
        # Run Mammamal model
        with open("data/models/Mammal_encoder.pkl", "rb") as f:
            label_encoder = pickle.load(f)
        model2 = mammal
    else:
        return "error"

    predictions = model2.predict(image_tf)
    predicted_class_idx = np.argmax(predictions, axis=1)
    predicted_label = label_encoder.inverse_transform(predicted_class_idx)

    return predicted_label[0]


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5001)
