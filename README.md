Open Sceheduler Extender
=======================

# Aim
- easy extension of the scheduler regardless of where you are hosting your kubernetes clusters!
  - easy to program (declarative Go syntax provided by the expr library allows for ease of writing extensions)
- flexiblility in scheduling


# Motivation
three main reasons:
1. most importantly, there are cases directly modifying the default kube-scheduler is impossible, such as in managed Kubernetes clusers like EKS where the control plane resorces are unmodifiable
2. allows for even greater flexibility as you can querying metrics that are not normally exposed to the standard scheduler! this includes things like number of pods with a dedicated ENI attached in EKS etc.
3. such a general schedular extending tool does not yet exist. moreover, most tools that help with the scheduling process are specific to a certain use case.


# Usage

### 1. Buid docker image

```

docker build . -t ose/ose:latest

```

### 2. Deploy `ose` in `kube-system` namespace
ConfigMap in [ose.yaml](ose.yaml) contains the scheduler config.

```

kubectl apply -f ose.yaml

```



### 3. Test scheduler

You should see that `ose-test-container` will be scheduled by `ose-scheduler`.

```

$ kubectl create -f test.yaml

$ kubectl describe pod/ose-test-container

```


# Drawbacks
1. will be slower than when directly modyfying the scheduler in the control plane and recompiling
2. since this is a pod that runs alongside ur other applications, some additional maintanence might be required
