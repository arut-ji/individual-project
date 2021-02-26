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
    corr = df.corr()

    # TODO: loc-smells-density plot
    print(df.loc[:, df.columns != 'LoC'].head())

    # sns.heatmap(corr, annot=True)
    # plt.show()


if __name__ == '__main__':
    main()
