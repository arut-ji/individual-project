from collections import defaultdict
from util import abbreviate, load_detections
import matplotlib.pyplot as plt
import pandas as pd


def name() -> str:
    return 'effected-scripts'


def run():
    detections = load_detections()
    detection_results = list(
        map(
            lambda x: x['detectionResult'],
            list(detections.find())
        )
    )
    smells = detection_results[0].keys()
    effected_scripts = defaultdict()

    for smell in smells:
        num_scripts = len(list(filter(lambda x: x[smell] != 0, detection_results)))
        effected_scripts[abbreviate(smell)] = num_scripts
    fig, ax = plt.subplots(1, 1)

    ax.bar(effected_scripts.keys(), effected_scripts.values())
    ax.set_xlabel("Smell Category")
    ax.set_ylabel("Number of effected scripts")
    ax.set_title("Number of scripts effected by each smell (n = {})".format(len(detection_results)))
    return fig


def main():
    run()
    plt.show()


if __name__ == '__main__':
    main()
