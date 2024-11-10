import torch
import torch.nn as nn
from flask import Flask, request, jsonify
import json
import numpy as np
import torchvision.models as models
from torchvision.transforms import transforms
from PIL import Image


model = models.mobilenet_v2(pretrained=False)
model.classifier[1] = nn.Linear(model.classifier[1].in_features, 4)
model.load_state_dict(torch.load("data/mobilenetft.pth"))
model.eval()


app = Flask(__name__)


categories = {0: "Plantae", 1: "Insecta", 2: "Aves", 3: "Mammalia"}


@app.route("/inference", methods=["POST"])
def inference():
    data = request.files["image"].stream
    image = Image.open(data).convert("RGB")

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
        pass
    elif output_data[1] == 1:
        # Run insect model
        pass
    elif output_data[1] == 2:
        # Run Aves(bird) model
        pass
    else:
        # Run Mammamal model
        pass

    # return Scientitic name in the output
    return "output"


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5001)
