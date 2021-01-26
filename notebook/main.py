from pymongo import MongoClient
import matplotlib.pyplot as plt

client = MongoClient('localhost', 27017)

db = client['kubernetes']

detections = db['detections']

detection_results = list(
    map(
        lambda x: x['detectionResult'],
        list(detections.find())
    )
)

smells = detection_results[0].keys()

occurrences = {}

for smell in smells:
    count = len(list(filter(lambda x: x[smell] is True, detection_results)))
    occurrences[smell] = count

x_pos = [i for i, _ in enumerate(occurrences.keys())]

plt.bar(occurrences.keys(), occurrences.values())
plt.xlabel("Smell Category")
plt.ylabel("Occurrence")
plt.title("Occurrences for each code smells (n = {})".format(len(detection_results)))
plt.show()
