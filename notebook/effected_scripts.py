from pymongo import MongoClient
from collections import defaultdict
from util import abbreviate, load_detections
import matplotlib.pyplot as plt


def main():
    detections = load_detections()
    detection_results = list(
        map(
            lambda x: x['detectionResult'],
            list(detections.find())
        )
    )
    smells = detection_results[0].keys()
    effected_scripts = defaultdict()

    for smell in smells:
        num_scripts = len(list(filter(lambda x: x[smell] != 0, detection_results)))
        effected_scripts[abbreviate(smell)] = num_scripts
    print(effected_scripts)

    plt.bar(effected_scripts.keys(), effected_scripts.values())
    plt.xlabel("Smell Category")
    plt.ylabel("Number of effected scripts")
    plt.title("Number of scripts effected by each smell (n = {})".format(len(detection_results)))
    plt.show()


if __name__ == '__main__':
    main()
