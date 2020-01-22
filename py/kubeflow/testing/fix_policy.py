import os
import yaml

if __name__ == "__main__":
  with open("/home/jlewi/Downloads/kubeflow.ci.deployment.policy.txt") as hf:
    policy = yaml.load(hf)

  for b in policy["bindings"]:
    members = b["members"]

    new_members = []
    for m in members:
      exclude = False
      for p in [
        "serviceAccount:kf-vmaster-",
      ]:
        if m.startswith(p):
          exclude = True
          continue

      if exclude:
        continue
      new_members.append(m)

    b["members"] = new_members

  out_file = "/tmp/kubeflow.ci.deployment.policy.restored.txt"
  with open(out_file, "w") as hf:
    yaml.dump(policy, hf)
  print(f"Wrote file {out_file}")
