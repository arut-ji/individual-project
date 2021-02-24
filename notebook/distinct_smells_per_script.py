from util import load_detections
import matplotlib.pyplot as plt


def main():
    detections_db = load_detections()
    detection_results = list(
        map(
            lambda x: x['detectionResult'],
            list(detections_db.find())
        )
    )
    distinct_smells = list(
        map(
            lambda detection: len(list(filter(lambda x: x != 0, detection.values()))),
            detection_results
        )
    )
    plt.hist(distinct_smells, bins=8)
    plt.title('Distinct Smells per script')
    plt.ylabel('Frequency')
    plt.xlabel('Number of Distinct Smells per script')
    plt.show()


if __name__ == '__main__':
    main()
