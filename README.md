# GO Templates

Program for generating files using go templates with json as input. Additional KVP's can be passed which will also be applied on the template. The additional KVP's may overlap with values in the json file, in such an event the KVP's will overrule the value in the json file.  

Initialy developed for generating K8s Custom Resources from within a workflow using json from a ConfigMap and workflow specific parameters.

The main program takes the following arguments:
- a json inputfile: location to a json file used for populating the template
- a templatefile: the actual go template
- a outputfile
- any key value pairs: zero or more kvp's in the format 'key(.key(.key)...)=value' these values wil also be used for populating the template

## example

using the input and template from the exmaple folder .... 

```cmd
docker run -it -v $(pwd)/example:/templates -v $(pwd)/output:/output pdokdev/go-templates:0.0.1 go-templates /templates/krm2018_intake.json /templates/atom.template /output/atom.yaml delivery.update_version=25-09-2019T09:00 delivery.source_key=rws/kaderrichtlijnmarienestrategie2018/25-09-2019T09:00
```

