GITHUB_PAGES_BRANCH=gh-pages
SITENAME=bitcoin-class.org

github: 
	cd web; make html
	cd web/latex; make
	echo $(SITENAME) > web/public/CNAME
	git add -A
	git commit -m "Rebuilt site"
	git push origin master
	git subtree push --prefix=web/public ssh://git@github.com/evansuva/CryptocurrencyCabal.git gh-pages
git subtree push --prefix=web/public https://github.com/evansuva/CryptocurrencyCabal.git gh-pages

.PHONY: html clean develop
