# ShopSearcher

this is a simple command line tool to query multiple websites/shop  
for a search-term.   
Querying means it opens the urls defined in the config with the entered search term.   
the default browser is used.

there is one pre defined config found in `search_config/data-embedded`  
this config will be extracted to a `./data` folder relativ to where your binary is.   

## pre defined configs
- swiss_wine (searches swiss wine shops)

## Usage
the CLI currently supports the following arguments.

- --config   
must be a config file found in './data', the file-ending '.txt' can be ommited.
if you omit it all configs will be used.
  all arguments which are followed are used as a search term
  
### Examples (for windows)
`shop-search.exe Gantenbein`   
searches all configs for the term Gantenbein

`shop-search.exe --config swiss_wine Gantenbein`   
searches the config swiss_wine.txt for Gantenbein

