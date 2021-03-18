from util import load_detections_as_df
import matplotlib.pyplot as plt


def name() -> str:
    return 'distinct-smells-per-scripts'


def run():
    df = load_detections_as_df()
    occurrences_df = df.loc[:, (df.columns != 'LoC') & (df.columns != 'RpS')]
    distinct_smells_df = occurrences_df.apply(
        lambda x: sum(list(map(lambda item: 0 if item == 0 else 1, x))),
        axis=1,
        result_type='expand'
    ).value_counts(sort=False)
    fig, _ = plt.subplots(1, 1)
    ax = distinct_smells_df.plot(kind='bar')
    ax.set_xlabel('Number of distinct smells')
    ax.set_ylabel('Frequency')
    ax.set_title('Distinct smells count in each script')
    return fig


def main():
    run()
    plt.show()


if __name__ == '__main__':
    main()
