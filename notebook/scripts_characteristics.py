from util import load_detections
import matplotlib.pyplot as plt
import numpy as np


def main():
    detections_db = load_detections()
    detections = list(detections_db.find())
    # Line of Codes
    loc = np.array(
        list(
            map(
                lambda x: x['lineOfCodes'],
                detections
            )
        )
    )
    # Number of Resources
    nor = np.array(
        list(
            map(
                lambda x: x['numberOfResources'],
                detections
            )
        )
    )
    hist, bins = np.histogram(loc, bins=15)
    logbins = np.logspace(np.log10(bins[0]), np.log10(bins[-1]), len(bins))

    fig = plt.figure()

    ax1 = fig.add_subplot(211)
    ax1.hist(loc, bins=logbins)
    ax1.set_ylabel(r"$Frequency$")
    ax1.set_xscale('log')
    # ax1.title.set_text("Line of Code")

    ax2 = fig.add_subplot(212)
    ax2.hist(nor, bins=logbins)
    ax2.set_ylabel(r"$Frequency$")
    ax2.set_xscale('log')
    # ax2.title.set_text("Number of Resources")

    plt.show()


if __name__ == '__main__':
    main()
