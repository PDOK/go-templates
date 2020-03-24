# Go templates

Program for generating files using templates with json as input. Additional KVP's can be passed to the main program

Initialy developed for generating K8s Custom Resources



```cmd
docker run -it -v $(pwd)/templates:/templates -v $(pwd)/output:/output go-templates:0.0.1 go-templates /templates/intake.json /templates/atom.template /output/atom.yaml update_version=1 source_key=dir/dir
```

