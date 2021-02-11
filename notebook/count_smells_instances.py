from pymongo import MongoClient
import matplotlib.pyplot as plt
from collections import defaultdict
from util import abbreviate

client = MongoClient('localhost', 27017)

db = client['kubernetes']
detections = db['detections']
detection_results = list(
    map(
        lambda x: x['detectionResult'],
        list(detections.find())
    )
)

totalLoC = sum(map(
    lambda x: x['lineOfCodes'],
    list(detections.find())
))

print(totalLoC)

smells = detection_results[0].keys()

occurrences = defaultdict()

for smell in smells:
    count = sum(map(lambda x: x[smell], detection_results))
    occurrences[abbreviate(smell)] = count

print(occurrences)

x_pos = [i for i, _ in enumerate(occurrences.keys())]

plt.bar(occurrences.keys(), occurrences.values())
plt.xlabel("Smell Category")
plt.ylabel("Occurrence")
plt.title("Occurrences for each code smells (n = {})".format(len(detection_results)))
plt.show()
