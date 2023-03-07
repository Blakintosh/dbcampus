from __future__ import print_function

from apiclient import discovery
from httplib2 import Http
from oauth2client import client, file, tools

def retrieveHandler(data):
    averages = getAverages(getResults(data["form_id"]))
    return averages


def getResults(form_id):
    SCOPES = "https://www.googleapis.com/auth/forms.responses.readonly"
    DISCOVERY_DOC = "https://forms.googleapis.com/$discovery/rest?version=v1"

    store = file.Storage('token.json')
    creds = None
    if not creds or creds.invalid:
        flow = client.flow_from_clientsecrets('credentials.json', SCOPES)
        creds = tools.run_flow(flow, store)
    service = discovery.build('forms', 'v1', http=creds.authorize(
        Http()), discoveryServiceUrl=DISCOVERY_DOC, static_discovery=False)


    result = service.forms().responses().list(formId=form_id).execute()
    return result



def getAverages(result):
    responses = dict()
    for x in result["responses"]:
        for k,j in x["answers"].items():
            if k not in responses:
                responses[k] = 0
            responses[k] += int(j["textAnswers"]["answers"][0]["value"])
            # print(j,k)
        # print("\n\n")

    for l in responses:
        responses[l] /= len(result["responses"])

    return responses


# form_id = '1eBp4r7shhohOAAi36qpf2Wkbj1QGJsxjmRvbXQXmJs4'
# result = getResults(form_id)
# print(getAverages(result))