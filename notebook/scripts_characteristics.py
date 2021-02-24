from util import load_detections
import matplotlib.pyplot as plt
import numpy as np


def main():
    detections = load_detections()
    loc = np.array(
        list(
            map(
                lambda x: x['lineOfCodes'],
                list(detections.find())
            )
        )
    )
    hist, bins = np.histogram(loc, bins=8)
    logbins = np.logspace(np.log10(bins[0]), np.log10(bins[-1]), len(bins))
    (fig, ax) = plt.subplots(1, 1)
    ax.hist(loc, bins=logbins)
    ax.set_title(r'$\mathrm{Line-of-code \/ Histogram}$')
    ax.set_ylabel(r"$Frequency$")
    ax.set_xlabel(r'$log_{10}\/(LOC)$')
    ax.set_xscale('log')
    plt.show()


if __name__ == '__main__':
    main()
