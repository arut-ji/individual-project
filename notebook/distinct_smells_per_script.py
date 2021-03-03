from util import load_detections, load_detections_as_df
import matplotlib.pyplot as plt


def main():
    df = load_detections_as_df()
    occurrences_df = df.loc[:, (df.columns != 'LoC') & (df.columns != 'RpS')]
    distinct_smells_df = occurrences_df.apply(
        lambda x: sum(list(map(lambda item: 0 if item == 0 else 1, x))),
        axis=1,
        result_type='expand'
    ).value_counts()
    print(distinct_smells_df.head())
    distinct_smells_df.plot(kind='bar')
    plt.show()


if __name__ == '__main__':
    main()
