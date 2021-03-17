import pandas as pd
import matplotlib.pyplot as plt
from util import load_detections_as_df
import seaborn as sns

sns.set_theme()


def main():
    df = load_detections_as_df()
    df = df.loc[:, (df.columns != 'LoC') & (df.columns != 'RpS')]

    fig, ax = plt.subplots(1, 2, figsize=(10, 5))
    fig.suptitle('Correlation between Smells Category')
    plot_heatmap(ax[0], df)
    plot_heatmap(ax[1], df.loc[:, df.columns != 'AC'])
    plt.show()


def plot_heatmap(ax, df):
    corr = df.corr()
    sns.heatmap(corr, ax=ax)


if __name__ == '__main__':
    main()
