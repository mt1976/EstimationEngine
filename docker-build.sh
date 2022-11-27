#Before we can use the script, we have to make it executable with the chmod command:
#chmod +x ./go-executable-build.sh
#then we can use it  ./go-executable-build.sh yourpackage
#!/usr/bin/env bash
docker image build -t mt1976/mc .
