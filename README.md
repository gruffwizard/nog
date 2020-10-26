# nog
Containerised Quarkus development
version 0.0.1-alpha

# options

nog qs name
nog qs name

nog dev -s ~/foo
nog dev -l srcvol

nog

--src-vol src volume
--mvn-vol mvn volume

-d mvn dir (defaults to ~/.m2)
-s src dir  (defaults to . )
-i launch theia ide (defaults to false) sets -m to nog-maven and -l to nog-
-m mvn vol  
-l  src vol

//






# run a quickstart locally

nog quickstart name

# run a quickstart using provided ide and in a src vol

nog quickstart -i name


# run dev mode on local src
nog dev

# run full container mode
nog dev -i

# run local dev mode using a maven volume
nog dev -m  vol-name

# run local dev mode using a different but local maven repo
nog dev  -d  file-name


# run full container mode with a different maven repo
nog dev -i -m mvn-vol-name


# run full container mode with a different maven repo and a name src repo
nog dev -i -s foo -m mvn-vol-name


nog volume ls



local (aka current dir) or a named volume using the theia editor

maven: shared local cache, shared volume or named volume.  



---
short cuts

-q demo     Checkouts the demo repo into a volume and then copies the specific demo into current directory.    Then runs in dev mode.

-e  lauches editor mode

-s uses a named volume for holding the source (implies -e )


----

nog dev -e -s -q -m


nog dev


starts a local dev session with src in current dir and maven in a shared volume "nog-maven"   Quarkus runs on the command line.

nog dev -e  starts a local dev with src in current dir and maven in a volume and runs quarkus via the editor...   (launches a browser to the editor)

nog dev -s (uses a named volume to hold source)
nog dev -m sets a different vol name for maven repo

nog dev -g gitrepo  

nog dev -q sample - checkouts out quarkus demo git repo into src vol and copies sample dir into

nog dev  

starts a simple container dev session with src in current dir



nog dev -e -s srcloc -m mavenloc -g [ cmds ]


nog dev -g dddd

checkout git repo into src volume using default maven volume  nog-maven-default and fire up theia and quarkus

nog dev -g ddd -s vol:  /  -s .  

-s vol means checkout into this volume
-s dir means checkout as subdir in the dir then launch as is
