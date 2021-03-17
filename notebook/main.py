import correlation_between_categories as cbc
import correlation_between_loc_and_rc as cblac
import count_smells_instances as csi
import matplotlib.pyplot as plt


def main():
    analyses = [cbc, cblac, csi]
    for analysis in analyses:
        analysis.run()
        dir_path = './figures/'
        filename = analysis.name() + ".pdf"
        plt.savefig(dir_path + filename)


if __name__ == '__main__':
    main()
