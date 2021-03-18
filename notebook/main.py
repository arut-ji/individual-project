import correlation_between_categories as cbc
import correlation_between_loc_and_rc as cblac
import count_smells_instances as csi
import distinct_smells_per_script as dsps
import effected_scripts as es
import scripts_characteristics as sc
import smells_co_occurrence as sco
import matplotlib.pyplot as plt


def main():
    analyses = [cbc, cblac, csi, dsps, es, sc, sco]
    for analysis in analyses:
        analysis.run()
        dir_path = '../assets/'
        filename = analysis.name() + ".png"
        plt.savefig(dir_path + filename)


if __name__ == '__main__':
    main()
