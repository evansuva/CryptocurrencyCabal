all: ../public/syllabus.pdf ../public/class1.pdf ../public/class2.pdf ../public/class3.pdf checkup1.pdf ../public/class5.pdf ../public/class6.pdf ../public/class7.pdf ../public/class8.pdf checkup2.pdf

checkup1.pdf: ../../staff/checkup1/checkup1.md
	python latexify.py ../../staff/checkup1/checkup1.md > checkup1-latex.md
	pandoc --template=./cabal.template checkup1-latex.md -o checkup1.pdf

checkup2.pdf: checkup2.md
	python latexify.py checkup2.md > checkup2-latex.md
	pandoc --template=./cabal.template checkup2-latex.md -o checkup2.pdf

feedback.pdf: ../../staff/checkup1/feedback.md
	python latexify.py ../../staff/checkup1/feedback.md > feedback-latex.md
	pandoc --template=./cabal.template feedback-latex.md -o feedback.pdf

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

../public/class5.pdf: ../content/classes/class5-latex.md
	python latexify.py ../content/classes/class5-latex.md > class5-latex.md
	pandoc --template=./cabal.template class5-latex.md -o ../public/class5.pdf

../public/class6.pdf: ../content/classes/class6.md
	python latexify.py ../content/classes/class6.md > class6-latex.md
	pandoc --template=./cabal.template class6-latex.md -o ../public/class6.pdf

../public/class7.pdf: ../content/classes/class7.md
	python latexify.py ../content/classes/class7.md > class7-latex.md
	pandoc --template=./cabal.template class7-latex.md -o ../public/class7.pdf

../public/class8.pdf: class8-latex.md
	# python latexify.py ../content/classes/class8.md > class8-latex.md
	pandoc --template=./cabal.template class8-latex.md -o ../public/class8.pdf

### Do not update class3.pdf
# ../public/class3.pdf: ../content/classes/class3.md
#	python latexify.py ../content/classes/class3.md > class3-latex.md
#	pandoc --template=./cabal.template class3-latex.md -o ../public/class3.pdf
