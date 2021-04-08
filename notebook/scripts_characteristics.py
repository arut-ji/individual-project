from util import load_detections
import matplotlib.pyplot as plt
import numpy as np
import pandas as pd


def name():
    return 'scripts-characteristics'


def histplot(data, title: str):
    hist, bins = np.histogram(data, bins=15)
    log_bins = np.logspace(np.log10(bins[0]), np.log10(bins[-1]), len(bins))
    fig, ax = plt.subplots(1, 1)
    ax.hist(data, bins=log_bins)
    ax.set_xscale('log')
    ax.set_title(title)
    return fig


def boxplot(data, title: str):
    fig, ax = plt.subplots(1, 1)
    ax.boxplot(data, vert=False)
    ax.set_xscale('log')
    ax.set_title(title)
    return fig


def run():
    detections_db = load_detections()
    detections = list(detections_db.find())
    # Line of Codes
    loc = np.array(list(map(lambda x: x['lineOfCodes'], detections)))
    # Number of Resources
    nor = np.array(list(map(lambda x: x['numberOfResources'], detections)))

    df = pd.DataFrame.from_records(list(zip(loc, nor)))
    print(df.describe())
    print(df.corr())

    loc_hist = histplot(loc, "Line-of-code Distribution")
    nor_hist = histplot(nor, "Resources Count")

    fig, ax = plt.subplots(1, 1)
    ax.scatter(loc, nor)
    ax.set_ylabel("Resources Count")
    ax.set_xlabel("Line-of-Code")
    ax.set_xscale('log')
    ax.set_yscale('log')
    ax.set_title("Line-of-Code versus Resources Count")

    return [loc_hist, nor_hist, fig]


def main():
    run()
    plt.show()


if __name__ == '__main__':
    main()
