# Code Smells Extraction

For this study, code smells are gathered from multiple sources. The research by Sharmann et al. on the code smells in Puppet is initially looked at to set the direction of the study. As the study suggested, the smells are categorized by its character, which results in two categories of code smells — implementation smells and design smells. However, as Kubernetes is set to be in focus of this project, implementation smells will be first put in concern.

In accordance with the study conducted by Schwarz et al., the smells can also be categorized regarding their technological aspect i.e. there are 3 types of smells: Technology Specific, Technology Dependent, and Technology Agnostic. Each of the type corresponds to a different level of technological reliant, i.e.:

- **Technology Specific Smells:** a group of smells that exist only in Kubernetes context, which can be derived from style guides and best practices;
- **Technology Dependent Smells:** ones that cannot be applied directly from one IaC tool to another, and required adjustment in its detection method; and
- **Technology Agnostic Smells:** smells that exists across IaCs, which its detection method remain unchange with consider a different IaC.

Therefore, this set the exploration topics for this project. Each of the categories will be examined. Ones laid in Technology Agnostic will be attempted to be ported to Kubernetes to answer the first research question.

Hence, two tasks are establisehd at this stage: porting code smells from the existing studies, and dissecting a set of technological specific and technological dependent smells from Kubernetes best practices in grey literatures — best practices and style guides.

## Porting Code Smells

By looking at smell catalogues from both Puppet-based and Chef-base studies, similarities is shown that both are likely to share a set of smells, which occurs in the implementation i.e. can be found in a single file and has no relationship between others in the same project. Therefore, this suggests that those should be fisrtly considered.

The smells will be listed along with their portability and justification. The source of the smell will be firstly presented, followed by its class, description, portability report, and modifications made for that specific smell if there are.

### Improper Alignment (IA)

- **Source:** Sharma et al., Schwarz et al.
- **Class:** Tech Agnostic
- **Description:** This smell is detected when the code contains incorrect and inconsisttent indetations. This could lead to incorrect or invalid interpretation when the code is executed.
- **Portability:** It is portable. Considering the fact that Kubernetes manifests are defined using YAML, the script is already constrained by YAML syntax.
- **Detection Method:** Using a general YAML parser would be enough. As said earlier, YAML enforces this already, if a script failed the detector, it would be deemed to contain the smell.

### Long Statements (LS)

- **Source:** Sharma et al., Schwarz et al.
- **Class:** Tech Agnostic
- **Description:** too long line of codes, which exceeds the screen.
- **Portability:** Portable.
- **Detection Method:** The smells can be detected by counting the characters existing on each line. The limit set to 140.

### Unguarded Variable (UV)

- **Source:**
- **Class:**
- **Description:**
- **Portability:**
- **Detection Method:**

### Misplaced Attribute (MAt)

- **Source:**
- **Class:**
- **Description:**
- **Portability:**
- **Detection Method:**

### Long Resource (LR)

- **Source:**
- **Class:**
- **Description:**
- **Portability:**
- **Detection Method:**

### Avoid Comments (AC)

- **Source:**
- **Class:**
- **Description:**
- **Portability:**
- **Detection Method:**

### Improper Quote Usage (IQU)

- **Source:**
- **Class:**
- **Description:**
- **Portability:**
- **Detection Method:**

### Hyphens (HP)

- **Source:**
- **Class:**
- **Description:**
- **Portability:**
- **Detection Method:**

### Missing Default Case (MD)

- **Source:**
- **Class:**
- **Description:**
- **Portability:**
- **Detection Method:**

### Inconsistent Naming Convention (INC)

- **Source:**
- **Class:**
- **Description:**
- **Portability:**
- **Detection Method:**

### Complex Expression (CE)

- **Source:**
- **Class:**
- **Description:**
- **Portability:**
- **Detection Method:**

### Duplicate Entity (DE)

- **Source:**
- **Class:**
- **Description:**
- **Portability:**
- **Detection Method:**

### Invalid Property Value (PV)

- **Source:**
- **Class:**
- **Description:**
- **Portability:**
- **Detection Method:**

### Incomplete Tasks (IT)

- **Source:**
- **Class:**
- **Description:**
- **Portability:**
- **Detection Method:**

### Deprecated Statement Usage (DS)

- **Source:**
- **Class:**
- **Description:**
- **Portability:**
- **Detection Method:**

### Incomple Condition (IC)

- **Source:**
- **Class:**
- **Description:**
- **Portability:**
- **Detection Method:**

## Kubernetes Specific Smells

After porting the existing smells found in the existing studies, the smells scoped to only Kubernetes is put in concern. By searching through what documented on the best practices about writing Kubernetes manifests, a specific set of patterns can be found, those are:

Smells

- Missing Readiness Probes,
- Duplicated Livenes and Readiness Probes,
- No Resources Limits,
- Hard Coded Secrets.
