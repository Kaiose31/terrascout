import pandas as pd
import json
import shutil
import requests
import multiprocessing
import requests
from typing import List
import time


def downloadtaxon(taxon_ids: List[int]):
    all_results = []

    for i in range(0, len(taxon_ids), 30):
        print(i)
        batch = taxon_ids[i : i + 30]
        resp = requests.get(f"https://api.inaturalist.org/v1/taxa/{",".join([str(i) for i in batch])}")

        j = resp.json()

        for res in j["results"]:
            result = {"id": res["id"], "wikipedia_summary": res["wikipedia_summary"]}
            all_results.append(result)
        time.sleep(1)

    with open("data/taxon.json", "w") as f:
        json.dump(all_results, f, indent=4)


def convert(df):

    df = df.fillna("", axis=1)
    db = df.to_dict(orient="records")
    with open("data/obs_small.json", "w") as f:
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
    df = pd.read_csv("data/observation_large.csv")
    # df = df.groupby("iconic_taxon_name").head(100)
    # df.to_csv("data/obs_small.csv")

    # convert(df)

    downloadtaxon(df["taxon_id"].unique().tolist())
