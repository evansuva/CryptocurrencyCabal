all: ../public/syllabus.pdf ../public/class1.pdf ../public/class2.pdf

../public/syllabus.pdf: ../content/syllabus.md
	python latexify.py ../content/syllabus.md > syllabus-latex.md
	pandoc --template=./cabal.template syllabus-latex.md -o syllabus.pdf
	mv syllabus.pdf ../public/

../public/class1.pdf: ../content/classes/class1.md
	python latexify.py ../content/classes/class1.md > class1-latex.md
	pandoc --template=./cabal.template class1-latex.md -o ../public/class1.pdf

../public/class2.pdf: ../content/classes/class2.md
	python latexify.py ../content/classes/class2.md > class2-latex.md
	pandoc --template=./cabal.template class2-latex.md -o ../public/class2.pdf
