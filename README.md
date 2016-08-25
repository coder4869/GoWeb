# GoWeb


# Introduction
This WEB is an go web demo, now including WEB API used and mongodb DataBase.

Project Directory:
	|--GoWeb
		|--	src
			|-- WebMain (entry of the Project)
			|-- conf (config operations for DataBase)
			|-- core (business operation)
				|-- controllers 
				|-- defines 
				|-- mgodaos (including init of mongodb config)
				|-- models 
				|-- mysqldaos (including init of mysql DataBase config)
				|-- serves 
			|-- libs (depends libs)
			|-- log (business logs)
			|-- tests (test cases)	
		|--	docs

Function Call:
    WebMain --> serves --> controllers --> daos(mgodaos/mysqldaos) --> models & conf

Developers:
    coder4869@gmail.com

Depends:
    "github.com/coder4869/golibs"
    "gopkg.in/mgo.v2"


# Run & Deployment
    Refer to "docs/Run&Deploy.md", this document provides the modify points 
    for run and deploy the project. 


# More
Other README.md or docs


