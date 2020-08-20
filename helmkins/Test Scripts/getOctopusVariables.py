import sys, requests, json
headers = {"X-Octopus-ApiKey": "API-111111111111111111111"}
environments = {
"Environments-1": "test",
"Environments-2": "pilot",
"Environments-21": "live",
"Environments-23": "dev",
"Environments-41": "alive",
"Environments-61": "drc",
"Environments-121": "preview"
}

finalList = {}
for env in environments:
  finalList[env] = []

def getName(key):
  return environments[key]

def getAllProjects():
  return requests.get("http://octopus.dev.example.com.tr/api/projects/all", headers=headers).text


def findVariablesURL(projectName):
  allProjects = json.loads(getAllProjects())
  for i in allProjects:
        if i["Name"] == projectName:
                return i["Links"]["Variables"]

if len(sys.argv) != 2:
  print("One project name must be provided")
  exit(1)

variables = json.loads(requests.get("http://octopus.dev.example.com.tr" + \
            findVariablesURL(sys.argv[1]), \
            headers=headers).text)["Variables"]

for var in variables:
  if not var["Scope"]:
    for env in environments.keys():
      finalList[env].append({var["Name"]: var["Value"]})
  else:
    for env in var["Scope"]["Environment"]:
      finalList[env].append({var["Name"]: var["Value"]})

for env in finalList:
  print("\n  "+ getName(env) + ":")
  for var in finalList[env]:
    for k, v in var.items():
      if k == "Docker.File": continue
      print ("    {0}: {1}".format(k.replace(":","__"), v))
