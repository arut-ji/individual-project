from util import load_detections
import matplotlib.pyplot as plt
import numpy as np

def name():
    return 'scripts-characteristics'

def run():
    detections_db = load_detections()
    detections = list(detections_db.find())
    # Line of Codes
    loc = np.array(list(map(lambda x: x['lineOfCodes'], detections)))
    # Number of Resources
    nor = np.array(list(map(lambda x: x['numberOfResources'], detections)))
    hist, bins = np.histogram(loc, bins=15)
    log_bins = np.logspace(np.log10(bins[0]), np.log10(bins[-1]), len(bins))

    fig, (ax1, ax2) = plt.subplots(2, 1)
    ax1.hist(loc, bins=log_bins)
    ax1.set_ylabel(r"$Frequency$")
    ax1.set_xscale('log')
    ax1.title.set_text("Line of Code")

    ax2.hist(nor, bins=log_bins)
    ax2.set_ylabel(r"$Frequency$")
    ax2.set_xscale('log')
    ax2.title.set_text("Number of Resources")
    fig.tight_layout()
    return fig


def main():
    run()
    plt.show()


if __name__ == '__main__':
    main()
