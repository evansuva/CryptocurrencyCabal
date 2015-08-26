all:
	pandoc --template=../latex/cabal.template syllabus-latex.md -o syllabus.pdf
