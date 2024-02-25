## Modules

Each folder and its files(top-level or immediate) must have same package name;
like src folder should have one and only one package named `anythingYouLike`.

- If there more than one package inside top-level files of a foder; go will not work
- For simplicity; use package name same as folder name
- Since, package name cannot contain dashes or other symbols; name you folder like that
- thus, prefer camelCase for folder and package names


## public/private
- func or any other top level values defined with first letter as capital letter will be publicly available in other packages or folders.
- same folder contains same package but as mentioned a folder can have 10 files at top-level with package name; so all these files can use camelCase func, values in other files of same package.

folder: src; files: geography.go(package src) and users.go(package src)
geography.go has 2 functions: GetCountry, getCity
- GetCountry is available everywhere
- getCity is available in geography.go and users.go; since both share same package