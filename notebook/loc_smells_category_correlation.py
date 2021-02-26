import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
from util import load_detections, abbreviate
from collections import defaultdict


def main():
    data = defaultdict(list)

    detections_db = load_detections()
    detection_docs = detections_db.find()
    for doc in detection_docs:
        data['LoC'].append(doc['lineOfCodes'])
        for key, value in doc['detectionResult'].items():
            data[abbreviate(key)].append(value)

    df = pd.DataFrame.from_dict(data)
    occurrences = df.loc[:, df.columns != 'LoC']
    df['total-occurrences'] = occurrences.sum(axis=1)

    fig, ax = plt.subplots(1, 1)
    df.plot.scatter('LoC', 'total-occurrences', ax=ax)

    ax.set_ylabel("Smells Existence per Script")

    fig.show()


if __name__ == '__main__':
    main()
