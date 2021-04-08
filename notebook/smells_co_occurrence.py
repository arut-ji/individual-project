import matplotlib.pyplot as plt
from util import load_detections_as_df
import seaborn as sns


def name() -> str:
    return 'smells-co-occurrence'


def run():
    sns.set_theme()
    df = load_detections_as_df()
    df = df.loc[:, (df.columns != 'LoC') & (df.columns != 'RpS')]
    pg = sns.pairplot(df, kind="scatter", diag_kind="kde")
    pg.set(xscale="log")
    pg.tight_layout()
    return pg.fig


def main():
    run()
    plt.show()


if __name__ == '__main__':
    main()
