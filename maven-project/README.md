# Maven demo project

Demo maven project to build Java library/app. I've got this working with VSCode and openjdk16 on Mac M1.


```bash
# install maven
brew install mvn

# check java version for maven (this is used in pom.xml)
# make sure maven.compiler.source|target in pom.xml is less than maven's java
mvn --version | grep -i java

# build a jar
mvn clean install

# run
java -cp target/demo-1.0-SNAPSHOT.jar com.example.App
```

## Useful VSCode extensions

- Java Extension Pack
- Java Test Runner
- Language Support for Java(TM) by Red Hat
