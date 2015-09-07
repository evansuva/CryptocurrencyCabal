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
                line = line.replace('&rarr;', '$\\rightarrow$')
                line = line.replace('KU<sub>X</sub>', '$KU_{X}$')
                line = line.replace('KR<sub>X</sub>', '$KR_{X}$')
                line = line.replace('_k_<sub>i</sub>', '$k_i$')
                line = line.replace('<!--page-->', '\clearpage')

                if line.startswith('---'):
                    inheader = True
                elif line.startswith('<div class="gap"></div>'):
                    print "\n#\n"
                elif line.startswith('<div class="biggap"></div>'):
                    print "\n#\n#\n"
                elif line.startswith('<div class="biggergap"></div>'):
                    print "\n#\n#\n#\n#\n"
                else:
                    print line,


if __name__ == "__main__":
    fin = sys.argv[1]
    process_latex(fin)
