from pymongo import MongoClient
import matplotlib.pyplot as plt
from collections import defaultdict
from util import abbreviate


def main():
    run()
    plt.show()


def name() -> str:
    return 'smells-instance-count'


def run():
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

    occurrences = defaultdict()

    for smell in smells:
        count = sum(map(lambda x: x[smell], detection_results))
        occurrences[abbreviate(smell)] = count

    fig, ax = plt.subplots(1, 1)
    print(occurrences)
    ax.bar(occurrences.keys(), occurrences.values())
    ax.set_xlabel("Smell Category")
    ax.set_yscale('log')
    ax.set_ylabel("Occurrence")
    ax.set_title("Occurrences for each code smells (n = {})".format(len(detection_results)))
    return fig


if __name__ == '__main__':
    main()
