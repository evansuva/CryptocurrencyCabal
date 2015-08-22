GITHUB_PAGES_BRANCH=gh-pages
SITENAME=bitcoin-class.org

html:
	hugo --theme=cabal

develop:
	hugo server --theme=cabal --watch

github: 
	cp -r spring2015 public/
	make html
	git add -A
	git commit -m "Rebuilt site"
	git push origin master
	git subtree push --prefix=public https://evansuva@github.com/CryptocurrencyCabal.git gh-pages

.PHONY: html clean develop
