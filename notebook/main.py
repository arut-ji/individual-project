import correlation_between_categories as cbc
import matplotlib.pyplot as plt


def main():
    analyses = [cbc]
    for analysis in analyses:
        analysis.run()
        dir_path = './figures/'
        filename = analysis.name() + ".pdf"
        plt.savefig(dir_path + filename)


if __name__ == '__main__':
    main()
