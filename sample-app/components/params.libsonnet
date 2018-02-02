{
  global: {
    // User-defined global parameters; accessible to all component and environments, Ex:
    // replicas: 4,
  },
  components: {
    prow_pod: {
      app_dir: "test/workflows",
      component: "workflows",
      name: "test",
      namespace: "test-pod",
      prow_env: "JOB_NAME=tf-k8s-presubmit,JOB_TYPE=presubmit,PULL_NUMBER=358,REPO_NAME=k8s,REPO_OWNER=tensorflow,BUILD_NUMBER=7759",
      testing_subdir: "kubeflow_testing",
    },
  },
}
