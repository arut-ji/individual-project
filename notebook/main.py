import correlation_between_categories as cbc
import correlation_between_loc_and_rc as cblac
import count_smells_instances as csi
import distinct_smells_per_script as dsps
import effected_scripts as es
import scripts_characteristics as sc
import smells_co_occurrence as sco


def main():
    analyses = [cbc, cblac, csi, dsps, es, sc, sco]
    dir_path = '../assets/'
    for analysis in analyses:
        figs = analysis.run()
        if isinstance(figs, list):
            for i in range(len(figs)):
                filename = dir_path + '{}-{}.png'.format(analysis.name(), i)
                figs[i].savefig(filename)
        else:
            filename = analysis.name() + '.png'
            figs.savefig(dir_path + filename)


if __name__ == '__main__':
    main()
