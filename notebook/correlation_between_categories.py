import pandas as pd
import matplotlib.pyplot as plt
from util import load_detections_as_df
import seaborn as sns

sns.set_theme()


def main():
    run()
    plt.show()


def name() -> str:
    return "correlation-between-categories"


def run():
    df = load_detections_as_df()
    df = df.loc[:, (df.columns != 'LoC') & (df.columns != 'RpS')]

    fig, ax = plt.subplots(1, 1, figsize=(10, 5))
    fig.suptitle('Correlation between Smells Category')
    plot_heatmap(ax, df)
    return fig


def plot_heatmap(ax, df):
    corr = df.corr()
    sns.heatmap(corr, ax=ax)


if __name__ == '__main__':
    main()
