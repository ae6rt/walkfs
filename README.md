A simple filesystem walker with depth control, similar to 

```
$ find DIR -maxdepth N -name T"
```

```
$ walkfs -h
Usage of walkfs:
  -filename string
    	file to search for
  -max-depth int
    	max depth (default 1)
  -root-dir string
    	root path (default ".")
```
