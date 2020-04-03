## *How To Use generate-cls.sh*

### *PreCondition*
- [ ] [doxygen](http://www.doxygen.nl/)
- [ ] [graphviz](https://www.graphviz.org/)

*showing usage*
```sh
$ ~/generate-cls.sh -h
Usage: /Users/hacknatural/generate-cls.sh [-d]
[-d]:set your root project dir
```

*execute exapmle*
```sh
$ ~/generate-cls.sh -d ~/test_dir/source_prj
```

*After executing the script, DoxyFile is generated under the specified directory, so if customization is required, modify the template and execute the doxygen command*

```sh
$ doxygen
```
