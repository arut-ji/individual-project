import pandas as pd
import matplotlib.pyplot as plt
from util import load_detections_as_df
import seaborn as sns

sns.set_theme()


def main():
    df = load_detections_as_df()
    occurrences_df = df.loc[:, (df.columns != 'LoC') & (df.columns != 'RpS')]

    corr = occurrences_df.corr()

    fig, ax = plt.subplots(1, 1)
    fig.suptitle('Correlation between Smells Category')
    ax = sns.heatmap(corr)
    plt.show()


if __name__ == '__main__':
    main()
