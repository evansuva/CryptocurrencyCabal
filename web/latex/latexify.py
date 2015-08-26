#
# convert .md file for latex
#

import sys

def process_latex(finname):
    inheader = False
    with open(finname, 'r') as file:
        for line in file:
            if inheader:
                if line.startswith('---'):
                    inheader = False
                else:
                    if line.startswith('title:'):
                        startquote = line.find('"')
                        endquote = line.find('"', startquote + 1)
                        assert startquote != -1
                        assert endquote != -1
                        assert endquote > startquote
                        title = line[startquote + 1:endquote]
                        print "% " + title
                    elif line.startswith('date:'):
                        print "% " + line[6:],
                    else:
                        pass # ingore rest of header
            else:
                if line.startswith('---'):
                    inheader = True
                else:
                    print line,


if __name__ == "__main__":
    fin = sys.argv[1]
    process_latex(fin)
