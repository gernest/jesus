# jesus [![Build Status](https://travis-ci.org/gernest/jesus.svg?branch=master)](https://travis-ci.org/gernest/jesus)

Work with M-Pesa messages froma `gammu-smsd` daemon

Docummentation
_____________

Check it here [Docs](http://gernest.github.io/jesus)_


Features
_________

* Compatible with the smsd schema
* Filters m-pesa sms for important components
* Exetensible, you can extend the API whatever you like
* multiple database supports, postgresql, mysql, sqlite and many more

Installation
------------

Install Jesus using the "go get" command:

    go get github.com/gernest/jesus

Jesus depends on the `gorm` ORM linrary for maintaining of Database interactions

Contributing
------------

Contributions are welcome. 

Before writing code, send mail to geofreyernest@live.com to discuss what you
plan to do. This gives me a chance to validate the design, avoid duplication of
effort and ensure that the changes fit the goals of the project. Do not start
the discussion with a pull request. 

License
-------

jesus is available under the 
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)
