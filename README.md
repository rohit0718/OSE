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


# Drawbacks
1. will be slower than when directly modyfying the scheduler in the control plane and recompiling
2. since this is a pod that runs alongside ur other applications, some additional maintanence might be required
