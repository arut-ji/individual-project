from pymongo import MongoClient
from collections import defaultdict
import pandas as pd


def abbreviate(s: str) -> str:
    words = s.split('-')
    return ''.join(map(lambda x: x[0].upper(), words))


def load_detections():
    client = MongoClient('localhost', 27017)
    return client['kubernetes']['detections']


def load_detections_as_df() -> pd.DataFrame:
    data = defaultdict(list)
    # Load all documents
    detections_db = load_detections()
    docs = detections_db.find()

    for doc in docs:
        # Resource per script
        data['RpS'].append(doc['numberOfResources'])
        # Line of Code
        data['LoC'].append(doc['lineOfCodes'])
        # Detection Result
        for key, value in doc['detectionResult'].items():
            data[abbreviate(key)].append(value)

    return pd.DataFrame.from_dict(data)
