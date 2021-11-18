# galasa-kubernetes-operator

```
~/go/src/k8s.io/code-generator/generate-groups.sh all github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis "galasaecosystem:v2alpha1" -h hack/boilerplate/boilerplate.go.txt
hack/generate-knative.sh injection github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis "galasaecosystem:v2alpha1"
ko apply -f config

```
/Users/jamesdavies/go/src/github.com/Jimbo4794/galasa-kubernetes-operator/hack/boilerplate/boilerplate.go.txt


Things left to sort out:
- replicas need to be tied to something so it can scale them down
- reset if something changes
- get the logging better