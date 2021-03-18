import matplotlib.pyplot as plt
from util import load_detections_as_df


def main():
    run()
    plt.show()


def name() -> str:
    return "correlation-between-loc-and-rc"


def run():
    df = load_detections_as_df()
    occurrences = df.loc[:, df.columns != 'LoC']
    df['total-occurrences'] = occurrences.sum(axis=1)
    df = df.loc[:, df.columns != 'AC']
    print(df.head())
    fig, ax = plt.subplots(1, 2, figsize=(15, 7), sharex=True)
    fig.suptitle('Correlation between Smell Occurrences versus LoC and Resource Count')
    df.plot.scatter('LoC', 'total-occurrences', ax=ax[0])
    df.plot.scatter('RpS', 'total-occurrences', ax=ax[1])

    ax[0].set_xlabel('Line-of-code per Script')
    ax[0].set_xscale('log')
    ax[0].set_yscale('log')
    ax[0].set_ylabel('Occurrences')

    ax[1].set_xlabel('Resource Count per Script')
    ax[1].set_xscale('log')
    ax[1].set_yscale('log')
    ax[1].set_ylabel('Occurrences')
    return fig


if __name__ == '__main__':
    main()
