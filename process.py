import pandas as pd
import json
import shutil
import requests
import multiprocessing


def convert(df):
    df = df.fillna("", axis=1)
    db = df.to_dict(orient="records")
    with open("data/obs.json", "w") as f:
        json.dump(db, f, indent=4)


def download(df: pd.DataFrame):
    df = df[["id", "image_url"]]
    for idx, row in df.iterrows():
        url = row["image_url"]
        id = row["id"]
        ext = url.split(".")[-1]
        res = requests.get(url)
        with open(f"data/images/{id}.jpg", "wb") as f:
            f.write(res.content)


if __name__ == "__main__":
    df = pd.read_csv("data/observations-497968.csv")
    # download(df)
