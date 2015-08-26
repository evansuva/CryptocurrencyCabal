all: syllabus class1

syllabus:
	python latexify.py ../content/syllabus.md > syllabus-latex.md
	pandoc --template=./cabal.template syllabus-latex.md -o syllabus.pdf
	mv syllabus.pdf ../public/

class1:
	python latexify.py ../content/classes/class1.md > class1-latex.md
	pandoc --template=./cabal.template class1-latex.md -o ../public/class1.pdf
