# text-search-engine
A text search engine built using GO

### Download the data dump 
[data-dump](https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz)

### Data structure 
* Title  
* URL 
* Abstract

### Basic approaches
* strings.Contains
* regexp 
* matchstring 

### Problem with basic approaches
The problem is that they don't scale
Takes upto 2 seconds for 600k docs , but what if we have 10M docs ? 
The time will keep increasing 

### Solution 
We will use the approach called `inverted index` 
* We will pre-process the data and create inverted-index from the text.
* We will keep a track of each word and its existence document wise. 
* We do [Tokenization](https://www.mckinsey.com/featured-insights/mckinsey-explainers/what-is-tokenization)
* Then we will do filtering (lowercasing , dropping common words, stemming).
* Lastly we will do searching. 
* We don't go through the docs for searching we will simply search the index.
