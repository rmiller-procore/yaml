CloudFormation Investigation
============================

*Goal* Spike an investigation to disover the usability of CloudFormation for
infrastructure provisioning in the Procore Envs.


Tools in Context
----------------
- AWS CloudFormation [CloudFormation Docs](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html)
- GoFormation [GoFOrmation on GitHub](https://github.com/awslabs/goformation)
    AWS Go Implementation to generate CloudFormation templates programmatically
- Go YAML library [Go YAML on GitHub](https://github.com/go-yaml/yaml)
    Ancilliary investigation. Go does not natively support YAML, so 
    a library to support YAML was found online for this purpose. This was
    due to the original design decision of a Procore YAML input file
- cloudformation-ruby-dsl [Ruby DSL on GitHub](https://github.com/bazaarvoice/cloudformation-ruby-dsl)
    Secondary investigation. Due to our strong association to Ruby code,
    this facility was also investigated as a possible approach to IAC.


Investigations Performed
------------------------
1. CloudFormation YAML templates written by hand
2. Ruby coded CloudFormation applications (CloudFormation Ruby DSL)
3. Creation of Sandbox VPC for AMI creation process.
4. GoFormation generated YAML templates
5. YAML driven GoFormation generation of YAML templates


Investigative Outcomes
----------------------
1. CloudFormation YAML templates written by hand
    This approach was rather clunky, but notably it was our first exposure to
    the template formats, so we didn't have the knowledge of the template
    parameters necessary to create the test infrastructure off the cuff.
2. Ruby Coded CloudFormation applications
    This investigation focussed on Ruby generated infrastructure using the 
    cloudfomration-ruby-dsl gem. This was one of the more efficient
    coding paths we took, likely due to the Ruby familiarity. The main
    drawback of this approach is that the code that is geneated are Ruby
    programs that directly control the infrastructure, rather then 
    producing the YAML template to be executed outside the scope of the program.
3. Creation of Sandbox VPC for AMI generation
    We reflexed our Ruby coding expertise to create a basic VPC required
    for Brian to build custom AMIs in the sandbox account. Though it was quick
    to create, we had to fix some odd issues in the AWS Console which we could
    not solve programmatically in the time allotted.
4. GoFormation generated YAML templates
    In this path of investigation we wrote some simple Go programs which will
    generate CloudFormation YAML templates to be executed outside the scrope
    of the Go program. This was fairly straight-forward, as we had a better
    grasp of the template parameters to pass in Go. The only concern in this 
    pass was that Go, for Reid and Martina, is a newer language, and learning some of
    its more obscure concepts will be a learning experience. I suspect this 
    may also affect other team members, but is not seen as a significant
    impediment.
5. YAML driven GoFormation generation of YAML templates
    The final phase of investigation included as stab at parsing the proposed
    procore.yml file, and generating CloudFormation YAML for external execution.
    We succeeded in performing these tasks. There were a couple of short-cuts
    taken that will require further research, for instance, we've split the
    procore.yml file into it's subsidiary sections for processing at this
    stage. We need to gain deeper understanding of the Go YAML library to best
    use it for our needs.


Additional Efforts
------------------
One follow-on experiment was to better understand the data strcuture generated
by the GoYAML library when it slurps up a YAML file. I put a small effort
to flesh out this data structure, and was able to put together a first
run at parsing the `procore.yml` file's "interesting" sections from the 
perspective of CloudFormation.

Conclusions
-----------
CloudFormation seems like it could be a path forward for us. Especially, if we
wish to generate the templates to be committed somewhere. Though the Ruby DSL
solution was very efficient for coding speed, does not seem to be well
supported that I can tell. The other negative of this solution was that it
generates Ruby programs to control the infrastucture directly. And if that is
not the path of control we seek, then that would not be the way forward. 

GoFormation seems promising for a number of reasons:
1. It creates the YAML templates for external execution
2. It does not directly control anything but the output of templates.
3. The tool is produced and supported by AWS Labs

In general, there is an assumed learning curve, since some will be learning
both CloudFormation and Go coding constructs along the way. The steepness of the
curve will likely vary with the individuals exposure to programing language
and data structures in Go.


Next Steps
==========
Planning session - Initial steps forward should include a planning session to
design the system to be built. This will allow designation of the parallelism
and separation of work efforts.
