# jsontheexperimenter

Helper for generating experiments from a JSON file

A simple example that shows how to use "CommandFormat", a single value arg, an array value arg, a range value arg, and nested value args.

```json
{
	"Name" : "Experiment1",

	"CommentPrefix" : "# ",

	"Executable" : "MyExecutable",

	"OutputFile" : "NUMBERED",

	"OutputDirectory" : "data",
	
	"CommandFormat" : {
		"Format" : "%v %v > %v/%v",
		"Args" : [ "Executable", "Args", "OutputDirectory", "OutputFile"  ]
	},

	"Args" : {
		
		"Timeout" : {
			"Format" : "timeout=%v",
			"Value" : 30
		},

		"Seed" : {
			"Format" : "seed=%v",
			"Value" : {
				"Min" : 1,
				"Max" : 5,
				"Increment" : 1
			}
		},

		"Algorithms" : {
			"Format" : "-algorithm=%v",
			"Value" : [ "A*", "Dijkstra", "Greedy" ]
		},

		"Domains" : {
			"Format" : "-domain=%v",
			"Value" : [
				{
					"Value" : "gridnav",
					"Args" :{
						"Width" : {
							"Format" : "width=%v",
							"Value" : 10
						},
						"Height" : {
							"Format" : "heigt=%v",
							"Value" : 10
						}
					}
				},
				{ "Value" : "tiles" }
			]
		}

	}
}
```

Will yield:


```
# Experiment1
MyExecutable -algorithm=A* -domain=gridnav width=10 heigt=10 timeout=30 seed=1 > data/0
MyExecutable -algorithm=A* -domain=gridnav width=10 heigt=10 timeout=30 seed=2 > data/1
MyExecutable -algorithm=A* -domain=gridnav width=10 heigt=10 timeout=30 seed=3 > data/2
MyExecutable -algorithm=A* -domain=gridnav width=10 heigt=10 timeout=30 seed=4 > data/3
MyExecutable -algorithm=A* -domain=gridnav width=10 heigt=10 timeout=30 seed=5 > data/4
MyExecutable -algorithm=A* -domain=tiles timeout=30 seed=1 > data/5
MyExecutable -algorithm=A* -domain=tiles timeout=30 seed=2 > data/6
MyExecutable -algorithm=A* -domain=tiles timeout=30 seed=3 > data/7
MyExecutable -algorithm=A* -domain=tiles timeout=30 seed=4 > data/8
MyExecutable -algorithm=A* -domain=tiles timeout=30 seed=5 > data/9
MyExecutable -algorithm=Dijkstra -domain=gridnav width=10 heigt=10 timeout=30 seed=1 > data/10
MyExecutable -algorithm=Dijkstra -domain=gridnav width=10 heigt=10 timeout=30 seed=2 > data/11
MyExecutable -algorithm=Dijkstra -domain=gridnav width=10 heigt=10 timeout=30 seed=3 > data/12
MyExecutable -algorithm=Dijkstra -domain=gridnav width=10 heigt=10 timeout=30 seed=4 > data/13
MyExecutable -algorithm=Dijkstra -domain=gridnav width=10 heigt=10 timeout=30 seed=5 > data/14
MyExecutable -algorithm=Dijkstra -domain=tiles timeout=30 seed=1 > data/15
MyExecutable -algorithm=Dijkstra -domain=tiles timeout=30 seed=2 > data/16
MyExecutable -algorithm=Dijkstra -domain=tiles timeout=30 seed=3 > data/17
MyExecutable -algorithm=Dijkstra -domain=tiles timeout=30 seed=4 > data/18
MyExecutable -algorithm=Dijkstra -domain=tiles timeout=30 seed=5 > data/19
MyExecutable -algorithm=Greedy -domain=gridnav width=10 heigt=10 timeout=30 seed=1 > data/20
MyExecutable -algorithm=Greedy -domain=gridnav width=10 heigt=10 timeout=30 seed=2 > data/21
MyExecutable -algorithm=Greedy -domain=gridnav width=10 heigt=10 timeout=30 seed=3 > data/22
MyExecutable -algorithm=Greedy -domain=gridnav width=10 heigt=10 timeout=30 seed=4 > data/23
MyExecutable -algorithm=Greedy -domain=gridnav width=10 heigt=10 timeout=30 seed=5 > data/24
MyExecutable -algorithm=Greedy -domain=tiles timeout=30 seed=1 > data/25
MyExecutable -algorithm=Greedy -domain=tiles timeout=30 seed=2 > data/26
MyExecutable -algorithm=Greedy -domain=tiles timeout=30 seed=3 > data/27
MyExecutable -algorithm=Greedy -domain=tiles timeout=30 seed=4 > data/28
MyExecutable -algorithm=Greedy -domain=tiles timeout=30 seed=5 > data/29
```
