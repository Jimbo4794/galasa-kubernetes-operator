# galasa-kubernetes-operator

```
# Deep copy and Client gen
hack/generate-groups.sh all github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis "galasaecosystem:v2alpha1" -h hack/boilerplate/boilerplate.go.txt

#Knative Injection clients
hack/generate-knative.sh injection github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis "galasaecosystem:v2alpha1" -h hack/boilerplate/boilerplate.go.txt

#Build and deploy
ko apply -f config
```
