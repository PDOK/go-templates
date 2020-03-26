# Go templates

Program for generating files using templates with json as input. Additional KVP's can be passed, they will aslo be applied on the template.

Initialy developed for generating K8s Custom Resources

The main program take the following arguments:
- json inputfile: location to a json file used for populating the template
- templatefile: the actual go template
- outputfile
- key value pair .... : zero or more kvp in the format 'key(.key)=value' these values wil also be used for populating the template

you can run the docker container from the source after a docker build
```cmd
docker run -it -v $(pwd)/templates:/templates -v $(pwd)/output:/output go-templates:0.0.1 go-templates /templates/intake.json /templates/atom.template /output/atom.yaml delivery.update_version=1 delivery.source_key=dir/dir
```

