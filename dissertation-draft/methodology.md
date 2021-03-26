# Code Smells Extraction

For this study, code smells are gathered from multiple sources. The research by Sharmann et al. on the code smells in Puppet is initially looked at to set the direction of the study. As the study suggested, the smells are categorized by its character, which results in two categories of code smells — implementation smells and design smells. However, as Kubernetes is set to be in focus of this project, implementation smells will be first put in concern.

In accordance with the study conducted by Schwarz et al., the smells can also be categorized regarding their technological aspect i.e. there are 3 types of smells: Technology Specific, Technology Dependent, and Technology Agnostic. Each of the type corresponds to a different level of technological reliant, i.e.:
Technology Specific Smells: a group of smells that exist only in Kubernetes context, which can be derived from style guides and best practices;
Technology Dependent Smells: ones that cannot be applied directly from one IaC tool to another, and required adjustment in its detection method; and
Technology Agnostic Smells: smells that exist across IaCs, which its detection method remains unchanged with consideration of a different IaC.
Therefore, this set the exploration topics for this project. Each of the categories will be examined. Ones laid in Technology Agnostic kind will be attempted to port to Kubernetes and to find the first question asked in the research questions section.

Therefore, this set the exploration topics for this project. Each of the categories will be examined. Ones laid in Technology Agnostic will be attempted to be ported to Kubernetes to answer the first research question.

Hence, two tasks are established at this stage: porting code smells from the existing studies, and dissecting a set of technological specific and technological dependent smells from Kubernetes best practices in grey literatures — best practices and style guides.

## Porting Code Smells

By looking at smell catalogues from both Puppet-based and Chef-base studies, similarities are shown that both are likely to share a set of smells, which occurs in the implementation i.e. can be found in a single file and has no relationship between others in the same project. Therefore, this suggests that those should be firstly considered.

The smells will be listed along with their portability and justification. The source of the smell will be firstly presented, followed by its class, description, portability report, and modifications made for that specific smell if there are.

### Improper Alignment (IA)

- Source: Sharma et al., Schwarz et al.
- Class: Tech Agnostic
- Description: This smell is detected when the code contains incorrect and inconsistent indentation. This could lead to incorrect or invalid interpretation when the code is executed.
- Portability: Portable
- Detection Method: Using a general YAML parser would be enough. As said earlier, YAML enforces this already, if a script failed the detector, it would be deemed to contain the smell.

### Long Statements (LS)

- Source: Sharma et al., Schwarz et al.
- Class: Tech Agnostic
- Description: This smell means too long line of codes, which exceeds the screen. It is applicable directly, as a statement should not be too long and this was dissected from Fowler, which used to be only applied to general purpose programming languages.
- Portability: Portable.
- Detection Method: The smells can be detected by counting the characters existing on each line. The limit is set to 140.

### Unguarded Variable (UV)

- Source: Sharma et al., Schwarz et al.
- Class: Tech Agnostic
- Description: A variable is not enclosed in braces when bein interpolated in a string. However, string interpolation cannot be performed using vanilla Kubernetes manifest. However, it is doable if some additional are used, e.g., Kustomize, helm and so on; we will not consider this as it goes beyond the scope of this study.
- Portability: Unportable
- Detection Method: Indefinable

### Misplaced Attribute (MAt)

- Source: Sharma et al., Schwarz et al.
- Class: Tech Agnostic
- Description: A specific order of information inside a resource should be followed.
- Portability: Unportable
- Justification: It is not required for Kubernetes manifests.
- Detection Method: Indefinable

### Long Manifest (LR)

- Source: Sharma et al., Schwarz et al.
- Class: Tech Agnostic
- Description: Execute and bash resources being too long.
- Portability: Portable
- Justification: In a Kubernetes file, one file can contain multiple resources, which in turn results in a file being too long.
- Detection Method: Undefined

### Avoid Comments (AC)

- Source: Sharma et al., Schwarz et al.
- Class: Tech Agnostic
- Description: Since IaCs are described in declarative languages, the scripts should always be self-explained and comments should be avoided.
- Portability: Portable
- Detection Method: This could be detected using comment syntax of YAML.

### Improper Quote Usage (IQU)

- Source: Sharma et al., Schwarz et al.
- Class: Tech Dependent
- Description: About Chef, it is suggested that Boolean values should not be quoted, and variables should not be single-quoted. Similarly for Kubernetes files, it is possible to use both single quote and double quote to signify strings, and even strings with no quote are considered as normal strings. Therefore, this is likely to rather be a YAML specific smell.
- Portability: Portable
- Detection Method: Undefined

### Hyphens (HP)

- Source: Schwarz et al.
- Class: Tech Specific
- Description: This smell appears in Chef. However, there is no such prohibition in Kubernetes.
- Portability: Unportable
- Detection Method: Undefinable

### Missing Default Case (MD)

- Source: Sharma et al.
- Class: Tech Specific
- Description: Sharma et al. describes this as the absence of default case in "case" or selector statement. In Kubernetes, this approach seems applicable to resource selectors. However, evidence preventing this kind of naming in resource selectors can be found. Therefore, this smell is considered as unportable.
- Portability: Unportable
- Detection Method: Undefinable

### Inconsistent Naming Convention (INC)

- Source: Sharma et al.
- Class: Tech Agnostic
- Description: The current naming deviates from the convention. This smell is likely to appear in Kubernetes manifest as well.
- Portability: Portable.
- Detection Method: Undefined

### Complex Expression (CE)

- Source: Sharma et al.
- Class: Tech Dependent
- Description: A Puppet script containing complex and difficult to understand expressions is considered to be smelly. However, it is seemingly impossible in Kubernetes, as there is no expression contained in Kubernetes-YAMl manifests.
- Portability: Unportable
- Detection Method: Undefinable

### Duplicate Entity (DE)

- Source: Sharma et al.
- Class: Tech Dependent
- Description: As raised by Sharma et al., this smell means duplicate parameters present in the code. In Kubernetes, an equivalent to attribute in Puppet would be a field in its YAML specification of resources. A field should not be defined multiple times in a resource definition.
- Portability: Portable
- Detection Method: Undefined

### Invalid Property Value (PV)

- Source: Sharma et al.
- Class: Tech Dependent
- Description: In Puppet, this smell is indicated from invalid property or attributes usages. Where in Kubernetes, it is unclear on how this smell can be applied in general. Nevertheless, there exists an applicable case, in which incorrect definition of the CPU limits and memory limits is considered as this smell.
- Portability: Portable
- Detection Method: Undefined

### Incomplete Tasks (IT)

- Source: Sharma et al.
- Class: Tech Agnostic
- Description: When "FIXME" and "TODO" appear in the comments, a script is considered to contain the smell.
- Portability: Portable
- Detection Method: Looking for comments with "FIXME" and "TODO" written.

### Deprecated Statement Usage (DS)

- Source: Sharma et al.
- Class: Tech Dependent
- Description: An instance of smell occurs when a deprecated statement is used. In Kubernetes, this smell seems applicable. However, a version used by Kubernetes manifests is only required to match its targeted container cluster. Since it is cluster dependent, the smell cannot be detected regardless of the cluster context.
- Portability: Unportable
- Detection Method: Undefinable

### Incomplete Condition (IC)

- Source: Sharma et al.
- Class: Tech Specific
- Description: Sharma et al. describes this as "if ... elsif" block without terminating "else". This is obviously not usable in the context of Kubernetes, since YAML does not have such capability to use conditional statements.
- Portability: Unportable
- Detection Method: Undefinable

## Kubernetes Specific Smells

Next, this section attempts to define a set of smells scoped to Kubernetes to answer the research question 2. By searching through guidelines and best practices documented about writing Kubernetes manifests, There are numerous aspects found. However, due to the time constraint, there only 4 of them extracted as code smells.

The first conern selected is that resources relating to container specification should provide a readiness probe. To elaborate, the purpose of readiness probe is to supply a communiation channel between the container being created, and kubelet — the cluster activity controller, to tell if the container is ready for receving traffic.
Smells

- Missing Readiness Probes,
- Duplicated Livenes and Readiness Probes,
- No Resources Limits,
- Hard Coded Secrets.
