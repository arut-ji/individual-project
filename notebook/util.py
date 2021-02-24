from pymongo import MongoClient


def abbreviate(s: str) -> str:
    words = s.split('-')
    return ''.join(map(lambda x: x[0].upper(), words))


def load_detections():
    client = MongoClient('localhost', 27017)
    return client['kubernetes']['detections']
