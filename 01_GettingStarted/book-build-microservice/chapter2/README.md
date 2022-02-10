This chapter 2, you need to know some important keyword:
1. How do you oranginze data in Golang?
2. Apply abstraction words in MongoDB into Golang:  
    - **Library - Mongo**: Library includes all knowledge about the worlds, this keyword make you imagine that you need match all knowledge in your life into library.
    - **Sections - Databases**: Section is a way to organize the specific knowledge into a section.
    - **Book - Collections**: Each section will have many kind of books help you understand deeply about major.
    - **Pages - Documents**: Each book will have many pages help you to recognize about features, characteristics on book.  
    **Note:**   
    Let's get a sample for your imagination
      - Library contains many kind of books, technoly, math, history ...
      - And we have a book about aminal, and aniaml is called a section.
      - In animals, we have many different species, each specie is called a book.
      - For each specie, we investigate their habitat, breed in animals.. it is called a page.  
3. How Golang applies the above keywords to effectively implement REST API?

## Table of contents
* [RESTful APIs](#RESTfull-APIs) 
  * [URI path design for REST services](#URI-path-design-for-REST-services) 
    * [Collections](#Collections)
    * [Documents](#Documents)
    * [Controller](#Controller)
    * [Store](#Store)
    * [CRUD function names](#CRUD-function-names)
  * [HTTP verbs](#HTTP-verbs)
  * [URI query design](#URI-query-design)
  * [Response code](#Response-code)
* [Accessing APIs from JavaScript](#Access-Javascript)

## Questions
## RESTfull-APIs

### URI-path-design-for-REST-services
In golang, we have rule to desin the URI path for REST service into 
- Collections
- Documents
- Stores
- Controllers
Therefore, please remember rule to design it.

#### Collections
```
  GET /cats -> All cats in the collection
  GET /cats/1 -> Single document for a cat 1
```
#### Documents
```
  GET /cats/1 -> Single document for cat 1
  GET /cats/1/kittens -> All kittens belonging to cat 1
  GET /cats/1/kittens/1 -> Kitten 1 for cat 1
```
#### Controller
```
  POST /cats/1/feed -> Feed cat 1
  POST /cats/1/feed?food=fish ->Feed cat 1 a fish
```
#### Store
```
  PUT /cats/2
```
#### CRUD-function-names
```
In java, we often use:
  DELETE /cats/1234

In golang, we must use:
  GET /deleteCat/1234
  DELETE /deleteCat/1234
  POST /cats/1234/delete
```

### HTTP-verbs

### URI-query-design

### Response-code

## Access-Javascript
