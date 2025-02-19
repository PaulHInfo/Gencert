# Gencert
This is perhaps the worst way in the world to practise golang.  This is a certificate generator.

## Way to use it
First you need to install go and gofpdf : 
```go get github.com/jung-kurt/gofpdf```

Then in the root folder to creat the .exe :
```go build```

You need to write your data in a .csv file :
```<Course's name>,<Student's name>,<Date> ```

for the date use this format : 
Year - Month - Day -> **2024-09-19**




Here's how to make it work :
```.\gencert.exe -type <type> -file <outputfile> ```

## Note
This isn't a serious project and it's part of my golang apprenticeship. In my spare time I'm planning to add some new feature : Adding certificates in JSON / XML format or make the HTML rendering more attractive.

