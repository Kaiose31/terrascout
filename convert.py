import pandas as pd
import json

df = pd.read_csv("data/observations-497968.csv")
df = df.fillna("", axis=1)
db = df.to_dict(orient="records")
with open("data/obs.json", "w") as f:
    json.dump(db, f, indent=4)
