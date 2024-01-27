# jsonkeys

Do you for some reason need a list of all the nested keys in a very large (so hard to jq) JSON file? 

So for:

```
{
    "a": "one",
    "b": "two",
    "c": [
	    {
		"a2": [{"a3": 1}]
	    }
    ],
    "d" : {"e": [{"xxx":1}], "f": [{"xxx":2}]}
}
```

You get

```
a 
b
c
d
d.e
d.f
```

Note that we don't delve into arrays. If somehow your usage matches this is the place!

# Compiling / usage

To install this you'll need go, if you have this checked out you should be able to use go build, but otherwise:

```
go install github.com/draxil/jsonkeys@latest
```

You may need to add go's bin file to your path, see general instructions on this! If this gets some interest I'll add builds.

Then:

```
jsonkeys largefile.json
```

Or 
```
cat largefile.json | jsonkeys
```

# Notes

+ Does not delve into arrays.
+ Does not cope with ndjson/nljson.

# Status


Basically a thing I've knocked together to answer a query on [json2nd](https://github.com/draxil/json2nd)! I actually now find it a bit handy when looking at a new JSON file I don't know. Like json2nd it should not choke on large files, although the implmentation here is a bit more go standard library. I may expand this if anyone cares, so let me know if it's useful.

