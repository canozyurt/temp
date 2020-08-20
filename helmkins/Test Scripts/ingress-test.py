import subprocess
import re

units = [
["      branch & test", "-n test --set branch=feature/unity-1111 --set setenv=false"],
["      devel  & test", "-n test"],
["              pilot", "-n pilot"],
["               live", "-n live"],
]

finder = re.compile("  - host.*")
report = []

for unit in units:
  command = "helm upgrade --install denemeler helm-btpl/ " + unit[1] + " --debug --dry-run -f projects/somebrand-yksquestions-webclient/values.yaml"
  output = subprocess.check_output(command, shell=True)
  finding = re.finditer(finder, output.decode('utf-8'))
  report.append([unit[0], finding])

for repo in report:
  for rep in repo[1]:
    print("{0} {1}" .format(repo[0], rep.group(0)))
