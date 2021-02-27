import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
from util import load_detections, abbreviate
from collections import defaultdict


def main():
    data = defaultdict(list)

    # Load all documents
    detections_db = load_detections()
    detection_docs = detections_db.find()

    for doc in detection_docs:
        # Resource per script
        data['RpS'].append(doc['numberOfResources'])
        # Line of Code
        data['LoC'].append(doc['lineOfCodes'])
        # Detection Result
        for key, value in doc['detectionResult'].items():
            data[abbreviate(key)].append(value)

    df = pd.DataFrame.from_dict(data)
    occurrences = df.loc[:, df.columns != 'LoC']
    df['total-occurrences'] = occurrences.sum(axis=1)

    fig, ax = plt.subplots(1, 2, figsize=(15, 15), sharex=True)
    fig.suptitle('Correlation Analysis')
    df.plot.scatter('LoC', 'total-occurrences', ax=ax[0])
    df.plot.scatter('RpS', 'total-occurrences', ax=ax[1])

    ax[0].set_xlabel('Smells Existence per Script')
    ax[0].set_xscale('log')
    ax[0].set_yscale('log')

    ax[1].set_xlabel('Resource Count per Script')
    ax[1].set_xscale('log')
    ax[1].set_yscale('log')

    plt.show()


if __name__ == '__main__':
    main()
